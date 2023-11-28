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

// GetAccessToken 根据[登录预授权码](https://open.feishu.cn/document/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN) code 获取 `user_access_token`。
//
// 本接口用于网页应用免登录应用场景, 小程序应用获取 user_access_token 的方法, 请参考小程序应用提供的 [code2session](https://open.feishu.cn/document/uYjL24iN/ukjM04SOyQjL5IDN) 接口。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/access_token/create
// new doc: https://open.feishu.cn/document/server-docs/authentication-management/access-token/create-2
//
// Deprecated
func (r *AuthService) GetAccessToken(ctx context.Context, request *GetAccessTokenReq, options ...MethodOptionFunc) (*GetAccessTokenResp, *Response, error) {
	if r.cli.mock.mockAuthGetAccessToken != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Auth#GetAccessToken mock enable")
		return r.cli.mock.mockAuthGetAccessToken(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:              "Auth",
		API:                "GetAccessToken",
		Method:             "POST",
		URL:                r.cli.openBaseURL + "/open-apis/authen/v1/access_token",
		Body:               request,
		MethodOption:       newMethodOption(options),
		NeedAppAccessToken: true,
	}
	resp := new(getAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAuthGetAccessToken mock AuthGetAccessToken method
func (r *Mock) MockAuthGetAccessToken(f func(ctx context.Context, request *GetAccessTokenReq, options ...MethodOptionFunc) (*GetAccessTokenResp, *Response, error)) {
	r.mockAuthGetAccessToken = f
}

// UnMockAuthGetAccessToken un-mock AuthGetAccessToken method
func (r *Mock) UnMockAuthGetAccessToken() {
	r.mockAuthGetAccessToken = nil
}

// GetAccessTokenReq ...
type GetAccessTokenReq struct {
	GrantType string `json:"grant_type,omitempty"` // 授权类型, 固定值: 示例值: "authorization_code"
	Code      string `json:"code,omitempty"`       // 登录预授权码, 调用[获取登录预授权码](https://open.feishu.cn/document/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN)接口获取, 示例值: "xMSldislSkdK"
}

// GetAccessTokenResp ...
type GetAccessTokenResp struct {
	AccessToken      string `json:"access_token,omitempty"`       // user_access_token, 用于获取用户资源
	TokenType        string `json:"token_type,omitempty"`         // token 类型
	ExpiresIn        int64  `json:"expires_in,omitempty"`         // `access_token` 的有效期, 单位: 秒
	Name             string `json:"name,omitempty"`               // 用户姓名
	EnName           string `json:"en_name,omitempty"`            // 用户英文名称
	AvatarURL        string `json:"avatar_url,omitempty"`         // 用户头像
	AvatarThumb      string `json:"avatar_thumb,omitempty"`       // 用户头像 72x72
	AvatarMiddle     string `json:"avatar_middle,omitempty"`      // 用户头像 240x240
	AvatarBig        string `json:"avatar_big,omitempty"`         // 用户头像 640x640
	OpenID           string `json:"open_id,omitempty"`            // 用户在应用内的唯一标识
	UnionID          string `json:"union_id,omitempty"`           // 用户统一ID
	Email            string `json:"email,omitempty"`              // 用户邮箱, 字段权限要求: 获取用户邮箱信息
	EnterpriseEmail  string `json:"enterprise_email,omitempty"`   // 企业邮箱, 请先确保已在管理后台启用飞书邮箱服务, 字段权限要求: 获取用户受雇信息
	UserID           string `json:"user_id,omitempty"`            // 用户 user_id, 字段权限要求: 获取用户 user ID
	Mobile           string `json:"mobile,omitempty"`             // 用户手机号, 字段权限要求: 获取用户手机号
	TenantKey        string `json:"tenant_key,omitempty"`         // 当前企业标识
	RefreshExpiresIn int64  `json:"refresh_expires_in,omitempty"` // `refresh_token` 的有效期, 单位: 秒
	RefreshToken     string `json:"refresh_token,omitempty"`      // 刷新用户 `access_token` 时使用的 token
	Sid              string `json:"sid,omitempty"`                // 用户当前登录态session的唯一标识, 为空则不返回
}

// getAccessTokenResp ...
type getAccessTokenResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetAccessTokenResp `json:"data,omitempty"`
}
