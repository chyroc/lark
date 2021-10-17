package larkext

import (
	"github.com/chyroc/lark"
)

type Doc struct {
	larkClient *lark.Lark
	docToken   string
}

func NewDoc(larkClient *lark.Lark, docToken string) *Doc {
	r := new(Doc)
	r.larkClient = larkClient
	r.docToken = docToken
	return r
}
