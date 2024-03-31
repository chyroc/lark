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

	structType, ok := ast.Structs[structName]
	if !ok {
		return fmt.Errorf("not found struct %s", structName)
	}

	content := clearGenerate(string(bs), structName)
	attrs := generate(structName, structType.Fields)

	return ioutil.WriteFile(path, []byte(helper.ClearString(content+attrs, "\n\n\n", "\n\n")), 0o644)
}

func clearGenerate(s, structName string) string {
	l := strings.Split(s, "\n")
	res := []string{}
	start := false
	for i := 0; i < len(l); i++ {
		v := l[i]
		if strings.HasPrefix(v, "// toMap conv "+structName+" to map") {
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

func generate(structName string, fields []*helper.Field) string {
	fieldCount := 0
	for _, v := range fields {
		if v.JSONTag == "" || v.JSONTag == "-" {
			continue
		}
		fieldCount++
	}
	res := []string{}
	res = append(res, "\n")
	res = append(res, fmt.Sprintf(`// toMap conv %s to map`, structName))
	res = append(res, fmt.Sprintf(`func (r *%s) toMap() map[string]interface{} {`, structName))
	res = append(res, fmt.Sprintf("\tres := make(map[string]interface{}, %d)", fieldCount))
	for _, v := range fields {
		if v.JSONTag == "" || v.JSONTag == "-" {
			continue
		}
		if v.JSONOmitempty {
			if v.ZeroString() == "len" {
				res = append(res, fmt.Sprintf("\tif len(r.%s) != 0 {", v.Name))
			} else {
				res = append(res, fmt.Sprintf("\tif r.%s != %s {", v.Name, v.ZeroString()))
			}
			res = append(res, fmt.Sprintf("\t\tres[\"%s\"] = r.%s", v.JSONTag, v.Name))
			res = append(res, "\t}")
		} else {
			res = append(res, fmt.Sprintf("\tres[\"%s\"] = r.%s", v.JSONTag, v.Name))
		}
	}
	res = append(res, "\treturn res")
	res = append(res, "}")
	return strings.Join(res, "\n")
}
