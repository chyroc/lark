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

// BatchUpdateHireEcoBackgroundCheckPackage 更新帐号下指定背调套餐和附加调查项信息, 更新将影响已发起背调的表单项展示。如需新增背调套餐、附加项需使用创建接口进行添加
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check_package/batch_update
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check_package/batch_update
func (r *HireService) BatchUpdateHireEcoBackgroundCheckPackage(ctx context.Context, request *BatchUpdateHireEcoBackgroundCheckPackageReq, options ...MethodOptionFunc) (*BatchUpdateHireEcoBackgroundCheckPackageResp, *Response, error) {
	if r.cli.mock.mockHireBatchUpdateHireEcoBackgroundCheckPackage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#BatchUpdateHireEcoBackgroundCheckPackage mock enable")
		return r.cli.mock.mockHireBatchUpdateHireEcoBackgroundCheckPackage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "BatchUpdateHireEcoBackgroundCheckPackage",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/eco_background_check_packages/batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchUpdateHireEcoBackgroundCheckPackageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireBatchUpdateHireEcoBackgroundCheckPackage mock HireBatchUpdateHireEcoBackgroundCheckPackage method
func (r *Mock) MockHireBatchUpdateHireEcoBackgroundCheckPackage(f func(ctx context.Context, request *BatchUpdateHireEcoBackgroundCheckPackageReq, options ...MethodOptionFunc) (*BatchUpdateHireEcoBackgroundCheckPackageResp, *Response, error)) {
	r.mockHireBatchUpdateHireEcoBackgroundCheckPackage = f
}

// UnMockHireBatchUpdateHireEcoBackgroundCheckPackage un-mock HireBatchUpdateHireEcoBackgroundCheckPackage method
func (r *Mock) UnMockHireBatchUpdateHireEcoBackgroundCheckPackage() {
	r.mockHireBatchUpdateHireEcoBackgroundCheckPackage = nil
}

// BatchUpdateHireEcoBackgroundCheckPackageReq ...
type BatchUpdateHireEcoBackgroundCheckPackageReq struct {
	AccountID          string                                                       `json:"account_id,omitempty"`           // 背调账号 ID, 可在「账号绑定」事件中获取, 示例值: "ord_id"
	PackageList        []*BatchUpdateHireEcoBackgroundCheckPackageReqPackage        `json:"package_list,omitempty"`         // 背调套餐列表
	AdditionalItemList []*BatchUpdateHireEcoBackgroundCheckPackageReqAdditionalItem `json:"additional_item_list,omitempty"` // 附加项列表
}

// BatchUpdateHireEcoBackgroundCheckPackageReqAdditionalItem ...
type BatchUpdateHireEcoBackgroundCheckPackageReqAdditionalItem struct {
	ID          string  `json:"id,omitempty"`          // 附件调查项 ID, 示例值: "ext001"
	Name        string  `json:"name,omitempty"`        // 附加调查项名称, 示例值: "工作履历信息验证X2"
	Description *string `json:"description,omitempty"` // 附加调查项描述, 示例值: "详细调查"
}

// BatchUpdateHireEcoBackgroundCheckPackageReqPackage ...
type BatchUpdateHireEcoBackgroundCheckPackageReqPackage struct {
	ID          string  `json:"id,omitempty"`          // 套餐 ID, 示例值: "pkg001"
	Name        string  `json:"name,omitempty"`        // 背调名称, 示例值: "基础套餐", 最大长度: `36` 字符
	Description *string `json:"description,omitempty"` // 套餐描述, 示例值: "工作履历信息验证X1, 工作表现鉴定评价X1, 教育背景核实, 公民身份信息验证, 简历对比, 民事诉讼调查"
}

// BatchUpdateHireEcoBackgroundCheckPackageResp ...
type BatchUpdateHireEcoBackgroundCheckPackageResp struct {
}

// batchUpdateHireEcoBackgroundCheckPackageResp ...
type batchUpdateHireEcoBackgroundCheckPackageResp struct {
	Code  int64                                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                        `json:"msg,omitempty"`  // 错误描述
	Data  *BatchUpdateHireEcoBackgroundCheckPackageResp `json:"data,omitempty"`
	Error *ErrorDetail                                  `json:"error,omitempty"`
}
