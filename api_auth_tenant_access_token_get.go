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
	"time"
)

// GetTenantAccessToken ...
//
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token
func (r *AuthService) GetTenantAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	if r.cli.mock.mockGetTenantAccessToken != nil {
		r.cli.Log(ctx, LogLevelTrace, "[lark] Auth#GetTenantAccessToken mock enable")
		return r.cli.mock.mockGetTenantAccessToken(ctx)
	}

	val, ttl, err := r.cli.store.Get(ctx, genTenantTokenKey(r.cli.isISV, r.cli.appID, r.cli.tenantKey))
	if err != nil && err != ErrStoreNotFound {
		r.cli.Log(ctx, LogLevelError, "[lark] Auth#GetTenantAccessToken get_from_cache failed, err=%s", err)
	} else if val != "" && ttl > 0 {
		return &TokenExpire{Token: val, Expire: int64(ttl.Seconds())}, &Response{}, nil
	}

	uri := r.cli.openBaseURL + "/open-apis/auth/v3/tenant_access_token/internal"
	body := getTenantAccessTokenReq{
		AppID:     r.cli.appID,
		AppSecret: r.cli.appSecret,
	}
	if r.cli.isISV {
		appAccessToken, response, err := r.GetAppAccessToken(ctx)
		if err != nil {
			return nil, response, err
		}
		uri = r.cli.openBaseURL + "/open-apis/auth/v3/tenant_access_token"
		body = getTenantAccessTokenReq{
			AppAccessToken: appAccessToken.Token,
			TenantKey:      r.cli.tenantKey,
		}
	}
	req := &RawRequestReq{
		Scope:        "Auth",
		API:          "GetTenantAccessToken",
		Method:       "POST",
		URL:          uri,
		Body:         body,
		MethodOption: newMethodOption(nil),
	}
	resp := new(getTenantAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	}

	if err = r.cli.store.Set(ctx, genTenantTokenKey(r.cli.isISV, r.cli.appID, r.cli.tenantKey), resp.TenantAccessToken, time.Second*time.Duration(resp.Expire)); err != nil {
		r.cli.Log(ctx, LogLevelError, "[lark] Auth#GetTenantAccessToken set_token_cache failed, err=%s", err)
	}
	return &TokenExpire{
		Token:  resp.TenantAccessToken,
		Expire: resp.Expire,
	}, response, nil
}

// MockGetTenantAccessToken ...
func (r *Mock) MockGetTenantAccessToken(f func(ctx context.Context) (*TokenExpire, *Response, error)) {
	r.mockGetTenantAccessToken = f
}

// UnMockGetTenantAccessToken ...
func (r *Mock) UnMockGetTenantAccessToken() {
	r.mockGetTenantAccessToken = nil
}

// TokenExpire ...
type TokenExpire struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type getTenantAccessTokenReq struct {
	AppID          string `json:"app_id,omitempty"`
	AppSecret      string `json:"app_secret,omitempty"`
	AppAccessToken string `json:"app_access_token,omitempty"`
	TenantKey      string `json:"tenant_key,omitempty"`
	AppTicket      string `json:"app_ticket,omitempty"`
}

type getTenantAccessTokenResp struct {
	Code              int64        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg               string       `json:"msg,omitempty"`  // 错误描述
	TenantAccessToken string       `json:"tenant_access_token"`
	AppAccessToken    string       `json:"app_access_token"`
	Expire            int64        `json:"expire"`
	Error             *ErrorDetail `json:"error,omitempty"`
}
