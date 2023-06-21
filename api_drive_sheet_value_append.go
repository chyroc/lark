// Code generated by lark_sdk_gen. DO NOT EDIT.
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

package lark

import (
	"context"
)

// AppendSheetValue 该接口用于在已有数据的末尾追加写入给定的数据。该接口会从给定的range中的起始行列开始向下寻找（如range为"sheet1!A1:B2", 将会依次寻找A1、A2、A3...）, 找到第一个空白位置后将数据写入到该区域。单次写入不得超过5000行, 100列, 每个格子不得超过5万字符。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMjMzUjLzIzM14yMyMTN
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/data-operation/append-data
func (r *DriveService) AppendSheetValue(ctx context.Context, request *AppendSheetValueReq, options ...MethodOptionFunc) (*AppendSheetValueResp, *Response, error) {
	if r.cli.mock.mockDriveAppendSheetValue != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#AppendSheetValue mock enable")
		return r.cli.mock.mockDriveAppendSheetValue(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "AppendSheetValue",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_append",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(appendSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveAppendSheetValue mock DriveAppendSheetValue method
func (r *Mock) MockDriveAppendSheetValue(f func(ctx context.Context, request *AppendSheetValueReq, options ...MethodOptionFunc) (*AppendSheetValueResp, *Response, error)) {
	r.mockDriveAppendSheetValue = f
}

// UnMockDriveAppendSheetValue un-mock DriveAppendSheetValue method
func (r *Mock) UnMockDriveAppendSheetValue() {
	r.mockDriveAppendSheetValue = nil
}

// AppendSheetValueReq ...
type AppendSheetValueReq struct {
	SpreadSheetToken string                         `path:"spreadsheetToken" json:"-"`  // spreadsheet 的 token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	InsertDataOption *string                        `query:"insertDataOption" json:"-"` // 遇到空行追加, 默认 OVERWRITE, 若空行的数量小于追加数据的行数, 则会覆盖已有数据；可选 INSERT_ROWS, 会在插入足够数量的行后再进行数据追加
	ValueRange       *AppendSheetValueReqValueRange `json:"valueRange,omitempty"`       // 值与范围
}

// AppendSheetValueReqValueRange ...
type AppendSheetValueReqValueRange struct {
	Range  string           `json:"range,omitempty"`  // ⁣查询范围, 包含 sheetId 与单元格范围两部分, 目前支持三种索引方式, 详见 [在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)。range所表示的范围需要大于等于values占用的范围。
	Values [][]SheetContent `json:"values,omitempty"` // 需要写入的值, 如要写入公式、超链接、email、@人等, 可详看附录[sheet 支持写入数据类型](https://open.feishu.cn/document/ukTMukTMukTM/ugjN1UjL4YTN14CO2UTN)
}

// AppendSheetValueResp ...
type AppendSheetValueResp struct {
	SpreadSheetToken string                       `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	TableRange       string                       `json:"tableRange,omitempty"`       // 写入的范围
	Revision         int64                        `json:"revision,omitempty"`         // sheet 的版本号
	Updates          *AppendSheetValueRespUpdates `json:"updates,omitempty"`          // 插入数据的范围、行列数等
}

// AppendSheetValueRespUpdates ...
type AppendSheetValueRespUpdates struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 写入的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 写入的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 写入的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 写入的单元格总数
	Revision         int64  `json:"revision,omitempty"`         // sheet 的版本号
}

// appendSheetValueResp ...
type appendSheetValueResp struct {
	Code int64                 `json:"code,omitempty"`
	Msg  string                `json:"msg,omitempty"`
	Data *AppendSheetValueResp `json:"data,omitempty"`
}
