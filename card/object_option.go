package card

import (
	"github.com/chyroc/lark"
)

func Option(value string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Text:     Text(value),
		Value:    value,
		URL:      "",
		MultiURL: nil,
	}
}
