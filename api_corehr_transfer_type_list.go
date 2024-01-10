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

// GetCoreHRTransferTypeList 获取异动类型列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/transfer_type/query
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/job_change/query
func (r *CoreHRService) GetCoreHRTransferTypeList(ctx context.Context, request *GetCoreHRTransferTypeListReq, options ...MethodOptionFunc) (*GetCoreHRTransferTypeListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRTransferTypeList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRTransferTypeList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRTransferTypeList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRTransferTypeList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/transfer_types/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRTransferTypeListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRTransferTypeList mock CoreHRGetCoreHRTransferTypeList method
func (r *Mock) MockCoreHRGetCoreHRTransferTypeList(f func(ctx context.Context, request *GetCoreHRTransferTypeListReq, options ...MethodOptionFunc) (*GetCoreHRTransferTypeListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRTransferTypeList = f
}

// UnMockCoreHRGetCoreHRTransferTypeList un-mock CoreHRGetCoreHRTransferTypeList method
func (r *Mock) UnMockCoreHRGetCoreHRTransferTypeList() {
	r.mockCoreHRGetCoreHRTransferTypeList = nil
}

// GetCoreHRTransferTypeListReq ...
type GetCoreHRTransferTypeListReq struct {
	Active                       *bool    `query:"active" json:"-"`                          // 异动类型状态, 示例值: true
	TransferTypeUniqueIdentifier []string `query:"transfer_type_unique_identifier" json:"-"` // 异动类型唯一标识, 多条时最多数量为10, 示例值: job_status_change, 最大长度: `10`
}

// GetCoreHRTransferTypeListResp ...
type GetCoreHRTransferTypeListResp struct {
	Items []*GetCoreHRTransferTypeListRespItem `json:"items,omitempty"` // 异动类型列表
}

// GetCoreHRTransferTypeListRespItem ...
type GetCoreHRTransferTypeListRespItem struct {
	TransferTypeUniqueIdentifier string                                       `json:"transfer_type_unique_identifier,omitempty"` // 异动类型唯一标识
	Name                         []*GetCoreHRTransferTypeListRespItemName     `json:"name,omitempty"`                            // 异动类型名称
	Active                       bool                                         `json:"active,omitempty"`                          // 异动类型状态
	FlowID                       string                                       `json:"flow_id,omitempty"`                         // 关联流程唯一标识符
	FlowName                     []*GetCoreHRTransferTypeListRespItemFlowName `json:"flow_name,omitempty"`                       // 关联流程名称
	CreatedTime                  string                                       `json:"created_time,omitempty"`                    // 创建时间
	UpdatedTime                  string                                       `json:"updated_time,omitempty"`                    // 更新时间
}

// GetCoreHRTransferTypeListRespItemFlowName ...
type GetCoreHRTransferTypeListRespItemFlowName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRTransferTypeListRespItemName ...
type GetCoreHRTransferTypeListRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRTransferTypeListResp ...
type getCoreHRTransferTypeListResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRTransferTypeListResp `json:"data,omitempty"`
}
