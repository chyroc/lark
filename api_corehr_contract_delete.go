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

// DeleteCoreHrContract 删除合同。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/contract/delete
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/contract/delete
func (r *CoreHrService) DeleteCoreHrContract(ctx context.Context, request *DeleteCoreHrContractReq, options ...MethodOptionFunc) (*DeleteCoreHrContractResp, *Response, error) {
	if r.cli.mock.mockCoreHrDeleteCoreHrContract != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#DeleteCoreHrContract mock enable")
		return r.cli.mock.mockCoreHrDeleteCoreHrContract(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "DeleteCoreHrContract",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/contracts/:contract_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHrContractResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrDeleteCoreHrContract mock CoreHrDeleteCoreHrContract method
func (r *Mock) MockCoreHrDeleteCoreHrContract(f func(ctx context.Context, request *DeleteCoreHrContractReq, options ...MethodOptionFunc) (*DeleteCoreHrContractResp, *Response, error)) {
	r.mockCoreHrDeleteCoreHrContract = f
}

// UnMockCoreHrDeleteCoreHrContract un-mock CoreHrDeleteCoreHrContract method
func (r *Mock) UnMockCoreHrDeleteCoreHrContract() {
	r.mockCoreHrDeleteCoreHrContract = nil
}

// DeleteCoreHrContractReq ...
type DeleteCoreHrContractReq struct {
	ContractID string `path:"contract_id" json:"-"` // 需要删除的合同 ID, 示例值: "4137834332"
}

// DeleteCoreHrContractResp ...
type DeleteCoreHrContractResp struct {
}

// deleteCoreHrContractResp ...
type deleteCoreHrContractResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCoreHrContractResp `json:"data,omitempty"`
}
