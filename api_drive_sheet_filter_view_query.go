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

// QuerySheetFilterView 查询子表内所有的筛选视图基本信息, 包括 id、name 和 range
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/query
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/spreadsheet-sheet-filter_view/query
func (r *DriveService) QuerySheetFilterView(ctx context.Context, request *QuerySheetFilterViewReq, options ...MethodOptionFunc) (*QuerySheetFilterViewResp, *Response, error) {
	if r.cli.mock.mockDriveQuerySheetFilterView != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#QuerySheetFilterView mock enable")
		return r.cli.mock.mockDriveQuerySheetFilterView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "QuerySheetFilterView",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(querySheetFilterViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveQuerySheetFilterView mock DriveQuerySheetFilterView method
func (r *Mock) MockDriveQuerySheetFilterView(f func(ctx context.Context, request *QuerySheetFilterViewReq, options ...MethodOptionFunc) (*QuerySheetFilterViewResp, *Response, error)) {
	r.mockDriveQuerySheetFilterView = f
}

// UnMockDriveQuerySheetFilterView un-mock DriveQuerySheetFilterView method
func (r *Mock) UnMockDriveQuerySheetFilterView() {
	r.mockDriveQuerySheetFilterView = nil
}

// QuerySheetFilterViewReq ...
type QuerySheetFilterViewReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值: "shtcnmBA*yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值: "0b**12"
}

// QuerySheetFilterViewResp ...
type QuerySheetFilterViewResp struct {
	Items []*QuerySheetFilterViewRespItem `json:"items,omitempty"` // 子表的所有筛选视图信息, id、name、range
}

// QuerySheetFilterViewRespItem ...
type QuerySheetFilterViewRespItem struct {
	FilterViewID   string `json:"filter_view_id,omitempty"`   // 筛选视图 id
	FilterViewName string `json:"filter_view_name,omitempty"` // 筛选视图名字
	Range          string `json:"range,omitempty"`            // 筛选视图的筛选范围
}

// querySheetFilterViewResp ...
type querySheetFilterViewResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *QuerySheetFilterViewResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
