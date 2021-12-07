package card

import (
	"github.com/chyroc/lark"
)

func Overflow(options ...*lark.MessageContentCardObjectOption) *lark.MessageContentCardElementOverflow {
	return &lark.MessageContentCardElementOverflow{
		Options: options,
	}
}
