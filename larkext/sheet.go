package larkext

import (
	"context"

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

func (r *Sheet) Meta(ctx context.Context) (*lark.GetSheetMetaResp, error) {
	return r.meta(ctx)
}

func (r *Sheet) SheetToken() string {
	return r.sheetToken
}

func (r *Sheet) Delete(ctx context.Context) error {
	return r.delete(ctx)
}

func (r *Sheet) SetTitle(ctx context.Context, title string) error {
	return r.setTitle(ctx, title)
}

func (r *Sheet) NewSheet(ctx context.Context, title string) (string, error) {
	return r.newSheet(ctx, title)
}

func (r *Sheet) DeleteSheet(ctx context.Context, sheetID string) error {
	return r.deleteSheet(ctx, sheetID)
}

func (r *Sheet) CopySheet(ctx context.Context, fromSheetID string, toTitle *string) (string, error) {
	return r.copySheet(ctx, fromSheetID, toTitle)
}

func (r *Sheet) SetActiveSheet(sheetID string) {
	r.activeSheetID = sheetID
}

func (r *Sheet) GetActiveSheet() string {
	return r.activeSheetID
}

func (r *Sheet) SetSheetName(ctx context.Context, sheetID string, name string) error {
	return r.setSheetName(ctx, sheetID, name)
}

func (r *Sheet) InsertCols(ctx context.Context, sheetID string, startIndex int, count int) error {
	return r.insertCols(ctx, sheetID, startIndex, count)
}

func (r *Sheet) InsertRows(ctx context.Context, sheetID string, startIndex int, count int) error {
	return r.insertRows(ctx, sheetID, startIndex, count)
}

type Cols struct {
	sheet   *Sheet
	sheetID string
}

func (r *Sheet) Cols(sheetID string) *Cols {
	return &Cols{sheet: r, sheetID: sheetID}
}

type Rows struct {
	sheet   *Sheet
	sheetID string
}

func (r *Sheet) Rows(sheetID string) *Rows {
	return &Rows{sheet: r, sheetID: sheetID}
}

func (r *Sheet) SearchSheet(ctx context.Context, sheetID, value string, condition *lark.FindSheetReqFindCondition) (*lark.FindSheetRespFindResult, error) {
	return r.searchSheet(ctx, sheetID, value, condition)
}

func (r *Sheet) MoveRows(ctx context.Context, sheetID string, fromStartIndex, count, destIndex int) error {
	return r.moveRows(ctx, sheetID, fromStartIndex, count, destIndex)
}
