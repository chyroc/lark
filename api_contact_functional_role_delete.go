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

// DeleteContactFunctionalRole 通过本接口可以删除某个角色
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role/delete
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/functional_role/delete
func (r *ContactService) DeleteContactFunctionalRole(ctx context.Context, request *DeleteContactFunctionalRoleReq, options ...MethodOptionFunc) (*DeleteContactFunctionalRoleResp, *Response, error) {
	if r.cli.mock.mockContactDeleteContactFunctionalRole != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteContactFunctionalRole mock enable")
		return r.cli.mock.mockContactDeleteContactFunctionalRole(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "DeleteContactFunctionalRole",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/functional_roles/:role_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteContactFunctionalRoleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactDeleteContactFunctionalRole mock ContactDeleteContactFunctionalRole method
func (r *Mock) MockContactDeleteContactFunctionalRole(f func(ctx context.Context, request *DeleteContactFunctionalRoleReq, options ...MethodOptionFunc) (*DeleteContactFunctionalRoleResp, *Response, error)) {
	r.mockContactDeleteContactFunctionalRole = f
}

// UnMockContactDeleteContactFunctionalRole un-mock ContactDeleteContactFunctionalRole method
func (r *Mock) UnMockContactDeleteContactFunctionalRole() {
	r.mockContactDeleteContactFunctionalRole = nil
}

// DeleteContactFunctionalRoleReq ...
type DeleteContactFunctionalRoleReq struct {
	RoleID string `path:"role_id" json:"-"` // 角色的唯一标识, 单租户下唯一, 示例值: "7vrj3vk70xk7v5r"
}

// DeleteContactFunctionalRoleResp ...
type DeleteContactFunctionalRoleResp struct {
}

// deleteContactFunctionalRoleResp ...
type deleteContactFunctionalRoleResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteContactFunctionalRoleResp `json:"data,omitempty"`
}
