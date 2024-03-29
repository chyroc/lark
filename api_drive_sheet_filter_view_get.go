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

// GetSheetFilterView 获取指定筛选视图 id 的名字和筛选范围。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/get
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/spreadsheet-sheet-filter_view/get
func (r *DriveService) GetSheetFilterView(ctx context.Context, request *GetSheetFilterViewReq, options ...MethodOptionFunc) (*GetSheetFilterViewResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetFilterView != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetSheetFilterView mock enable")
		return r.cli.mock.mockDriveGetSheetFilterView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSheetFilterView",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetFilterViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetSheetFilterView mock DriveGetSheetFilterView method
func (r *Mock) MockDriveGetSheetFilterView(f func(ctx context.Context, request *GetSheetFilterViewReq, options ...MethodOptionFunc) (*GetSheetFilterViewResp, *Response, error)) {
	r.mockDriveGetSheetFilterView = f
}

// UnMockDriveGetSheetFilterView un-mock DriveGetSheetFilterView method
func (r *Mock) UnMockDriveGetSheetFilterView() {
	r.mockDriveGetSheetFilterView = nil
}

// GetSheetFilterViewReq ...
type GetSheetFilterViewReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值: "shtcnmBA*yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值: "0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值: "pH9hbVcCXA"
}

// GetSheetFilterViewResp ...
type GetSheetFilterViewResp struct {
	FilterView *GetSheetFilterViewRespFilterView `json:"filter_view,omitempty"` // 筛选视图信息, 包括 id、name、range
}

// GetSheetFilterViewRespFilterView ...
type GetSheetFilterViewRespFilterView struct {
	FilterViewID   string `json:"filter_view_id,omitempty"`   // 筛选视图 id
	FilterViewName string `json:"filter_view_name,omitempty"` // 筛选视图名字
	Range          string `json:"range,omitempty"`            // 筛选视图的筛选范围
}

// getSheetFilterViewResp ...
type getSheetFilterViewResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *GetSheetFilterViewResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
