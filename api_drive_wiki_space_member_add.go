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

// AddWikiSpaceMember 添加知识空间成员或管理员。
//
// 知识空间具有[类型](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)和[可见性](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)的概念。不同的类型或可见性可以对本操作做出限制:
// - 可见性限制: 公开知识空间（visibility为public）对租户所有用户可见, 因此不支持再添加成员, 但可以添加管理员。
// - 类型限制: 个人知识空间 （type为person）为个人管理的知识空间, 不支持添加其他管理员（包括应用/机器人）。但可以添加成员。
// 知识空间权限要求, 当前用户或应用:
// - 为知识空间管理员
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/create
// new doc: https://open.feishu.cn/document/server-docs/docs/wiki-v2/space-member/create
func (r *DriveService) AddWikiSpaceMember(ctx context.Context, request *AddWikiSpaceMemberReq, options ...MethodOptionFunc) (*AddWikiSpaceMemberResp, *Response, error) {
	if r.cli.mock.mockDriveAddWikiSpaceMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#AddWikiSpaceMember mock enable")
		return r.cli.mock.mockDriveAddWikiSpaceMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "AddWikiSpaceMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/spaces/:space_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(addWikiSpaceMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveAddWikiSpaceMember mock DriveAddWikiSpaceMember method
func (r *Mock) MockDriveAddWikiSpaceMember(f func(ctx context.Context, request *AddWikiSpaceMemberReq, options ...MethodOptionFunc) (*AddWikiSpaceMemberResp, *Response, error)) {
	r.mockDriveAddWikiSpaceMember = f
}

// UnMockDriveAddWikiSpaceMember un-mock DriveAddWikiSpaceMember method
func (r *Mock) UnMockDriveAddWikiSpaceMember() {
	r.mockDriveAddWikiSpaceMember = nil
}

// AddWikiSpaceMemberReq ...
type AddWikiSpaceMemberReq struct {
	SpaceID          string `path:"space_id" json:"-"`           // 知识空间id, 示例值: "1565676577122621"
	NeedNotification *bool  `query:"need_notification" json:"-"` // 添加权限后是否通知对方, 示例值: true
	MemberType       string `json:"member_type,omitempty"`       // “openchat” - 群id, “userid” - 用户id, “email” - 邮箱, “opendepartmentid” - 部门id, “openid” - 应用openid, “unionid” - [unionid](/:ssltoken/home/user-identity-introduction/union-id, ), 示例值: "userid"
	MemberID         string `json:"member_id,omitempty"`         // 用户id, 值的类型由上面的 member_type 参数决定, 示例值: "1565676577122621"
	MemberRole       string `json:"member_role,omitempty"`       // 角色: “admin” - 管理员, “member” - 成员, 示例值: "admin"
}

// AddWikiSpaceMemberResp ...
type AddWikiSpaceMemberResp struct {
	Member *AddWikiSpaceMemberRespMember `json:"member,omitempty"` // 知识空间成员
}

// AddWikiSpaceMemberRespMember ...
type AddWikiSpaceMemberRespMember struct {
	MemberType string `json:"member_type,omitempty"` // “openchat” - 群id, “userid” - 用户id, “email” - 邮箱, “opendepartmentid” - 部门id, “openid” - 应用openid, “unionid” - [unionid](/:ssltoken/home/user-identity-introduction/union-id, )
	MemberID   string `json:"member_id,omitempty"`   // 用户id, 值的类型由上面的 member_type 参数决定
	MemberRole string `json:"member_role,omitempty"` // 角色: “admin” - 管理员, “member” - 成员
}

// addWikiSpaceMemberResp ...
type addWikiSpaceMemberResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *AddWikiSpaceMemberResp `json:"data,omitempty"`
}
