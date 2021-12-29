/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package larkext

import (
	"context"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
)

// Sheet is sheet client
type Sheet struct {
	larkClient    *lark.Lark
	sheetToken    string
	activeSheetID string
}

// NewSheet new sheet client
func NewSheet(larkClient *lark.Lark, sheetToken string) *Sheet {
	r := new(Sheet)
	r.larkClient = larkClient
	r.sheetToken = sheetToken

	return r
}

// Meta get sheet meta
func (r *Sheet) Meta(ctx context.Context) (*lark.GetSheetMetaResp, error) {
	return r.meta(ctx)
}

// SheetToken get sheet token
func (r *Sheet) SheetToken() string {
	return r.sheetToken
}

// Delete delete sheet
func (r *Sheet) Delete(ctx context.Context) error {
	return r.delete(ctx)
}

// SetTitle set sheet title
func (r *Sheet) SetTitle(ctx context.Context, title string) error {
	return r.setTitle(ctx, title)
}

// NewSheet create new sheet
func (r *Sheet) NewSheet(ctx context.Context, title string) (string, error) {
	return r.newSheet(ctx, title)
}

// DeleteSheet delete sheet
func (r *Sheet) DeleteSheet(ctx context.Context, sheetID string) error {
	return r.deleteSheet(ctx, sheetID)
}

// CopySheet copy sheet
func (r *Sheet) CopySheet(ctx context.Context, fromSheetID string, toTitle *string) (string, error) {
	return r.copySheet(ctx, fromSheetID, toTitle)
}

// SetSheetTitle set sheet title
func (r *Sheet) SetSheetTitle(ctx context.Context, sheetID string, title string) error {
	return r.setSheetTitle(ctx, sheetID, title)
}

// SetSheetIndex set sheet index
func (r *Sheet) SetSheetIndex(ctx context.Context, sheetID string, index int64) error {
	return r.setSheetIndex(ctx, sheetID, index)
}

// SetSheetTitle hide sheet
func (r *Sheet) HideSheet(ctx context.Context, sheetID string, hidden bool) error {
	return r.hideSheet(ctx, sheetID, hidden)
}

// FrozenSheet freeze worksheet, 0 means unfreeze
func (r *Sheet) FrozenSheet(ctx context.Context, sheetID string, rowCount, colCount int64) error {
	return r.frozenSheet(ctx, sheetID, rowCount, colCount)
}

// LockSheet lock worksheet
func (r *Sheet) LockSheet(ctx context.Context, sheetID string, lockInfo string, editableUserIDs []string) error {
	return r.lockSheet(ctx, sheetID, false, lockInfo, editableUserIDs)
}

// UnlockSheet unlock worksheet
func (r *Sheet) UnlockSheet(ctx context.Context, sheetID string) error {
	return r.lockSheet(ctx, sheetID, true, "", nil)
}

// MoveRows 移动行，将 sheetID 表中，从 startIndex 行后，数量为 count 的行，移动 diff 行，diff 小于 0 表示上移，diff 大于 0 ，表示下移
func (r *Sheet) MoveRows(ctx context.Context, sheetID string, startIndex, count, diff int64) error {
	return r.moveDimension(ctx, "ROWS", sheetID, startIndex, count, diff)
}

// MoveCols 移动列，将 sheetID 表中，从 startIndex 列后，数量为 count 的列，移动 diff 列，diff 小于 0 表示左移，diff 大于 0 ，表示右移
func (r *Sheet) MoveCols(ctx context.Context, sheetID string, startIndex, count, diff int64) error {
	return r.moveDimension(ctx, "COLUMNS", sheetID, startIndex, count, diff)
}

// InsertRows 插入行，表示从 startIndex 开始插入数量为 count 的行
func (r *Sheet) InsertRows(ctx context.Context, sheetID string, startIndex, count int64) error {
	return r.insertDimension(ctx, "ROWS", sheetID, startIndex-1, startIndex-1+count)
}

// InsertCols 插入列，表示从 startIndex 开始插入数量为 count 的列
func (r *Sheet) InsertCols(ctx context.Context, sheetID string, startIndex, count int64) error {
	return r.insertDimension(ctx, "COLUMNS", sheetID, startIndex-1, startIndex-1+count)
}

// Append 追加数据
//
// option 是追加数据模式，默认 OVERWRITE，还可以传入 INSERT_ROWS
// 当 option == OVERWRITE 的时候，sheet 会找到 range 范围内左上角第一个空白的格子，然后将 values 写入该范围，如果除了「左上角第一个空白的格子」外有其他非空格子，这些格子会被覆盖。
// 当 option == INSERT_ROWS 的时候，sheet 会找到 range 范围内左上角第一个空白的格子，尝试然后将 values 写入该范围，如果除了「左上角第一个空白的格子」外有其他非空格子，会先将这些格子整体下移几格，知道空间够放下 values
func (r *Sheet) Append(ctx context.Context, cellRange string, values [][]lark.SheetContent, option *string) error {
	return r.appendDimension(ctx, cellRange, values, option)
}

// AddRows 增加行
func (r *Sheet) AddRows(ctx context.Context, sheetID string, count int64) error {
	return r.addDimension(ctx, "ROWS", sheetID, count)
}

