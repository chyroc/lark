package main

import (
	"context"

	"github.com/chyroc/lark"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	ctx := context.Background()
	appToken := "Plu9bzHFbaviJ2sljPwcj51bn8d"
	tableID := "tblXGxioGCS7ekUN"
	r := lark.New(lark.WithAppCredential("cli_a1ecf77320f8d00e", "DTrdblys6V1TKWYJNtq1HdmSwghIspQR"))

	res, _, err := r.Bitable.GetBitableRecordList(ctx, &lark.GetBitableRecordListReq{
		AppToken: appToken,
		TableID:  tableID,
		// ViewID:            nil,
		// Filter:            nil,
		// Sort:              nil,
		// FieldNames:        nil,
		// TextFieldAsArray:  nil,
		// UserIDType:        nil,
		// DisplayFormulaRef: nil,
		// AutomaticFields:   nil,
		// PageToken:         nil,
		// PageSize:          nil,
	})
	if err != nil {
		panic(err)
	}

	spew.Dump(res)
}
