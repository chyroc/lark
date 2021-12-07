package card

import (
	"github.com/chyroc/lark"
)

func Confirm(title, text string) *lark.MessageContentCardObjectConfirm {
	return &lark.MessageContentCardObjectConfirm{
		Title: Text(title),
		Text:  Text(text),
	}
}