// AddCols 增加列
func (r *Sheet) AddCols(ctx context.Context, sheetID string, count int64) error {
	return r.addDimension(ctx, "COLUMNS", sheetID, count)
}

// SetRowsVisible 设置行的可见性
func (r *Sheet) SetRowsVisible(ctx context.Context, sheetID string, startIndex, count int64, visible bool) error {
	return r.updateDimension(ctx, "ROWS", sheetID, startIndex, count, &visible, nil)
}

// SetColsVisible 设置列的可见性
func (r *Sheet) SetColsVisible(ctx context.Context, sheetID string, startIndex, count int64, visible bool) error {
	return r.updateDimension(ctx, "COLUMNS", sheetID, startIndex, count, &visible, nil)
}

// SetRowsSize 设置行的高度
func (r *Sheet) SetRowsSize(ctx context.Context, sheetID string, startIndex, count, size int64) error {
	return r.updateDimension(ctx, "ROWS", sheetID, startIndex, count, nil, ptr.Int64(size))
}

// SetColsSize 设置列的宽度
func (r *Sheet) SetColsSize(ctx context.Context, sheetID string, startIndex, count, size int64) error {
	return r.updateDimension(ctx, "COLUMNS", sheetID, startIndex, count, nil, ptr.Int64(size))
}

// DeleteRows 删除行
func (r *Sheet) DeleteRows(ctx context.Context, sheetID string, startIndex, count int64) error {
	return r.deleteDimension(ctx, "ROWS", sheetID, startIndex, count)
}

// DeleteCols 删除列
func (r *Sheet) DeleteCols(ctx context.Context, sheetID string, startIndex, count int64) error {
	return r.deleteDimension(ctx, "COLUMNS", sheetID, startIndex, count)
}

// // 获取单元格内容
// func (r *Sheet) Get(ctx context.Context, cellRange string, option *lark.GetSheetValueReq) error {
// 	return r.getValue(ctx, cellRange, option)
// }

// SetCellStyle 设置单元格样式
func (r *Sheet) SetCellStyle(ctx context.Context, cellRange string, style *lark.SetSheetStyleReqAppendStyleStyle) error {
	return r.setCellStyle(ctx, cellRange, style)
}

// CleanCellStyle 清除单元格样式
func (r *Sheet) CleanCellStyle(ctx context.Context, cellRange string) error {
	return r.setCellStyle(ctx, cellRange, &lark.SetSheetStyleReqAppendStyleStyle{Clean: ptr.Bool(true)})
}

// BatchSetCellStyle 批量设置单元格样式
func (r *Sheet) BatchSetCellStyle(ctx context.Context, styles []*lark.BatchSetSheetStyleReqData) error {
	return r.batchSetCellStyle(ctx, styles)
}

// BatchCleanCellStyle 批量清除单元格样式
func (r *Sheet) BatchCleanCellStyle(ctx context.Context, cellRanges []string) error {
	styles := []*lark.BatchSetSheetStyleReqData{{
		Ranges: cellRanges,
		Style: &lark.BatchSetSheetStyleReqDataStyle{
			Clean: ptr.Bool(true),
		},
	}}
	return r.batchSetCellStyle(ctx, styles)
}

// TODO：批量 style

// MergeCell 合并单元格
//
// mergeType 支持三种格式
// 当 mergeType == MERGE_ALL，合并所选区域
// 当 mergeType == MERGE_ROWS，按行合并
// 当 mergeType == MERGE_COLUMNS，按列合并
func (r *Sheet) MergeCell(ctx context.Context, cellRange, mergeType string) error {
	return r.mergeCell(ctx, cellRange, mergeType)
}

// UnmergeCell 取消合并单元格
func (r *Sheet) UnmergeCell(ctx context.Context, cellRange string) error {
	return r.unmergeCell(ctx, cellRange)
}

// SetSheetValue 将内容写入单元格
func (r *Sheet) SetSheetValue(ctx context.Context, cellRange string, contents [][]lark.SheetContent) error {
	return r.setSheetValue(ctx, cellRange, contents)
}

// BatchSetSheetValue 批量将内容写入单元格
func (r *Sheet) BatchSetSheetValue(ctx context.Context, values []*lark.BatchSetSheetValueReqValueRange) error {
	return r.batchSetSheetValue(ctx, values)
}

// SetSheetValueImage 将图片写入单元格
func (r *Sheet) SetSheetValueImage(ctx context.Context, cellRange string, image []byte) error {
	return r.setSheetValueImage(ctx, cellRange, image)
}

// Search 搜索
func (r *Sheet) Search(ctx context.Context, sheetID, value string, condition *lark.FindSheetReqFindCondition) (*lark.FindSheetRespFindResult, error) {
	return r.search(ctx, sheetID, value, condition)
}

// Replace 替换
func (r *Sheet) Replace(ctx context.Context, sheetID, old, new string, condition *lark.ReplaceSheetReqFindCondition) (*lark.ReplaceSheetRespReplaceResult, error) {
	return r.replace(ctx, sheetID, old, new, condition)
}
