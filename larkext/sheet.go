package larkext

import (
	"context"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
)

type Sheet struct {
	larkClient    *lark.Lark
	sheetToken    string
	activeSheetID string
}

func NewSheet(larkClient *lark.Lark, sheetToken string) *Sheet {
	r := new(Sheet)
	r.larkClient = larkClient
	r.sheetToken = sheetToken

	return r
}

// Meta 获取 sheet meta 信息
func (r *Sheet) Meta(ctx context.Context) (*lark.GetSheetMetaResp, error) {
	return r.meta(ctx)
}

func (r *Sheet) SheetToken() string {
	return r.sheetToken
}

// Delete 删除自己
func (r *Sheet) Delete(ctx context.Context) error {
	return r.delete(ctx)
}

// SetTitle 设置标题
func (r *Sheet) SetTitle(ctx context.Context, title string) error {
	return r.setTitle(ctx, title)
}

// NewSheet 创建新的工作表
func (r *Sheet) NewSheet(ctx context.Context, title string) (string, error) {
	return r.newSheet(ctx, title)
}

// DeleteSheet 删除工作表
func (r *Sheet) DeleteSheet(ctx context.Context, sheetID string) error {
	return r.deleteSheet(ctx, sheetID)
}

// CopySheet 复制工作表
func (r *Sheet) CopySheet(ctx context.Context, fromSheetID string, toTitle *string) (string, error) {
	return r.copySheet(ctx, fromSheetID, toTitle)
}

// SetSheetName 设置工作表的名字
func (r *Sheet) SetSheetName(ctx context.Context, sheetID string, name string) error {
	return r.setSheetName(ctx, sheetID, name)
}

// MoveRows 移动行，将 sheetID 表中，从 startIndex 行后，数量为 count 的行，移动 diff 行，diff 小于 0 表示上移，diff 大于 0 ，表示下移
func (r *Sheet) MoveRows(ctx context.Context, sheetID string, startIndex, count, diff int) error {
	return r.moveDimension(ctx, "ROWS", sheetID, startIndex, count, diff)
}

// MoveCols 移动列，将 sheetID 表中，从 startIndex 列后，数量为 count 的列，移动 diff 列，diff 小于 0 表示左移，diff 大于 0 ，表示右移
func (r *Sheet) MoveCols(ctx context.Context, sheetID string, startIndex, count, diff int) error {
	return r.moveDimension(ctx, "COLUMNS", sheetID, startIndex, count, diff)
}

// InsertRows 插入行，表示从 startIndex 开始插入数量为 count 的行
func (r *Sheet) InsertRows(ctx context.Context, sheetID string, startIndex int, count int) error {
	return r.insertDimension(ctx, "ROWS", sheetID, startIndex-1, startIndex-1+count)
}

// InsertCols 插入列，表示从 startIndex 开始插入数量为 count 的列
func (r *Sheet) InsertCols(ctx context.Context, sheetID string, startIndex int, count int) error {
	return r.insertDimension(ctx, "COLUMNS", sheetID, startIndex-1, startIndex-1+count)
}

// Append 追加数据
//
// option 是追加数据模式，默认 OVERWRITE，还可以传入 INSERT_ROWS
// 当 option == OVERWRITE 的时候，sheet 会找到 range 范围内左上角第一个空白的格子，然后将 values 写入该范围，如果除了「左上角第一个空白的格子」外有其他非空格子，这些格子会被覆盖。
// 当 option == INSERT_ROWS 的时候，sheet 会找到 range 范围内左上角第一个空白的格子，尝试然后将 values 写入该范围，如果除了「左上角第一个空白的格子」外有其他非空格子，会先将这些格子整体下移几格，知道空间够放下 values
func (r *Sheet) Append(ctx context.Context, sheetID, ranges string, values [][]lark.SheetContent, option *string) error {
	return r.appendDimension(ctx, sheetID, ranges, values, option)
}

