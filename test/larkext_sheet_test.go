package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_SheetExt(t *testing.T) {
	as := assert.New(t)
	imageBytes := readFile("./test/file_1.png")
	ctx := context.Background()
	larkCli := AppAllPermission.Ins()
	sheetClient := testCreateSheet(larkCli)
	defer func() {
		as.Nil(sheetClient.Delete(ctx))
	}()

	meta, err := sheetClient.Meta(ctx)
	as.Nil(err)
	defaultSheetID := meta.Sheets[0].SheetID
	as.Equal(sheetClient.SheetToken(), meta.SpreadSheetToken)

	t.Run("set title", func(t *testing.T) {
		as.Nil(sheetClient.SetTitle(ctx, fmt.Sprintf("sheet title %d", randInt64())))
	})

	t.Run("set sheet title", func(t *testing.T) {
		as.Nil(sheetClient.SetSheetTitle(ctx, defaultSheetID, fmt.Sprintf("sheet title %d", randInt64())))
	})

	t.Run("new-sheet", func(t *testing.T) {
		sheetID, err := sheetClient.NewSheet(ctx, fmt.Sprintf("sheet %d", randInt64()))
		as.Nil(err)
		as.NotEmpty(sheetID)

		as.Nil(sheetClient.DeleteSheet(ctx, sheetID))
	})

	t.Run("copy-sheet", func(t *testing.T) {
		sheetID, err := sheetClient.CopySheet(ctx, defaultSheetID, ptr.String(randInt64String()))
		as.Nil(err)
		as.NotEmpty(sheetID)

		as.Nil(sheetClient.DeleteSheet(ctx, sheetID))
	})

	t.Run("", func(t *testing.T) {
		err = sheetClient.SetSheetValue(ctx, larkext.CellRange(defaultSheetID, 1, 1, 1, 7), [][]lark.SheetContent{
			{{String: ptr.String("1")}}, // 1
			{{Int: ptr.Int64(100)}},     // 2
			{{Link: &lark.SheetValueLink{Text: "link", Link: "https://google.com"}}}, // 3
			{{Formula: &lark.SheetValueFormula{Text: "=A1"}}},                        // 4
			{{String: ptr.String("5")}},                                              // 5
			{{String: ptr.String("6")}},                                              // 6
			{{String: ptr.String("7")}},                                              // 7
		})
		as.Nil(err)

		err = sheetClient.BatchSetSheetValue(ctx, []*lark.BatchSetSheetValueReqValueRange{
			{
				Range: larkext.CellRange(defaultSheetID, 1, 1, 1, 1),
				Values: [][]lark.SheetContent{
					{
						{String: ptr.String("(1,1) 字符串")},
					},
				},
			},
			{
				Range: larkext.CellRange(defaultSheetID, 2, 2, 2, 2),
				Values: [][]lark.SheetContent{
					{
						{Link: &lark.SheetValueLink{Text: "Google", Link: "https://google.com"}},
					},
				},
			},
			{
				Range: larkext.CellRange(defaultSheetID, 3, 3, 3, 3),
				Values: [][]lark.SheetContent{
					{
						{AtDoc: &lark.SheetValueAtDoc{ObjType: "sheet", Text: sheetClient.SheetToken()}},
					},
				},
			},
		})
		as.Nil(err)

		err = sheetClient.SetSheetValueImage(ctx, larkext.CellRange(defaultSheetID, 3, 3, 3, 3), imageBytes)
		as.Nil(err)

		// 1 2 行，往下移动 2 行
		as.Nil(sheetClient.MoveRows(ctx, defaultSheetID, 1, 2, 2))

		// 1 2 列，往右移动 2 列
		as.Nil(sheetClient.MoveCols(ctx, defaultSheetID, 1, 2, 2))

		as.Nil(sheetClient.InsertRows(ctx, defaultSheetID, 1, 1))

		as.Nil(sheetClient.InsertCols(ctx, defaultSheetID, 1, 1))

		as.Nil(sheetClient.AddRows(ctx, defaultSheetID, 1))

		as.Nil(sheetClient.AddCols(ctx, defaultSheetID, 1))

		err = sheetClient.Append(ctx, larkext.CellRange(defaultSheetID, 1, 1, 1, 7), [][]lark.SheetContent{
			{{String: ptr.String("1")}}, // 1
			{{Int: ptr.Int64(100)}},     // 2
			{{Link: &lark.SheetValueLink{Text: "link", Link: "https://google.com"}}}, // 3
			{{Formula: &lark.SheetValueFormula{Text: "=A1"}}},                        // 4
			{{String: ptr.String("5")}},                                              // 5
			{{String: ptr.String("6")}},                                              // 6
			{{String: ptr.String("7")}},                                              // 7
		}, nil)
		as.Nil(err)

		as.Nil(sheetClient.SetRowsVisible(ctx, defaultSheetID, 1, 1, false))

		as.Nil(sheetClient.SetColsVisible(ctx, defaultSheetID, 1, 1, false))

		as.Nil(sheetClient.SetRowsSize(ctx, defaultSheetID, 2, 2, 50))

		as.Nil(sheetClient.SetColsSize(ctx, defaultSheetID, 2, 2, 200))

		as.Nil(sheetClient.DeleteRows(ctx, defaultSheetID, 10, 1))

		as.Nil(sheetClient.DeleteCols(ctx, defaultSheetID, 10, 1))

		as.Nil(sheetClient.SetCellStyle(ctx, larkext.CellRange(defaultSheetID, 1, 1, 20, 20), &lark.SetSheetStyleReqAppendStyleStyle{
			BackColor: ptr.String("#21d11f"),
		}))

		as.Nil(sheetClient.CleanCellStyle(ctx, larkext.CellRange(defaultSheetID, 1, 1, 1, 1)))

		as.Nil(sheetClient.MergeCell(ctx, larkext.CellRange(defaultSheetID, 1, 1, 2, 1), ""))

		as.Nil(sheetClient.UnmergeCell(ctx, larkext.CellRange(defaultSheetID, 1, 1, 2, 1)))

		_, err = sheetClient.Search(ctx, defaultSheetID, "6", nil)
		as.Nil(err)
		_, err = sheetClient.Replace(ctx, defaultSheetID, "6", "7", nil)
		as.Nil(err)
	})

	err = sheetClient.InsertRows(ctx, defaultSheetID, 1, 1)
	as.Nil(err)
	err = sheetClient.InsertCols(ctx, defaultSheetID, 1, 1)
	as.Nil(err)
}

func testCreateSheet(larkCli *lark.Lark) *larkext.Sheet {
	folderIns := larkext.NewFolder(larkCli, "")

	sheetClient, err := folderIns.NewSheet(ctx, "")
	if err != nil {
		panic(err)
	}

	_, _, err = larkCli.Drive.UpdateDriveMemberPermission(context.Background(), &lark.UpdateDriveMemberPermissionReq{
		NeedNotification: ptr.Bool(true),
		Type:             "sheet",
		Token:            sheetClient.SheetToken(),
		MemberID:         UserAdmin.OpenID,
		MemberType:       "openid",
		Perm:             "full_access",
	})
	if err != nil {
		panic(err)
	}

	return sheetClient
}
