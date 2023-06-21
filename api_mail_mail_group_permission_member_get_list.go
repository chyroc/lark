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

// GetMailGroupPermissionMemberList 分页批量获取邮件组权限成员列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/list
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/mail-group/mailgroup-permission_member/list
func (r *MailService) GetMailGroupPermissionMemberList(ctx context.Context, request *GetMailGroupPermissionMemberListReq, options ...MethodOptionFunc) (*GetMailGroupPermissionMemberListResp, *Response, error) {
	if r.cli.mock.mockMailGetMailGroupPermissionMemberList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetMailGroupPermissionMemberList mock enable")
		return r.cli.mock.mockMailGetMailGroupPermissionMemberList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailGroupPermissionMemberList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailGroupPermissionMemberListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailGroupPermissionMemberList mock MailGetMailGroupPermissionMemberList method
func (r *Mock) MockMailGetMailGroupPermissionMemberList(f func(ctx context.Context, request *GetMailGroupPermissionMemberListReq, options ...MethodOptionFunc) (*GetMailGroupPermissionMemberListResp, *Response, error)) {
	r.mockMailGetMailGroupPermissionMemberList = f
}

// UnMockMailGetMailGroupPermissionMemberList un-mock MailGetMailGroupPermissionMemberList method
func (r *Mock) UnMockMailGetMailGroupPermissionMemberList() {
	r.mockMailGetMailGroupPermissionMemberList = nil
}

// GetMailGroupPermissionMemberListReq ...
type GetMailGroupPermissionMemberListReq struct {
	MailGroupID      string            `path:"mailgroup_id" json:"-"`        // 邮件组ID或者邮件组地址, 示例值: "xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "xxx"
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `200`
}

// GetMailGroupPermissionMemberListResp ...
type GetMailGroupPermissionMemberListResp struct {
	HasMore   bool                                        `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                      `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetMailGroupPermissionMemberListRespItem `json:"items,omitempty"`      // 邮件组权限成员列表
}

// GetMailGroupPermissionMemberListRespItem ...
type GetMailGroupPermissionMemberListRespItem struct {
	PermissionMemberID string       `json:"permission_member_id,omitempty"` // 权限组内成员唯一标识
	UserID             string       `json:"user_id,omitempty"`              // 租户内用户的唯一标识（当成员类型是USER时有值）
	DepartmentID       string       `json:"department_id,omitempty"`        // 租户内部门的唯一标识（当成员类型是DEPARTMENT时有值）
	Email              string       `json:"email,omitempty"`                // 成员邮箱地址（当成员类型是MAIL_GROUP/PUBLIC_MAILBOX时有值）
	Type               MailUserType `json:"type,omitempty"`                 // 成员类型, 可选值有: USER: 内部用户, DEPARTMENT: 部门, MAIL_GROUP: 邮件组, PUBLIC_MAILBOX: 公共邮箱
}

// getMailGroupPermissionMemberListResp ...
type getMailGroupPermissionMemberListResp struct {
	Code int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupPermissionMemberListResp `json:"data,omitempty"`
}
