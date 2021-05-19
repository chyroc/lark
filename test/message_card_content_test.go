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
	s := string(bs)
	as.Contains(s, `{"actions":[{`)
	as.Contains(s, `"tag":"button"`)
	as.Contains(s, `"text":{`)
	as.Contains(s, `"tag":"lark_md"`)
	as.Contains(s, `"content":"click-me"`)
	as.Contains(s, `"lines":1`)
	as.Contains(s, `"type":"default"`)
	as.Contains(s, `"url":"https://www.baidu.com"`)
	as.Contains(s, `"tag":"action"`)
}
