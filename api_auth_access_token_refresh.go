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

// RefreshAccessToken user_access_token 的最大有效期是 2小时左右。当 user_access_token 过期时, 可以调用本接口获取新的 user_access_token。
//
// 刷新后请更新本地user_access_token和refresh_token, 不要继续使用旧值重复刷新。保证参数是最新值
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/oidc-refresh_access_token/create
func (r *AuthService) RefreshAccessToken(ctx context.Context, request *RefreshAccessTokenReq, options ...MethodOptionFunc) (*RefreshAccessTokenResp, *Response, error) {
	if r.cli.mock.mockAuthRefreshAccessToken != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Auth#RefreshAccessToken mock enable")
		return r.cli.mock.mockAuthRefreshAccessToken(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:              "Auth",
		API:                "RefreshAccessToken",
		Method:             "POST",
		URL:                r.cli.openBaseURL + "/open-apis/authen/v1/oidc/refresh_access_token",
		Body:               request,
		MethodOption:       newMethodOption(options),
		NeedAppAccessToken: true,
	}
	resp := new(refreshAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAuthRefreshAccessToken mock AuthRefreshAccessToken method
func (r *Mock) MockAuthRefreshAccessToken(f func(ctx context.Context, request *RefreshAccessTokenReq, options ...MethodOptionFunc) (*RefreshAccessTokenResp, *Response, error)) {
	r.mockAuthRefreshAccessToken = f
}

// UnMockAuthRefreshAccessToken un-mock AuthRefreshAccessToken method
func (r *Mock) UnMockAuthRefreshAccessToken() {
	r.mockAuthRefreshAccessToken = nil
}

// RefreshAccessTokenReq ...
type RefreshAccessTokenReq struct {
	GrantType    string `json:"grant_type,omitempty"`    // 授权类型, 固定值: 示例值: "refresh_token"
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新和获取user_access_token接口均返回 `refresh_token`, 每次请求, 请注意使用最新获取到的`refresh_token`, 示例值: "ur-oQ0mMq6MCcueAv0pwx2fQQhxqv__CbLu6G8ySFwafeKww2Def2BJdOkW3.9gCFM.LBQgFri901QaqeuL"
}

// RefreshAccessTokenResp ...
type RefreshAccessTokenResp struct {
	AccessToken      string `json:"access_token,omitempty"`       // 字段`access_token`即user_access_token, 用于获取用户资源和访问某些open api
	RefreshToken     string `json:"refresh_token,omitempty"`      // 刷新 user_access_token时使用的 refresh_token
	TokenType        string `json:"token_type,omitempty"`         // token 类型, 固定值
	ExpiresIn        int64  `json:"expires_in,omitempty"`         // user_access_token有效期, 单位: 秒, 有效时间两个小时左右, 需要以返回结果为准
	RefreshExpiresIn int64  `json:"refresh_expires_in,omitempty"` // refresh_token有效期, 单位: 秒, 一般是30天左右, 需要以返回结果为准
	Scope            string `json:"scope,omitempty"`              // 用户授予app的权限全集
}

// refreshAccessTokenResp ...
type refreshAccessTokenResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *RefreshAccessTokenResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
