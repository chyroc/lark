package card

import (
	"github.com/chyroc/lark"
)

func Header(title string) *lark.MessageContentCardHeader {
	return &lark.MessageContentCardHeader{
		Title:    Text(title),
		Template: "",
	}
}
