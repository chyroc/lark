package helper

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type ASTResult struct {
	Consts  map[string][]Const
	Types   map[string]*ast.TypeSpec
	Structs map[string]*Struct
}

type Struct struct {
	Name   string   `json:"name,omitempty"`
	Fields []*Field `json:"fields,omitempty"`
}

type Field struct {
	Name          string         `json:"name,omitempty"`
	NameType      string         `json:"name_type,omitempty"`
	RealType      string         `json:"real_type,omitempty"`
	JSONTag       string         `json:"json_tag,omitempty"`
	JSONOmitempty bool           `json:"json_omitempty,omitempty"`
	UnmarshalTool string         `json:"unmarshal_tool,omitempty"`
	Builder       string         `json:"builder,omitempty"`
	BuilderArgs   map[int]*Field `json:"builder_args,omitempty"`
	BuilderIndex  int            `json:"builder_index,omitempty"`
}

type Const struct {
	NameType string `json:"name_type,omitempty"`
	RealType string `json:"real_type,omitempty"`
}

func (r Field) ZeroString() string {
	if strings.HasPrefix(r.RealType, "*") {
		return "nil"
	} else if strings.HasPrefix(r.RealType, "[]") {
		return "len"
	} else if strings.HasPrefix(r.RealType, "map") {
		return "len"
	} else if strings.HasPrefix(r.RealType, "string") {
		return "\"\""
	} else if strings.HasPrefix(r.RealType, "bool") {
		return "false"
	} else {
		return "0"
	}
}

func ParseAST(dir string) (*ASTResult, error) {
	fset := token.NewFileSet()
	nodes, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	res := &ASTResult{
		Consts:  map[string][]Const{},
		Types:   map[string]*ast.TypeSpec{},
		Structs: map[string]*Struct{},
	}

	for _, node := range nodes {
		ast.Inspect(node, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.TypeSpec:
				res.Types[n.Name.Name] = n
				return false
			case *ast.GenDecl:
				if n.Tok == token.CONST {
					for k, v := range parseConsts(n) {
						res.Consts[k] = append(res.Consts[k], v...)
					}
					return false
				}
			}
			return true
		})
	}

	for k, v := range res.Types {
		fields, err := res.parseStruct(v.Type)
		if err != nil {
			return nil, err
		} else if fields == nil {
			continue
		}
		res.Structs[k] = &Struct{
			Name:   k,
			Fields: fields,
		}
		builderToBuilderArgs := map[string]map[int]*Field{}
		for _, field := range res.Structs[k].Fields {
			field := field
			builderToBuilderArgs[field.Name] = map[int]*Field{}
			if field.Builder != "" {
				if _, ok := builderToBuilderArgs[field.Builder]; !ok {
					builderToBuilderArgs[field.Builder] = map[int]*Field{}
				}
				builderToBuilderArgs[field.Builder][field.BuilderIndex] = field
			}
		}
		for _, v := range res.Structs[k].Fields {
			if x, ok := builderToBuilderArgs[v.Builder]; ok {
				v.BuilderArgs = x
			}
		}
	}

	return res, nil
}

