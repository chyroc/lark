package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/chyroc/lark/.github/script/helper"
)

func main() {
	typ := flag.String("type", "", "")
	flag.Parse()
	fileName := os.Getenv("GOFILE")
	structName := *typ

	ast, err := helper.ParseAST(filepath.Dir(fileName))
	if err != nil {
		log.Fatalln(err)
	}

	err = rewritePath(fileName, structName, ast)
	if err != nil {
		log.Fatalln(err)
	}
}

func rewritePath(path string, structName string, ast *helper.ASTResult) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	structType, ok := ast.Structs[structName]
	if !ok {
		return fmt.Errorf("not found struct %s", structName)
	}
	content := clearGenerate(string(bs), structName)
	attrs := generate(structName, structType)

	return ioutil.WriteFile(path, []byte(helper.ClearString(content+attrs, "\n\n\n", "\n\n")), 0o644)
}

func clearGenerate(s, structName string) string {
	l := strings.Split(s, "\n")
	res := []string{}
	start := false
	for i := 0; i < len(l); i++ {
		v := l[i]
		if strings.Contains(v, fmt.Sprintf("generated to unmarshal %s", structName)) {
			start = true
			continue
		} else if v == "}" && start {
			start = false
			continue
		}

		if start {
			continue
		}
		res = append(res, v)
	}
	return strings.Join(res, "\n")
}

func generate(structName string, structType *helper.Struct) string {
	return generateNewType(structName, structType.Fields) + "\n\n" + generateUnmarshal(structName, structType.Fields)
}

func generateNewType(structName string, fields []*helper.Field) string {
	res := []string{}
	res = append(res, "// unmarshal"+structName+" generated to unmarshal "+structName+" iface")
	res = append(res, "type unmarshal"+structName+" struct {")
	for _, v := range fields {
		if v.UnmarshalTool != "" {
			typ := "json.RawMessage"
			if strings.HasPrefix(v.RealType, "[]") {
				typ = "[]json.RawMessage"
			} else if strings.HasPrefix(v.RealType, "map[string][]") {
				typ = "map[string][]json.RawMessage"
			}
			res = append(res, fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`", v.Name, typ, v.JSONTag))
		} else {
			res = append(res, fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`", v.Name, v.NameType, v.JSONTag))
		}
	}
	res = append(res, "}")
	return strings.Join(res, "\n")
}

func generateUnmarshal(structName string, fields []*helper.Field) string {
	res := []string{}
	res = append(res, "\n")
	res = append(res, fmt.Sprintf(`// UnmarshalJSON generated to unmarshal %s iface`, structName))
	res = append(res, fmt.Sprintf(`func (r *%s) UnmarshalJSON(bs []byte) error {`, structName))
	res = append(res, fmt.Sprintf("\tobj := new(unmarshal%s)", structName))
	res = append(res, "\terr := json.Unmarshal(bs, obj)")
	res = append(res, "\tif err != nil {")
	res = append(res, "\t\treturn err")
	res = append(res, "\t}")
	for _, v := range fields {
		if v.UnmarshalTool == "" {
			res = append(res, fmt.Sprintf("\tr.%s = obj.%s", v.Name, v.Name))
		} else {
			if strings.HasPrefix(v.RealType, "[]") {
				res = append(res, fmt.Sprintf("\tr.%s = make(%s, len(obj.%s))", v.Name, v.NameType, v.Name))
				res = append(res, fmt.Sprintf("\tfor i, v := range obj.%s {", v.Name))
				res = append(res, fmt.Sprintf("\t\tr.%s[i], err = %s(v)", v.Name, v.UnmarshalTool))
				res = append(res, "\t\tif err != nil {")
				res = append(res, "\t\t\treturn err")
				res = append(res, "\t\t}")
				res = append(res, "\t}")
			} else if strings.HasPrefix(v.RealType, "map[string][]") {
				res = append(res, fmt.Sprintf("\tr.%s = make(%s, len(obj.%s))", v.Name, v.NameType, v.Name))
				res = append(res, fmt.Sprintf("\tfor k, v := range obj.%s {", v.Name))
				res = append(res, fmt.Sprintf("\t\tfor _, vv := range v {"))
				res = append(res, fmt.Sprintf("\t\t\ttmp, err := %s(vv)", v.UnmarshalTool))
				res = append(res, "\t\t\tif err != nil {")
				res = append(res, "\t\t\t\treturn err")
				res = append(res, "\t\t\t}")
				res = append(res, fmt.Sprintf("\t\t\tr.%s[k] = append(r.%s[k], tmp)", v.Name, v.Name))
				res = append(res, "\t\t}")
				res = append(res, "\t}")
			} else {
				spew.Dump(v.Name, v.RealType, v.NameType)
				panic(v)
			}
		}
	}
	res = append(res, "\treturn nil")
	res = append(res, "}")
	return strings.Join(res, "\n")
}
