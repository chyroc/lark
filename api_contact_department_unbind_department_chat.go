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

// UnbindDepartmentChat 通过该接口将部门群转为普通群。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/unbind_department_chat
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/department/unbind_department_chat
func (r *ContactService) UnbindDepartmentChat(ctx context.Context, request *UnbindDepartmentChatReq, options ...MethodOptionFunc) (*UnbindDepartmentChatResp, *Response, error) {
	if r.cli.mock.mockContactUnbindDepartmentChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UnbindDepartmentChat mock enable")
		return r.cli.mock.mockContactUnbindDepartmentChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UnbindDepartmentChat",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/departments/unbind_department_chat",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(unbindDepartmentChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUnbindDepartmentChat mock ContactUnbindDepartmentChat method
func (r *Mock) MockContactUnbindDepartmentChat(f func(ctx context.Context, request *UnbindDepartmentChatReq, options ...MethodOptionFunc) (*UnbindDepartmentChatResp, *Response, error)) {
	r.mockContactUnbindDepartmentChat = f
}

// UnMockContactUnbindDepartmentChat un-mock ContactUnbindDepartmentChat method
func (r *Mock) UnMockContactUnbindDepartmentChat() {
	r.mockContactUnbindDepartmentChat = nil
}

// UnbindDepartmentChatReq ...
type UnbindDepartmentChatReq struct {
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 默认为"open_department_id", 示例值: "open_department_id", 可选值有: department_id: 用来标识租户内一个唯一的部门, open_department_id: 用来在具体某个应用中标识一个部门, 同一个部门 在不同应用中的 open_department_id 不相同。
	DepartmentID     string            `json:"department_id,omitempty"`      // 部门ID, 示例值: "D096"
}

// UnbindDepartmentChatResp ...
type UnbindDepartmentChatResp struct {
}

// unbindDepartmentChatResp ...
type unbindDepartmentChatResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *UnbindDepartmentChatResp `json:"data,omitempty"`
}
