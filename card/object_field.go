package card

import (
	"github.com/chyroc/lark"
)

func Field()*lark.MessageContentCardObjectField  {
	return &lark.MessageContentCardObjectField{
		IsShort: false,
		Text:    nil,
	}
}
