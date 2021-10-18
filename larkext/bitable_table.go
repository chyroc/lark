package larkext

import (
	"github.com/chyroc/lark"
)

type BitableTable struct {
	larkClient *lark.Lark
	appToken   string
	tableID    string
}

func NewBitableTable(larkClient *lark.Lark, appToken string, tableID string) *BitableTable {
	return &BitableTable{larkClient: larkClient, appToken: appToken, tableID: tableID}
}
