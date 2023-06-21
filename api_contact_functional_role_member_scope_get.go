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

// GetContactFunctionalRoleMemberScope 通过本接口可以查询某个成员的管理范围
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role-member/get
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/functional_role-member/get
func (r *ContactService) GetContactFunctionalRoleMemberScope(ctx context.Context, request *GetContactFunctionalRoleMemberScopeReq, options ...MethodOptionFunc) (*GetContactFunctionalRoleMemberScopeResp, *Response, error) {
	if r.cli.mock.mockContactGetContactFunctionalRoleMemberScope != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactFunctionalRoleMemberScope mock enable")
		return r.cli.mock.mockContactGetContactFunctionalRoleMemberScope(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactFunctionalRoleMemberScope",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/functional_roles/:role_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactFunctionalRoleMemberScopeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactFunctionalRoleMemberScope mock ContactGetContactFunctionalRoleMemberScope method
func (r *Mock) MockContactGetContactFunctionalRoleMemberScope(f func(ctx context.Context, request *GetContactFunctionalRoleMemberScopeReq, options ...MethodOptionFunc) (*GetContactFunctionalRoleMemberScopeResp, *Response, error)) {
	r.mockContactGetContactFunctionalRoleMemberScope = f
}

// UnMockContactGetContactFunctionalRoleMemberScope un-mock ContactGetContactFunctionalRoleMemberScope method
func (r *Mock) UnMockContactGetContactFunctionalRoleMemberScope() {
	r.mockContactGetContactFunctionalRoleMemberScope = nil
}

// GetContactFunctionalRoleMemberScopeReq ...
type GetContactFunctionalRoleMemberScopeReq struct {
	RoleID           string            `path:"role_id" json:"-"`             // 角色的唯一标识, 单租户下唯一, 示例值: "7vrj3vk70xk7v5r"
	MemberID         string            `path:"member_id" json:"-"`           // 要查询的角色内成员ID, 示例值: "od-123456"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
}

// GetContactFunctionalRoleMemberScopeResp ...
type GetContactFunctionalRoleMemberScopeResp struct {
	Member *GetContactFunctionalRoleMemberScopeRespMember `json:"member,omitempty"` // 成员的管理范围
}

// GetContactFunctionalRoleMemberScopeRespMember ...
type GetContactFunctionalRoleMemberScopeRespMember struct {
	UserID        string   `json:"user_id,omitempty"`        // 成员ID
	ScopeType     string   `json:"scope_type,omitempty"`     // 管理范围的类型, 可选值有: All: 管理范围是全部, Part: 管理范围是部分, None: 管理范围为空
	DepartmentIDs []string `json:"department_ids,omitempty"` // 表示该角色成员的管理范围, scope_type为“指定范围”时, 返回该值
}

// getContactFunctionalRoleMemberScopeResp ...
type getContactFunctionalRoleMemberScopeResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *GetContactFunctionalRoleMemberScopeResp `json:"data,omitempty"`
}
