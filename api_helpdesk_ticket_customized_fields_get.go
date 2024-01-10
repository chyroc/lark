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

// GetHelpdeskTicketCustomizedFields 该接口用于获取服务台自定义字段详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/customized_fields
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/ticket-management/ticket/customized_fields
func (r *HelpdeskService) GetHelpdeskTicketCustomizedFields(ctx context.Context, request *GetHelpdeskTicketCustomizedFieldsReq, options ...MethodOptionFunc) (*GetHelpdeskTicketCustomizedFieldsResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskTicketCustomizedFields != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskTicketCustomizedFields mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskTicketCustomizedFields(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskTicketCustomizedFields",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/customized_fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskTicketCustomizedFieldsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskTicketCustomizedFields mock HelpdeskGetHelpdeskTicketCustomizedFields method
func (r *Mock) MockHelpdeskGetHelpdeskTicketCustomizedFields(f func(ctx context.Context, request *GetHelpdeskTicketCustomizedFieldsReq, options ...MethodOptionFunc) (*GetHelpdeskTicketCustomizedFieldsResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskTicketCustomizedFields = f
}

// UnMockHelpdeskGetHelpdeskTicketCustomizedFields un-mock HelpdeskGetHelpdeskTicketCustomizedFields method
func (r *Mock) UnMockHelpdeskGetHelpdeskTicketCustomizedFields() {
	r.mockHelpdeskGetHelpdeskTicketCustomizedFields = nil
}

// GetHelpdeskTicketCustomizedFieldsReq ...
type GetHelpdeskTicketCustomizedFieldsReq struct {
	VisibleOnly *bool `query:"visible_only" json:"-"` // visible only, 示例值: true
}

// GetHelpdeskTicketCustomizedFieldsResp ...
type GetHelpdeskTicketCustomizedFieldsResp struct {
	UserCustomizedFields   []*GetHelpdeskTicketCustomizedFieldsRespUserCustomizedField   `json:"user_customized_fields,omitempty"`   // 用户自定义字段
	TicketCustomizedFields []*GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedField `json:"ticket_customized_fields,omitempty"` // 自定义工单字段
}

// GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedField ...
type GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedField struct {
	TicketCustomizedFieldID string                                                               `json:"ticket_customized_field_id,omitempty"` // 工单自定义字段ID
	HelpdeskID              string                                                               `json:"helpdesk_id,omitempty"`                // 服务台ID
	KeyName                 string                                                               `json:"key_name,omitempty"`                   // 键名
	DisplayName             string                                                               `json:"display_name,omitempty"`               // 名称
	Position                string                                                               `json:"position,omitempty"`                   // 字段在列表后台管理列表中的位置
	FieldType               string                                                               `json:"field_type,omitempty"`                 // 类型, string - 单行文本, multiline - 多行文本, dropdown - 下拉列表, dropdown_nested - 级联下拉
	Description             string                                                               `json:"description,omitempty"`                // 描述
	Visible                 bool                                                                 `json:"visible,omitempty"`                    // 是否可见
	Editable                bool                                                                 `json:"editable,omitempty"`                   // 是否可以修改
	Required                bool                                                                 `json:"required,omitempty"`                   // 是否必填
	CreatedAt               string                                                               `json:"created_at,omitempty"`                 // 创建时间
	UpdatedAt               string                                                               `json:"updated_at,omitempty"`                 // 更新时间
	CreatedBy               *GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedFieldCreatedBy `json:"created_by,omitempty"`                 // 创建用户
	UpdatedBy               *GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedFieldUpdatedBy `json:"updated_by,omitempty"`                 // 更新用户
	DropdownAllowMultiple   bool                                                                 `json:"dropdown_allow_multiple,omitempty"`    // 是否支持多选, 仅在字段类型是dropdown的时候有效
}

// GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedFieldCreatedBy ...
type GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedFieldCreatedBy struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Email      string `json:"email,omitempty"`      // 用户邮箱
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedFieldUpdatedBy ...
type GetHelpdeskTicketCustomizedFieldsRespTicketCustomizedFieldUpdatedBy struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Email      string `json:"email,omitempty"`      // 用户邮箱
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// GetHelpdeskTicketCustomizedFieldsRespUserCustomizedField ...
type GetHelpdeskTicketCustomizedFieldsRespUserCustomizedField struct {
	UserCustomizedFieldID string `json:"user_customized_field_id,omitempty"` // 字段ID
	ID                    string `json:"id,omitempty"`                       // 旧字段ID, 向后兼容用
	HelpdeskID            string `json:"helpdesk_id,omitempty"`              // 服务台ID
	KeyName               string `json:"key_name,omitempty"`                 // 字段键
	DisplayName           string `json:"display_name,omitempty"`             // 字段展示名称
	Position              string `json:"position,omitempty"`                 // 字段在列表中的展示位置
	FieldType             string `json:"field_type,omitempty"`               // 字段类型
	Description           string `json:"description,omitempty"`              // 字段描述信息
	Visible               bool   `json:"visible,omitempty"`                  // 字段是否可见
	Editable              bool   `json:"editable,omitempty"`                 // 字段是否可编辑
	Required              bool   `json:"required,omitempty"`                 // 字段是否必填
	CreatedAt             string `json:"created_at,omitempty"`               // 字段创建时间
	UpdatedAt             string `json:"updated_at,omitempty"`               // 字段修改时间
}

// getHelpdeskTicketCustomizedFieldsResp ...
type getHelpdeskTicketCustomizedFieldsResp struct {
	Code int64                                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskTicketCustomizedFieldsResp `json:"data,omitempty"`
}
