package test

import (
	"encoding/json"
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_SheetContent(t *testing.T) {
	as := assert.New(t)

	content := [][]lark.SheetContent{
		{
			&lark.SheetContentString{
				Text: "text",
			},
			&lark.SheetContentNumber{
				Number: "123.456",
			},
			&lark.SheetContentLink{
				Text: "text",
				Link: "https://github.com/",
			},
			&lark.SheetContentEmail{
				Email: "a@b.com",
			},
			&lark.SheetContentUser{
				TextType:            "email",
				Text:                "a@b.com",
				Notify:              false,
				GrantReadPermission: true,
			},
		},
		{
			&lark.SheetContentFormula{
				Text: "=A1",
			},
			&lark.SheetContentDoc{
				Text:    "doc-token",
				ObjType: "doc",
			},
			&lark.SheetContentMultipleValue{
				Values: []interface{}{"1", 123},
			},
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
