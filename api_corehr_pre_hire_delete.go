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

// DeleteCoreHRPreHire 删除待入职人员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/pre_hire/delete
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/pre_hire/delete
func (r *CoreHRService) DeleteCoreHRPreHire(ctx context.Context, request *DeleteCoreHRPreHireReq, options ...MethodOptionFunc) (*DeleteCoreHRPreHireResp, *Response, error) {
	if r.cli.mock.mockCoreHRDeleteCoreHRPreHire != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#DeleteCoreHRPreHire mock enable")
		return r.cli.mock.mockCoreHRDeleteCoreHRPreHire(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "DeleteCoreHRPreHire",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/pre_hires/:pre_hire_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHRPreHireResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRDeleteCoreHRPreHire mock CoreHRDeleteCoreHRPreHire method
func (r *Mock) MockCoreHRDeleteCoreHRPreHire(f func(ctx context.Context, request *DeleteCoreHRPreHireReq, options ...MethodOptionFunc) (*DeleteCoreHRPreHireResp, *Response, error)) {
	r.mockCoreHRDeleteCoreHRPreHire = f
}

// UnMockCoreHRDeleteCoreHRPreHire un-mock CoreHRDeleteCoreHRPreHire method
func (r *Mock) UnMockCoreHRDeleteCoreHRPreHire() {
	r.mockCoreHRDeleteCoreHRPreHire = nil
}

// DeleteCoreHRPreHireReq ...
type DeleteCoreHRPreHireReq struct {
	PreHireID string `path:"pre_hire_id" json:"-"` // 需要删除的待入职人员信息ID, 示例值: "76534545454"
}

// DeleteCoreHRPreHireResp ...
type DeleteCoreHRPreHireResp struct {
}

// deleteCoreHRPreHireResp ...
type deleteCoreHRPreHireResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCoreHRPreHireResp `json:"data,omitempty"`
}
