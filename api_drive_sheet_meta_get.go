// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetSheetMeta 该接口用于根据 spreadsheetToken 获取表格元数据。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN
func (r *DriveService) GetSheetMeta(ctx context.Context, request *GetSheetMetaReq, options ...MethodOptionFunc) (*GetSheetMetaResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetMeta != nil {
		r.cli.logDebug(ctx, "[lark] Drive#GetSheetMeta mock enable")
		return r.cli.mock.mockDriveGetSheetMeta(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#GetSheetMeta call api")
	r.cli.logDebug(ctx, "[lark] Drive#GetSheetMeta request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#GetSheetMeta GET https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#GetSheetMeta GET https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "GetSheetMeta", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#GetSheetMeta request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveGetSheetMeta(f func(ctx context.Context, request *GetSheetMetaReq, options ...MethodOptionFunc) (*GetSheetMetaResp, *Response, error)) {
	r.mockDriveGetSheetMeta = f
}

func (r *Mock) UnMockDriveGetSheetMeta() {
	r.mockDriveGetSheetMeta = nil
}

type GetSheetMetaReq struct {
	ExtFields        *string `query:"extFields" json:"-"`       // 额外返回的字段，extFields=protectedRange时返回保护行列信息
	SpreadsheetToken string  `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token；获取方式见[对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
}

type getSheetMetaResp struct {
	Code int               `json:"code,omitempty"`
	Msg  string            `json:"msg,omitempty"`
	Data *GetSheetMetaResp `json:"data,omitempty"`
}

type GetSheetMetaResp struct {
	SpreadsheetToken string                      `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	Properties       *GetSheetMetaRespProperties `json:"properties,omitempty"`       // spreadsheet 的属性
	Sheets           *GetSheetMetaRespSheets     `json:"sheets,omitempty"`           // spreadsheet 下的sheet列表
}

type GetSheetMetaRespProperties struct {
	Title      string `json:"title,omitempty"`      // spreadsheet 的标题
	OwnerUser  int    `json:"ownerUser,omitempty"`  // 所有者的 id
	SheetCount int    `json:"sheetCount,omitempty"` // spreadsheet 下的 sheet 数
	Revision   int    `json:"revision,omitempty"`   // 该 sheet 的版本
}

type GetSheetMetaRespSheets struct {
	SheetID        string                                `json:"sheetId,omitempty"`        // sheet 的 id
	Title          string                                `json:"title,omitempty"`          // sheet 的标题
	Index          int                                   `json:"index,omitempty"`          // sheet 的位置
	RowCount       int                                   `json:"rowCount,omitempty"`       // sheet 的最大行数
	ColumnCount    int                                   `json:"columnCount,omitempty"`    // sheet 的最大列数
	FrozenRowCount int                                   `json:"frozenRowCount,omitempty"` // 该 sheet 的冻结行数，小于等于 sheet 的最大行数，0表示未设置冻结
	FrozenColCount int                                   `json:"frozenColCount,omitempty"` // 该 sheet 的冻结列数，小于等于 sheet 的最大列数，0表示未设置冻结
	Merges         *GetSheetMetaRespSheetsMerges         `json:"merges,omitempty"`         // 该 sheet 中合并单元格的范围
	ProtectedRange *GetSheetMetaRespSheetsProtectedRange `json:"protectedRange,omitempty"` // 该 sheet 中保护范围
}

type GetSheetMetaRespSheetsMerges struct {
	StartRowIndex    int `json:"startRowIndex,omitempty"`    // 合并单元格范围的开始行下标，index 从 0 开始
	StartColumnIndex int `json:"startColumnIndex,omitempty"` // 合并单元格范围的开始列下标，index 从 0 开始
	RowCount         int `json:"rowCount,omitempty"`         // 合并单元格范围的行数量
	ColumnCount      int `json:"columnCount,omitempty"`      // 合并单元格范围的列数量
}

type GetSheetMetaRespSheetsProtectedRange struct {
	Dimension *GetSheetMetaRespSheetsProtectedRangeDimension `json:"dimension,omitempty"` // 保护行列的信息，如果为保护工作表，则该字段为空
	ProtectID string                                         `json:"protectId,omitempty"` // 保护范围ID
	LockInfo  string                                         `json:"lockInfo,omitempty"`  // 保护说明
	SheetID   string                                         `json:"sheetId,omitempty"`   // 保护工作表 ID
}

type GetSheetMetaRespSheetsProtectedRangeDimension struct {
	StartIndex     int    `json:"startIndex,omitempty"`     // 保护行列的起始位置，位置从1开始
	EndIndex       int    `json:"endIndex,omitempty"`       // 保护行列的结束位置，位置从1开始
	MajorDimension string `json:"majorDimension,omitempty"` // 若为ROWS，则为保护行；为COLUMNS，则为保护列
	SheetID        string `json:"sheetId,omitempty"`        // 保护范围所在工作表 ID
}
