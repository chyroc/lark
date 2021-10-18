package test

import (
	"encoding/json"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_SheetContent(t *testing.T) {
	as := assert.New(t)

	content := [][]lark.SheetContent{
		{
			{String: ptr.String("text")},
			{String: ptr.String("123.456")},
			{Link: &lark.SheetValueLink{Text: "text", Link: "https://github.com/"}},
			{String: ptr.String("a@b.com")},
			{AtUser: &lark.SheetValueAtUser{Text: "a@b.com", TextType: "email", Notify: false, GrantReadPermission: true}},
		},
		{
			{Formula: &lark.SheetValueFormula{Text: "=A1"}},
			{AtDoc: &lark.SheetValueAtDoc{ObjType: "doc", Text: "doc-token"}},
			{MultiValue: &lark.SheetValueMultiValue{Values: []interface{}{"1", 123}}},
		},
	}
	bs, err := json.Marshal(content)
	as.Nil(err)
	as.Contains(string(bs), `"text"`)
	as.Contains(string(bs), `123.456`)
	as.Contains(string(bs), `"link":"https://github.com/"`)
	as.Contains(string(bs), `"a@b.com"`)
	as.Contains(string(bs), `"grantReadPermission":true`)
	as.Contains(string(bs), `"text":"=A1"`)
	as.Contains(string(bs), `"values":["1",123]`)
}
