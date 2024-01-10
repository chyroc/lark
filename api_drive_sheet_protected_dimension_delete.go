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

// DeleteSheetProtectedDimension 该接口用于根据保护范围ID删除保护范围, 最多支持同时删除10个ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYTM5YjL2ETO24iNxkjN
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/protect-range/delete-protection-scopes
func (r *DriveService) DeleteSheetProtectedDimension(ctx context.Context, request *DeleteSheetProtectedDimensionReq, options ...MethodOptionFunc) (*DeleteSheetProtectedDimensionResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetProtectedDimension != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetProtectedDimension mock enable")
		return r.cli.mock.mockDriveDeleteSheetProtectedDimension(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteSheetProtectedDimension",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_del",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteSheetProtectedDimensionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteSheetProtectedDimension mock DriveDeleteSheetProtectedDimension method
func (r *Mock) MockDriveDeleteSheetProtectedDimension(f func(ctx context.Context, request *DeleteSheetProtectedDimensionReq, options ...MethodOptionFunc) (*DeleteSheetProtectedDimensionResp, *Response, error)) {
	r.mockDriveDeleteSheetProtectedDimension = f
}

// UnMockDriveDeleteSheetProtectedDimension un-mock DriveDeleteSheetProtectedDimension method
func (r *Mock) UnMockDriveDeleteSheetProtectedDimension() {
	r.mockDriveDeleteSheetProtectedDimension = nil
}

// DeleteSheetProtectedDimensionReq ...
type DeleteSheetProtectedDimensionReq struct {
	SpreadSheetToken string   `path:"spreadsheetToken" json:"-"` // sheet 的 token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	ProtectIDs       []string `json:"protectIds,omitempty"`      // 需要删除的保护范围ID, 可以通过[获取表格元数据](https://open.feishu.cn/document/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN)接口获取
}

// DeleteSheetProtectedDimensionResp ...
type DeleteSheetProtectedDimensionResp struct {
	DelProtectIDs []string `json:"delProtectIds,omitempty"` // 成功删除的保护范围ID
}

// deleteSheetProtectedDimensionResp ...
type deleteSheetProtectedDimensionResp struct {
	Code int64                              `json:"code,omitempty"`
	Msg  string                             `json:"msg,omitempty"`
	Data *DeleteSheetProtectedDimensionResp `json:"data,omitempty"`
}
