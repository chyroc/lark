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
	need := `[["text",123.456,{"link":"https://github.com/","text":"text","type":"url"},"a@b.com",{"grantReadPermission":true,"notify":false,"text":"a@b.com","text_type":"email","type":"mention"}],[{"text":"=A1","type":"formula"},{"objType":"doc","text":"doc-token","textType":"fileToken","type":"mention"},{"type":"multipleValue","values":["1",123]}]]`
	as.Equal(need, string(bs))
}
