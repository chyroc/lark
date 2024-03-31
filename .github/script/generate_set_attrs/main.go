package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	struct_, ok := ast.Structs[structName]
	if !ok {
		return fmt.Errorf("not found struct %s", structName)
	}

	content := clearSetAttr(string(bs), structName)
	attrs := genSetAttrs(structName, struct_, ast)

	return ioutil.WriteFile(path, []byte(helper.ClearString(content+attrs, "\n\n\n", "\n\n")), 0o644)
}

func clearSetAttr(s, structName string) string {
	l := strings.Split(s, "\n")
	// SetTextTagList set ComponentHeader.TextTagList attribute
	res := []string{}
	start := false
	for i := 0; i < len(l); i++ {

		v := l[i]
		if strings.HasPrefix(v, "// Set") && strings.Contains(v, "set "+structName) {
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

func genSetAttrs(structName string, structType *helper.Struct, ast *helper.ASTResult) string {
	res := []string{}
	for _, v := range structType.Fields {
		res = append(res, genSetAttr(structName, v, ast))
	}
	return strings.Join(res, "")
}

func genSetAttr(structName string, field *helper.Field, ast *helper.ASTResult) string {
	consts := ast.Consts[field.NameType]
	nameType := field.NameType
	if strings.HasPrefix(nameType, "[]") {
		nameType = "..." + nameType[2:]
	}
	res := []string{
		fmt.Sprintf(`
// Set%s set %s.%s attribute
func (r *%s) Set%s(val %s) *%s {
	r.%s = val
	return r
}
`, field.Name, structName, field.Name, structName, field.Name, nameType, structName, field.Name),
	}
	for _, const_ := range unique(consts) {
		constSuffix := const_.NameType[len(field.NameType):]
		method := field.Name + constSuffix
		res = append(res, fmt.Sprintf(`
// Set%s set %s.%s attribute to %s
func (r *%s) Set%s() *%s {
	r.%s = %s
	return r
}`, method, structName, field.Name, const_.NameType, structName, method, structName, field.Name, const_.NameType))
	}
	return strings.Join(res, "\n")
}

func unique(l []helper.Const) []helper.Const {
	res := []helper.Const{}
	done := map[string]bool{}
	for _, v := range l {
		if !done[v.NameType] {
			res = append(res, v)
			done[v.NameType] = true
		}
	}
	return res
}
