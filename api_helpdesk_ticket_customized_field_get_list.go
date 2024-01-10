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

// GetHelpdeskTicketCustomizedFieldList 该接口用于获取全部工单自定义字段。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/list-ticket-customized-fields
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/ticket-management/ticket_customized_field/list-ticket-customized-fields
func (r *HelpdeskService) GetHelpdeskTicketCustomizedFieldList(ctx context.Context, request *GetHelpdeskTicketCustomizedFieldListReq, options ...MethodOptionFunc) (*GetHelpdeskTicketCustomizedFieldListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskTicketCustomizedFieldList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskTicketCustomizedFieldList mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskTicketCustomizedFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskTicketCustomizedFieldList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/ticket_customized_fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskTicketCustomizedFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskTicketCustomizedFieldList mock HelpdeskGetHelpdeskTicketCustomizedFieldList method
func (r *Mock) MockHelpdeskGetHelpdeskTicketCustomizedFieldList(f func(ctx context.Context, request *GetHelpdeskTicketCustomizedFieldListReq, options ...MethodOptionFunc) (*GetHelpdeskTicketCustomizedFieldListResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskTicketCustomizedFieldList = f
}

// UnMockHelpdeskGetHelpdeskTicketCustomizedFieldList un-mock HelpdeskGetHelpdeskTicketCustomizedFieldList method
func (r *Mock) UnMockHelpdeskGetHelpdeskTicketCustomizedFieldList() {
	r.mockHelpdeskGetHelpdeskTicketCustomizedFieldList = nil
}

// GetHelpdeskTicketCustomizedFieldListReq ...
type GetHelpdeskTicketCustomizedFieldListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "6948728206392295444"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10；默认为20, 最大值: `100`
	Visible   *bool   `json:"visible,omitempty"`    // 是否可见, 示例值: true
}

// GetHelpdeskTicketCustomizedFieldListResp ...
type GetHelpdeskTicketCustomizedFieldListResp struct {
	HasMore       bool                                            `json:"has_more,omitempty"`        // 是否还有更多项
	NextPageToken string                                          `json:"next_page_token,omitempty"` // 下一分页标识
	Items         []*GetHelpdeskTicketCustomizedFieldListRespItem `json:"items,omitempty"`           // 工单自定义字段列表
}

// GetHelpdeskTicketCustomizedFieldListRespItem ...
type GetHelpdeskTicketCustomizedFieldListRespItem struct {
	TicketCustomizedFieldID string                                                 `json:"ticket_customized_field_id,omitempty"` // 工单自定义字段ID
	HelpdeskID              string                                                 `json:"helpdesk_id,omitempty"`                // 服务台ID
	KeyName                 string                                                 `json:"key_name,omitempty"`                   // 键名
	DisplayName             string                                                 `json:"display_name,omitempty"`               // 名称
	Position                string                                                 `json:"position,omitempty"`                   // 字段在列表后台管理列表中的位置
	FieldType               string                                                 `json:"field_type,omitempty"`                 // 类型
	Description             string                                                 `json:"description,omitempty"`                // 描述
	Visible                 bool                                                   `json:"visible,omitempty"`                    // 是否可见
	Editable                bool                                                   `json:"editable,omitempty"`                   // 是否可以修改
	Required                bool                                                   `json:"required,omitempty"`                   // 是否必填
	CreatedAt               string                                                 `json:"created_at,omitempty"`                 // 创建时间
	UpdatedAt               string                                                 `json:"updated_at,omitempty"`                 // 更新时间
	CreatedBy               *GetHelpdeskTicketCustomizedFieldListRespItemCreatedBy `json:"created_by,omitempty"`                 // 创建用户
	UpdatedBy               *GetHelpdeskTicketCustomizedFieldListRespItemUpdatedBy `json:"updated_by,omitempty"`                 // 更新用户
	DropdownOptions         *HelpdeskDropdownOption                                `json:"dropdown_options,omitempty"`           // 下拉列表选项
	DropdownAllowMultiple   bool                                                   `json:"dropdown_allow_multiple,omitempty"`    // 是否支持多选, 仅在字段类型是dropdown的时候有效
}

// GetHelpdeskTicketCustomizedFieldListRespItemCreatedBy ...
type GetHelpdeskTicketCustomizedFieldListRespItemCreatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

// GetHelpdeskTicketCustomizedFieldListRespItemUpdatedBy ...
type GetHelpdeskTicketCustomizedFieldListRespItemUpdatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

// getHelpdeskTicketCustomizedFieldListResp ...
type getHelpdeskTicketCustomizedFieldListResp struct {
	Code int64                                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                    `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskTicketCustomizedFieldListResp `json:"data,omitempty"`
}
