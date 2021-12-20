package larkext

import (
	"github.com/chyroc/lark"
)

// BitableTable is a table that can be used to store bitable data.
type BitableTable struct {
	larkClient *lark.Lark
	appToken   string
	tableID    string
}

// NewBitableTable new a bitable table.
func NewBitableTable(larkClient *lark.Lark, appToken string, tableID string) *BitableTable {
	return &BitableTable{larkClient: larkClient, appToken: appToken, tableID: tableID}
}
