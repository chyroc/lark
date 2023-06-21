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

// CreateSheetFilterViewCondition 在筛选视图的筛选范围的某一列创建筛选条件。
//
// 筛选条件参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/create
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/spreadsheet-sheet-filter_view/spreadsheet-sheet-filter_view-condition/create
func (r *DriveService) CreateSheetFilterViewCondition(ctx context.Context, request *CreateSheetFilterViewConditionReq, options ...MethodOptionFunc) (*CreateSheetFilterViewConditionResp, *Response, error) {
	if r.cli.mock.mockDriveCreateSheetFilterViewCondition != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheetFilterViewCondition mock enable")
		return r.cli.mock.mockDriveCreateSheetFilterViewCondition(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateSheetFilterViewCondition",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createSheetFilterViewConditionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateSheetFilterViewCondition mock DriveCreateSheetFilterViewCondition method
func (r *Mock) MockDriveCreateSheetFilterViewCondition(f func(ctx context.Context, request *CreateSheetFilterViewConditionReq, options ...MethodOptionFunc) (*CreateSheetFilterViewConditionResp, *Response, error)) {
	r.mockDriveCreateSheetFilterViewCondition = f
}

// UnMockDriveCreateSheetFilterViewCondition un-mock DriveCreateSheetFilterViewCondition method
func (r *Mock) UnMockDriveCreateSheetFilterViewCondition() {
	r.mockDriveCreateSheetFilterViewCondition = nil
}

// CreateSheetFilterViewConditionReq ...
type CreateSheetFilterViewConditionReq struct {
	SpreadSheetToken string   `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值: "shtcnmBA*yGehy8"
	SheetID          string   `path:"sheet_id" json:"-"`          // 子表 id, 示例值: "0b**12"
	FilterViewID     string   `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值: "pH9hbVcCXA"
	ConditionID      *string  `json:"condition_id,omitempty"`     // 设置筛选条件的列, 使用字母号, 示例值: "E"
	FilterType       *string  `json:"filter_type,omitempty"`      // 筛选类型, 示例值: "number"
	CompareType      *string  `json:"compare_type,omitempty"`     // 比较类型, 示例值: "less"
	Expected         []string `json:"expected,omitempty"`         // 筛选参数, 示例值: 6
}

// CreateSheetFilterViewConditionResp ...
type CreateSheetFilterViewConditionResp struct {
	Condition *CreateSheetFilterViewConditionRespCondition `json:"condition,omitempty"` // 创建的筛选条件
}

// CreateSheetFilterViewConditionRespCondition ...
type CreateSheetFilterViewConditionRespCondition struct {
	ConditionID string   `json:"condition_id,omitempty"` // 设置筛选条件的列, 使用字母号
	FilterType  string   `json:"filter_type,omitempty"`  // 筛选类型
	CompareType string   `json:"compare_type,omitempty"` // 比较类型
	Expected    []string `json:"expected,omitempty"`     // 筛选参数
}

// createSheetFilterViewConditionResp ...
type createSheetFilterViewConditionResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *CreateSheetFilterViewConditionResp `json:"data,omitempty"`
}