// 增加行
func (r *Sheet) AddRows(ctx context.Context, sheetID string, count int) error {
	return r.addDimension(ctx, "ROWS", sheetID, count)
}

// 增加列
func (r *Sheet) AddCols(ctx context.Context, sheetID string, count int) error {
	return r.addDimension(ctx, "COLUMNS", sheetID, count)
}

// 设置行的可见性
func (r *Sheet) SetRowsVisible(ctx context.Context, sheetID string, startIndex int, count int, visible bool) error {
	return r.updateDimension(ctx, "ROWS", sheetID, startIndex, count, &visible, nil)
}

// 设置列的可见性
func (r *Sheet) SetColsVisible(ctx context.Context, sheetID string, startIndex int, count int, visible bool) error {
	return r.updateDimension(ctx, "COLUMNS", sheetID, startIndex, count, &visible, nil)
}

// 设置行的高度
func (r *Sheet) SetRowsSize(ctx context.Context, sheetID string, startIndex int, count int, size int) error {
	return r.updateDimension(ctx, "ROWS", sheetID, startIndex, count, nil, ptr.Int64(int64(size)))
}

// 设置列的宽度
func (r *Sheet) SetColsSize(ctx context.Context, sheetID string, startIndex int, count int, size int) error {
	return r.updateDimension(ctx, "COLUMNS", sheetID, startIndex, count, nil, ptr.Int64(int64(size)))
}

// 删除行
func (r *Sheet) DeleteRows(ctx context.Context, sheetID string, startIndex int, count int) error {
	return r.deleteDimension(ctx, "ROWS", sheetID, startIndex, count)
}

// 删除列
func (r *Sheet) DeleteCols(ctx context.Context, sheetID string, startIndex int, count int) error {
	return r.deleteDimension(ctx, "COLUMNS", sheetID, startIndex, count)
}

// 删除列
func (r *Sheet) Get(ctx context.Context, sheetID string) error {
	return r.getValue(ctx, sheetID)
}

// 设置单元格样式
func (r *Sheet) SetCellStyle(ctx context.Context, sheetID, ranges string, style *lark.SetSheetStyleReqAppendStyleStyle) error {
	return r.setCellStyle(ctx, sheetID, ranges, style)
}

// 清除单元格样式
func (r *Sheet) CleanCellStyle(ctx context.Context, sheetID, ranges string) error {
	return r.setCellStyle(ctx, sheetID, ranges, &lark.SetSheetStyleReqAppendStyleStyle{Clean: ptr.Bool(true)})
}

// TODO：批量 style

// MergeCell 合并单元格
//
// mergeType 支持三种格式
// 当 mergeType == MERGE_ALL，合并所选区域
// 当 mergeType == MERGE_ROWS，按行合并
// 当 mergeType == MERGE_COLUMNS，按列合并
func (r *Sheet) MergeCell(ctx context.Context, ranges, mergeType string) error {
	return r.mergeCell(ctx, ranges, mergeType)
}

// UnmergeCell 取消合并单元格
func (r *Sheet) UnmergeCell(ctx context.Context, ranges string) error {
	return r.unmergeCell(ctx, ranges)
}

// SetSheetValueImage 将图片写入单元格
func (r *Sheet) SetSheetValueImage(ctx context.Context, ranges string, image []byte) error {
	return r.setSheetValueImage(ctx, ranges, image)
}

// 搜索
func (r *Sheet) Search(ctx context.Context, sheetID, value string, condition *lark.FindSheetReqFindCondition) (*lark.FindSheetRespFindResult, error) {
	return r.search(ctx, sheetID, value, condition)
}

// 替换
func (r *Sheet) Replace(ctx context.Context, sheetID, old, new string, condition *lark.ReplaceSheetReqFindCondition) (*lark.ReplaceSheetRespReplaceResult, error) {
	return r.replace(ctx, sheetID, old, new, condition)
}
