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

// BatchDeleteHireEcoBackgroundCheckPackage 删除指定帐号的指定背调套餐和附加调查项信息, 删除不会影响已创建的背调。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check_package/batch_delete
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check_package/batch_delete
func (r *HireService) BatchDeleteHireEcoBackgroundCheckPackage(ctx context.Context, request *BatchDeleteHireEcoBackgroundCheckPackageReq, options ...MethodOptionFunc) (*BatchDeleteHireEcoBackgroundCheckPackageResp, *Response, error) {
	if r.cli.mock.mockHireBatchDeleteHireEcoBackgroundCheckPackage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#BatchDeleteHireEcoBackgroundCheckPackage mock enable")
		return r.cli.mock.mockHireBatchDeleteHireEcoBackgroundCheckPackage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "BatchDeleteHireEcoBackgroundCheckPackage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/eco_background_check_packages/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchDeleteHireEcoBackgroundCheckPackageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireBatchDeleteHireEcoBackgroundCheckPackage mock HireBatchDeleteHireEcoBackgroundCheckPackage method
func (r *Mock) MockHireBatchDeleteHireEcoBackgroundCheckPackage(f func(ctx context.Context, request *BatchDeleteHireEcoBackgroundCheckPackageReq, options ...MethodOptionFunc) (*BatchDeleteHireEcoBackgroundCheckPackageResp, *Response, error)) {
	r.mockHireBatchDeleteHireEcoBackgroundCheckPackage = f
}

// UnMockHireBatchDeleteHireEcoBackgroundCheckPackage un-mock HireBatchDeleteHireEcoBackgroundCheckPackage method
func (r *Mock) UnMockHireBatchDeleteHireEcoBackgroundCheckPackage() {
	r.mockHireBatchDeleteHireEcoBackgroundCheckPackage = nil
}

// BatchDeleteHireEcoBackgroundCheckPackageReq ...
type BatchDeleteHireEcoBackgroundCheckPackageReq struct {
	AccountID            string   `json:"account_id,omitempty"`              // 背调账号 ID, 可在「账号绑定」事件中获取, 示例值: "xd_bc_001"
	PackageIDList        []string `json:"package_id_list,omitempty"`         // 要删除的套餐 ID 列表, 删除套餐不影响已安排的背调, 示例值: ["7230753910687080001"]
	AdditionalItemIDList []string `json:"additional_item_id_list,omitempty"` // 要删除的附加调查项 ID 列表, 删除附加调查项不影响已安排的背调, 示例值: ["7230753910687080002"]
}

// BatchDeleteHireEcoBackgroundCheckPackageResp ...
type BatchDeleteHireEcoBackgroundCheckPackageResp struct {
}

// batchDeleteHireEcoBackgroundCheckPackageResp ...
type batchDeleteHireEcoBackgroundCheckPackageResp struct {
	Code int64                                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                        `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeleteHireEcoBackgroundCheckPackageResp `json:"data,omitempty"`
}
