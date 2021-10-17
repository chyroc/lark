package larkext

import (
	"github.com/chyroc/lark"
)

type Drive struct {
	larkClient *lark.Lark
}

func NewDrive(larkClient *lark.Lark) *Drive {
	r := new(Drive)
	r.larkClient = larkClient

	return r
}
