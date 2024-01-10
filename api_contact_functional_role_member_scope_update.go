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

// UpdateContactFunctionalRoleMemberScope 通过该接口可设置本租户下角色成员的管理范围, 以便在审批等场景中应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role-member/scopes
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/functional_role-member/scopes
func (r *ContactService) UpdateContactFunctionalRoleMemberScope(ctx context.Context, request *UpdateContactFunctionalRoleMemberScopeReq, options ...MethodOptionFunc) (*UpdateContactFunctionalRoleMemberScopeResp, *Response, error) {
	if r.cli.mock.mockContactUpdateContactFunctionalRoleMemberScope != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#UpdateContactFunctionalRoleMemberScope mock enable")
		return r.cli.mock.mockContactUpdateContactFunctionalRoleMemberScope(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateContactFunctionalRoleMemberScope",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/functional_roles/:role_id/members/scopes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateContactFunctionalRoleMemberScopeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateContactFunctionalRoleMemberScope mock ContactUpdateContactFunctionalRoleMemberScope method
func (r *Mock) MockContactUpdateContactFunctionalRoleMemberScope(f func(ctx context.Context, request *UpdateContactFunctionalRoleMemberScopeReq, options ...MethodOptionFunc) (*UpdateContactFunctionalRoleMemberScopeResp, *Response, error)) {
	r.mockContactUpdateContactFunctionalRoleMemberScope = f
}

// UnMockContactUpdateContactFunctionalRoleMemberScope un-mock ContactUpdateContactFunctionalRoleMemberScope method
func (r *Mock) UnMockContactUpdateContactFunctionalRoleMemberScope() {
	r.mockContactUpdateContactFunctionalRoleMemberScope = nil
}

// UpdateContactFunctionalRoleMemberScopeReq ...
type UpdateContactFunctionalRoleMemberScopeReq struct {
	RoleID           string            `path:"role_id" json:"-"`             // 角色的唯一标识, 单租户下唯一, 示例值: "7vrj3vk70xk7v5r"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: open_department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	Members          []string          `json:"members,omitempty"`            // 角色修改的角色成员列表（一批用户的UserID列表), 示例值: ["ou-12832197382"], 长度范围: `1` ～ `100`
	Departments      []string          `json:"departments,omitempty"`        // 角色内用户的管理范围, 示例值: ["ou-12343455"], 长度范围: `1` ～ `100`
}

// UpdateContactFunctionalRoleMemberScopeResp ...
type UpdateContactFunctionalRoleMemberScopeResp struct {
	Results []*UpdateContactFunctionalRoleMemberScopeRespResult `json:"results,omitempty"` // 批量更新角色成员管理范围结果集
}

// UpdateContactFunctionalRoleMemberScopeRespResult ...
type UpdateContactFunctionalRoleMemberScopeRespResult struct {
	UserID string `json:"user_id,omitempty"` // 用户ID
	Reason int64  `json:"reason,omitempty"`  // 成员处理结果, 可选值有: 1: 处理成功, 2: 用户ID无效, 3: 用户ID无权限, 4: 用户已存在在该角色中, 5: 用户不存在在该角色中, 6: 对该角色内该用户旧的管理范围无权限
}

// updateContactFunctionalRoleMemberScopeResp ...
type updateContactFunctionalRoleMemberScopeResp struct {
	Code int64                                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateContactFunctionalRoleMemberScopeResp `json:"data,omitempty"`
}
