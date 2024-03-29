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

// DeleteCoreHRCostCenterVersion 撤销成本中心版本
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/cost_center-version/delete
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/cost_center/cost_center-version/delete
func (r *CoreHRService) DeleteCoreHRCostCenterVersion(ctx context.Context, request *DeleteCoreHRCostCenterVersionReq, options ...MethodOptionFunc) (*DeleteCoreHRCostCenterVersionResp, *Response, error) {
	if r.cli.mock.mockCoreHRDeleteCoreHRCostCenterVersion != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#DeleteCoreHRCostCenterVersion mock enable")
		return r.cli.mock.mockCoreHRDeleteCoreHRCostCenterVersion(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "DeleteCoreHRCostCenterVersion",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/cost_centers/:cost_center_id/versions/:version_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHRCostCenterVersionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRDeleteCoreHRCostCenterVersion mock CoreHRDeleteCoreHRCostCenterVersion method
func (r *Mock) MockCoreHRDeleteCoreHRCostCenterVersion(f func(ctx context.Context, request *DeleteCoreHRCostCenterVersionReq, options ...MethodOptionFunc) (*DeleteCoreHRCostCenterVersionResp, *Response, error)) {
	r.mockCoreHRDeleteCoreHRCostCenterVersion = f
}

// UnMockCoreHRDeleteCoreHRCostCenterVersion un-mock CoreHRDeleteCoreHRCostCenterVersion method
func (r *Mock) UnMockCoreHRDeleteCoreHRCostCenterVersion() {
	r.mockCoreHRDeleteCoreHRCostCenterVersion = nil
}

// DeleteCoreHRCostCenterVersionReq ...
type DeleteCoreHRCostCenterVersionReq struct {
	CostCenterID    string `path:"cost_center_id" json:"-"`    // 成本中心ID, 示例值: "6862995757234914824"
	VersionID       string `path:"version_id" json:"-"`        // 版本ID, 示例值: "6862995757234914824"
	OperationReason string `json:"operation_reason,omitempty"` // 操作原因, 示例值: "随着组织架构调整, 该成本中心不再使用"
}

// DeleteCoreHRCostCenterVersionResp ...
type DeleteCoreHRCostCenterVersionResp struct {
}

// deleteCoreHRCostCenterVersionResp ...
type deleteCoreHRCostCenterVersionResp struct {
	Code  int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                             `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteCoreHRCostCenterVersionResp `json:"data,omitempty"`
	Error *ErrorDetail                       `json:"error,omitempty"`
}
