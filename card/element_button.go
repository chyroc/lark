package card

import (
	"github.com/chyroc/lark"
)

// 回传数据按钮
func Button(text string, value interface{}) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Text:  Text(text),
		Type:  "default",
		Value: value,
	}
}

func LinkButton(text string, url string) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Text: Text(text),
		Type: "default",
		URL:  url,
	}
}
