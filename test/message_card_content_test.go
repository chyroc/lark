package test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_CardContent(t *testing.T) {
	as := assert.New(t)

	action := lark.MessageContentCardModuleAction{
		Actions: []*lark.MessageContentCardElementButton{
			{
				Text: &lark.MessageContentCardObjectText{
					Tag:     "lark_md",
					Content: "click-me",
					Lines:   1,
				},
				URL:  "https://www.baidu.com",
				Type: "default",
			},
		},
		Layout: "bisected",
	}
	bs, err := action.MarshalJSON()
	as.Nil(err)
	need := `{"actions":[{"tag":"button","text":{"tag":"lark_md","content":"click-me","lines":1},"type":"default","url":"https://www.baidu.com"}],"layout":"bisected","tag":"action"}`
	as.Equal(need, string(bs))
}
