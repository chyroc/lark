package card

import (
	"github.com/chyroc/lark"
)

func URL(url string) *lark.MessageContentCardObjectURL {
	return &lark.MessageContentCardObjectURL{
		URL: url,
	}
}
