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

// DeleteAdminBadgeGrant 通过该接口可以删除特定授予名单的信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge-grant/delete
// new doc: https://open.feishu.cn/document/server-docs/admin-v1/badge/badge-grant/delete
func (r *AdminService) DeleteAdminBadgeGrant(ctx context.Context, request *DeleteAdminBadgeGrantReq, options ...MethodOptionFunc) (*DeleteAdminBadgeGrantResp, *Response, error) {
	if r.cli.mock.mockAdminDeleteAdminBadgeGrant != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#DeleteAdminBadgeGrant mock enable")
		return r.cli.mock.mockAdminDeleteAdminBadgeGrant(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "DeleteAdminBadgeGrant",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/admin/v1/badges/:badge_id/grants/:grant_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteAdminBadgeGrantResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAdminDeleteAdminBadgeGrant mock AdminDeleteAdminBadgeGrant method
func (r *Mock) MockAdminDeleteAdminBadgeGrant(f func(ctx context.Context, request *DeleteAdminBadgeGrantReq, options ...MethodOptionFunc) (*DeleteAdminBadgeGrantResp, *Response, error)) {
	r.mockAdminDeleteAdminBadgeGrant = f
}

// UnMockAdminDeleteAdminBadgeGrant un-mock AdminDeleteAdminBadgeGrant method
func (r *Mock) UnMockAdminDeleteAdminBadgeGrant() {
	r.mockAdminDeleteAdminBadgeGrant = nil
}

// DeleteAdminBadgeGrantReq ...
type DeleteAdminBadgeGrantReq struct {
	BadgeID string `path:"badge_id" json:"-"` // 企业勋章的唯一ID, 示例值: "m_DjMzaK", 长度范围: `1` ～ `64` 字符
	GrantID string `path:"grant_id" json:"-"` // 租户内授予名单的唯一标识, 该值由系统随机生成, 示例值: "g_uS4yux", 长度范围: `1` ～ `64` 字符
}

// DeleteAdminBadgeGrantResp ...
type DeleteAdminBadgeGrantResp struct {
}

// deleteAdminBadgeGrantResp ...
type deleteAdminBadgeGrantResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteAdminBadgeGrantResp `json:"data,omitempty"`
}
