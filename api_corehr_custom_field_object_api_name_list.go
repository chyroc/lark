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

// GetCoreHRCustomFieldObjectApiNameList 获取「飞书人事」中的对象列表, 含系统预置对象与自定义对象。使用方式可参考[【操作手册】如何通过 OpenAPI 维护自定义字段](https://feishu.feishu.cn/docx/QlUudBfCtosWMbxx3vxcOFDknn7)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/list_object_api_name
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/custom_field/list_object_api_name
func (r *CoreHRService) GetCoreHRCustomFieldObjectApiNameList(ctx context.Context, request *GetCoreHRCustomFieldObjectApiNameListReq, options ...MethodOptionFunc) (*GetCoreHRCustomFieldObjectApiNameListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCustomFieldObjectApiNameList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCustomFieldObjectApiNameList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCustomFieldObjectApiNameList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCustomFieldObjectApiNameList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/custom_fields/list_object_api_name",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCustomFieldObjectApiNameListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCustomFieldObjectApiNameList mock CoreHRGetCoreHRCustomFieldObjectApiNameList method
func (r *Mock) MockCoreHRGetCoreHRCustomFieldObjectApiNameList(f func(ctx context.Context, request *GetCoreHRCustomFieldObjectApiNameListReq, options ...MethodOptionFunc) (*GetCoreHRCustomFieldObjectApiNameListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCustomFieldObjectApiNameList = f
}

// UnMockCoreHRGetCoreHRCustomFieldObjectApiNameList un-mock CoreHRGetCoreHRCustomFieldObjectApiNameList method
func (r *Mock) UnMockCoreHRGetCoreHRCustomFieldObjectApiNameList() {
	r.mockCoreHRGetCoreHRCustomFieldObjectApiNameList = nil
}

// GetCoreHRCustomFieldObjectApiNameListReq ...
type GetCoreHRCustomFieldObjectApiNameListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 11
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100
}

// GetCoreHRCustomFieldObjectApiNameListResp ...
type GetCoreHRCustomFieldObjectApiNameListResp struct {
	Items     []*GetCoreHRCustomFieldObjectApiNameListRespItem `json:"items,omitempty"`      // 对象列表
	HasMore   bool                                             `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHRCustomFieldObjectApiNameListRespItem ...
type GetCoreHRCustomFieldObjectApiNameListRespItem struct {
	ObjectApiName string                                             `json:"object_api_name,omitempty"` // 对象的唯一标识
	Name          *GetCoreHRCustomFieldObjectApiNameListRespItemName `json:"name,omitempty"`            // 对象名称
	IsOpen        bool                                               `json:"is_open,omitempty"`         // 是否启用, True 为已启用, False 为未启用
	CreateTime    string                                             `json:"create_time,omitempty"`     // 创建时间, 秒级时间戳
	UpdateTime    string                                             `json:"update_time,omitempty"`     // 更新时间, 秒级时间戳
}

// GetCoreHRCustomFieldObjectApiNameListRespItemName ...
type GetCoreHRCustomFieldObjectApiNameListRespItemName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getCoreHRCustomFieldObjectApiNameListResp ...
type getCoreHRCustomFieldObjectApiNameListResp struct {
	Code int64                                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                     `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRCustomFieldObjectApiNameListResp `json:"data,omitempty"`
}
