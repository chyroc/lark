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

// GetCoreHrJobFamilyList 批量查询序列。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_family/list
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/job-management/job_family/list
func (r *CoreHrService) GetCoreHrJobFamilyList(ctx context.Context, request *GetCoreHrJobFamilyListReq, options ...MethodOptionFunc) (*GetCoreHrJobFamilyListResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrJobFamilyList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrJobFamilyList mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrJobFamilyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrJobFamilyList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_families",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrJobFamilyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrJobFamilyList mock CoreHrGetCoreHrJobFamilyList method
func (r *Mock) MockCoreHrGetCoreHrJobFamilyList(f func(ctx context.Context, request *GetCoreHrJobFamilyListReq, options ...MethodOptionFunc) (*GetCoreHrJobFamilyListResp, *Response, error)) {
	r.mockCoreHrGetCoreHrJobFamilyList = f
}

// UnMockCoreHrGetCoreHrJobFamilyList un-mock CoreHrGetCoreHrJobFamilyList method
func (r *Mock) UnMockCoreHrGetCoreHrJobFamilyList() {
	r.mockCoreHrGetCoreHrJobFamilyList = nil
}

// GetCoreHrJobFamilyListReq ...
type GetCoreHrJobFamilyListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100
}

// GetCoreHrJobFamilyListResp ...
type GetCoreHrJobFamilyListResp struct {
	Items     []*GetCoreHrJobFamilyListRespItem `json:"items,omitempty"`      // 查询的序列信息
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHrJobFamilyListRespItem ...
type GetCoreHrJobFamilyListRespItem struct {
	ID             string                                       `json:"id,omitempty"`              // 序列 ID
	Name           []*GetCoreHrJobFamilyListRespItemName        `json:"name,omitempty"`            // 名称
	Active         bool                                         `json:"active,omitempty"`          // 是否启用
	ParentID       string                                       `json:"parent_id,omitempty"`       // 上级序列 ID, 枚举值及详细信息可通过【批量查询序列】接口查询获得
	EffectiveTime  string                                       `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                       `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                       `json:"code,omitempty"`            // 编码
	CustomFields   []*GetCoreHrJobFamilyListRespItemCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// GetCoreHrJobFamilyListRespItemCustomField ...
type GetCoreHrJobFamilyListRespItemCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHrJobFamilyListRespItemName ...
type GetCoreHrJobFamilyListRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHrJobFamilyListResp ...
type getCoreHrJobFamilyListResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrJobFamilyListResp `json:"data,omitempty"`
}