func (r *ASTResult) parseStruct(expr ast.Expr) ([]*Field, error) {
	structType, ok := expr.(*ast.StructType)
	if !ok {
		return nil, nil
	}
	res := []*Field{}
	for _, field := range structType.Fields.List {
		tmp, err := r.parseField(field)
		if err != nil {
			return nil, err
		} else if tmp == nil {
			continue
		} else if tmp.NameType == "componentBase" {
			continue
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (r *ASTResult) parseField(f *ast.Field) (*Field, error) {
	name := ""
	if len(f.Names) > 0 {
		name = f.Names[0].Name
	}
	res := &Field{
		Name: name,
	}

	// tag
	if f.Tag != nil {
		tag := reflect.StructTag(strings.Trim(f.Tag.Value, "`"))

		if jsonTag, _ := tag.Lookup("json"); jsonTag != "" {
			res.JSONTag = strings.Split(jsonTag, ",")[0]
			res.JSONOmitempty = strings.Contains(jsonTag, "omitempty")
		}

		res.UnmarshalTool, _ = tag.Lookup("unmarshal")

		builder, _ := tag.Lookup("builder")
		if builder != "" {
			l := strings.Split(builder, ",")
			if len(l) > 0 {
				res.Builder = l[0]
			}
			if len(l) > 1 {
				tmp, err := strconv.Atoi(l[1])
				if err != nil {
					return nil, err
				}
				res.BuilderIndex = tmp
			} else {
				res.BuilderIndex = 1
			}
		}
	}

	// type
	{
		nameType, err := r.exprToNameString(f.Type, false)
		if err != nil {
			return nil, err
		}
		realType, err := r.exprToNameString(f.Type, true)
		if err != nil {
			return nil, err
		}
		res.NameType = nameType
		res.RealType = realType
	}

	return res, nil
}

func (r *ASTResult) exprToNameString(expr ast.Expr, isReal bool) (string, error) {
	switch expr := expr.(type) {
	case *ast.Ident:
		if isReal {
			identTypeSpec := r.getIdentTypeSpec(expr)
			if identTypeSpec == nil {
				return expr.Name, nil
			} else if isExprStruct(identTypeSpec.Type) || isExprInterface(identTypeSpec.Type) {
				// 如果是结构体, identTypeSpec.Type 中没有结构体的 name, 这里不递归调用
				return identTypeSpec.Name.Name, nil
			}
			return r.exprToNameString(identTypeSpec.Type, isReal)
		}
		return expr.Name, nil
	case *ast.ArrayType:
		name, err := r.exprToNameString(expr.Elt, isReal)
		if err != nil {
			return "", err
		}
		return "[]" + name, nil
	case *ast.StarExpr:
		name, err := r.exprToNameString(expr.X, isReal)
		if err != nil {
			return "", err
		}
		return "*" + name, nil
	case *ast.SelectorExpr:
		x, err := r.exprToNameString(expr.X, isReal)
		if err != nil {
			return "", err
		}
		return x + "." + expr.Sel.Name, nil
	case *ast.MapType:
		key, err := r.exprToNameString(expr.Key, isReal)
		if err != nil {
			return "", err
		}
		value, err := r.exprToNameString(expr.Value, isReal)
		if err != nil {
			return "", err
		}
		return "map[" + key + "]" + value, nil
	case *ast.InterfaceType:
		return "interface{}", nil
	default:
		spew.Dump(expr)
		panic("exprToNameString.default")
	}
}

func (r *ASTResult) getIdentTypeSpec(ident *ast.Ident) *ast.TypeSpec {
	name := ident.Name
	obj := ident.Obj
	if obj != nil {
		data, _ := obj.Decl.(*ast.TypeSpec)
		return data
	}
	typ := r.Types[name]
	if typ != nil {
		return typ
	}
	return nil
}

func isExprStruct(expr ast.Expr) bool {
	_, ok := expr.(*ast.StructType)
	return ok
}

func isExprInterface(expr ast.Expr) bool {
	_, ok := expr.(*ast.InterfaceType)
	return ok
}

func parseConsts(n *ast.GenDecl) map[string][]Const {
	consts := map[string][]Const{}
	for _, spec := range n.Specs {
		val, ok := spec.(*ast.ValueSpec)
		if ok {
			for _, name := range val.Names {
				if val.Type == nil {
					continue
				}
				realType := ""
				err := json.Unmarshal([]byte(name.Obj.Decl.(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value), &realType)
				if err != nil {
					panic(err)
				}
				consts[val.Type.(*ast.Ident).Name] = append(consts[val.Type.(*ast.Ident).Name], Const{
					NameType: name.Name,
					RealType: realType,
				})
			}
		}
	}
	return consts
}
