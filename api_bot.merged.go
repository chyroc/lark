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

// GetBotInfo 获取机器人的基本信息。
//
// 需要启用机器人能力
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM
func (r *BotService) GetBotInfo(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error) {
	if r.cli.mock.mockBotGetBotInfo != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bot#GetBotInfo mock enable")
		return r.cli.mock.mockBotGetBotInfo(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bot",
		API:                   "GetBotInfo",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bot/v3/info",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getBotInfoResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBotGetBotInfo mock BotGetBotInfo method
func (r *Mock) MockBotGetBotInfo(f func(ctx context.Context, request *GetBotInfoReq, options ...MethodOptionFunc) (*GetBotInfoResp, *Response, error)) {
	r.mockBotGetBotInfo = f
}

// UnMockBotGetBotInfo un-mock BotGetBotInfo method
func (r *Mock) UnMockBotGetBotInfo() {
	r.mockBotGetBotInfo = nil
}

// GetBotInfoReq ...
type GetBotInfoReq struct{}

// getBotInfoResp ...
type getBotInfoResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetBotInfoResp `json:"bot,omitempty"`  // 机器人信息
}

// GetBotInfoResp ...
type GetBotInfoResp struct {
	ActivateStatus int64    `json:"activate_status,omitempty"` // app 当前状态。,0: 初始化，租户待安装,1: 租户停用,2: 租户启用,3: 安装后待启用,4: 升级待启用,5: license过期停用,6: Lark套餐到期或降级停用
	AppName        string   `json:"app_name,omitempty"`        // app 名称
	AvatarURL      string   `json:"avatar_url,omitempty"`      // app 图像地址
	IpWhiteList    []string `json:"ip_white_list,omitempty"`   // app 的 IP 白名单地址
	OpenID         string   `json:"open_id,omitempty"`         // 机器人的open_id
}

// Code generated by lark_sdk_gen. DO NOT EDIT.

// CreateCalendarACL
//
// 该接口用于以当前身份（应用 / 用户）给日历添加访问控制权限，即日历成员。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/create
func (r *CalendarService) CreateCalendarACL(ctx context.Context, request *CreateCalendarACLReq, options ...MethodOptionFunc) (*CreateCalendarACLResp, *Response, error) {
	if r.cli.mock.mockCalendarCreateCalendarACL != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#CreateCalendarACL mock enable")
		return r.cli.mock.mockCalendarCreateCalendarACL(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "CreateCalendarACL",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/acls",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createCalendarACLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarCreateCalendarACL mock CalendarCreateCalendarACL method
func (r *Mock) MockCalendarCreateCalendarACL(f func(ctx context.Context, request *CreateCalendarACLReq, options ...MethodOptionFunc) (*CreateCalendarACLResp, *Response, error)) {
	r.mockCalendarCreateCalendarACL = f
}

// UnMockCalendarCreateCalendarACL un-mock CalendarCreateCalendarACL method
func (r *Mock) UnMockCalendarCreateCalendarACL() {
	r.mockCalendarCreateCalendarACL = nil
}

// CreateCalendarACLReq ...
type CreateCalendarACLReq struct {
	UserIDType *IDType                    `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	CalendarID string                     `path:"calendar_id" json:"-"`   // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	Role       CalendarRole               `json:"role,omitempty"`         // 对日历的访问权限, 示例值："writer", 可选值有: `unknown`：未知权限, `free_busy_reader`：游客，只能看到忙碌/空闲信息, `reader`：订阅者，查看所有日程详情, `writer`：编辑者，创建及修改日程, `owner`：管理员，管理日历及共享设置
	Scope      *CreateCalendarACLReqScope `json:"scope,omitempty"`        // 权限范围
}

// CreateCalendarACLReqScope ...
type CreateCalendarACLReqScope struct {
	Type   string  `json:"type,omitempty"`    // 权限类型，当type为User时，值为open_id/user_id/union_id, 示例值："user", 可选值有: `user`：用户
	UserID *string `json:"user_id,omitempty"` // 用户ID，参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值："ou_xxxxxx"
}

// createCalendarACLResp ...
type createCalendarACLResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarACLResp `json:"data,omitempty"`
}

// CreateCalendarACLResp ...
type CreateCalendarACLResp struct {
	ACLID string                      `json:"acl_id,omitempty"` // acl资源ID。参见[ACL ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/introduction)
	Role  CalendarRole                `json:"role,omitempty"`   // 对日历的访问权限, 可选值有: `unknown`：未知权限, `free_busy_reader`：游客，只能看到忙碌/空闲信息, `reader`：订阅者，查看所有日程详情, `writer`：编辑者，创建及修改日程, `owner`：管理员，管理日历及共享设置
	Scope *CreateCalendarACLRespScope `json:"scope,omitempty"`  // 权限范围
}

// CreateCalendarACLRespScope ...
type CreateCalendarACLRespScope struct {
	Type   string `json:"type,omitempty"`    // 权限类型，当type为User时，值为open_id/user_id/union_id, 可选值有: `user`：用户
	UserID string `json:"user_id,omitempty"` // 用户ID，参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
}
