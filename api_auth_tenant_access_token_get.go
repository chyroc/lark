package lark

import (
	"context"
	"time"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
func (r *AuthService) GetTenantAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	if r.cli.mock.mockGetTenantAccessToken != nil {
		r.cli.logDebug(ctx, "[lark] Auth#GetTenantAccessToken mock enable")
		return r.cli.mock.mockGetTenantAccessToken(ctx)
	}

	r.cli.logInfo(ctx, "[lark] Auth#GetTenantAccessToken call api")

	val, ttl, err := r.cli.store.Get(ctx, genTenantTokenKey(r.cli.isISV, r.cli.appID, r.cli.tenantKey))
	if err != nil && err != ErrStoreNotFound {
		r.cli.logError(ctx, "[lark] Auth#GetTenantAccessToken get token from store failed: %s", err)
	} else if val != "" && ttl > 0 {
		return &TokenExpire{Token: val, Expire: int(ttl.Seconds())}, &Response{}, nil
	}

	uri := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	body := getTenantAccessTokenReq{
		AppID:     r.cli.appID,
		AppSecret: r.cli.appSecret,
	}
	if r.cli.isISV {
		appAccessToken, response, err := r.GetAppAccessToken(ctx)
		if err != nil {
			return nil, response, err
		}
		uri = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/"
		body = getTenantAccessTokenReq{
			AppAccessToken: appAccessToken.Token,
			TenantKey:      r.cli.tenantKey,
		}
	}
	req := &RawRequestReq{
		Method: "POST",
		URL:    uri,
		Body:   body,
	}
	resp := new(getTenantAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Auth#GetTenantAccessToken GET %s failed: %s", uri, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Auth#GetTenantAccessToken GET %s failed, code: %d, msg: %s", uri, resp.Code, resp.Msg)
		return nil, response, NewError("Token", "GetTenantAccessToken", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Auth#GetTenantAccessToken request_id: %s, response: %s", response.RequestID, jsonString(resp))

	err = r.cli.store.Set(ctx, genTenantTokenKey(r.cli.isISV, r.cli.appID, r.cli.tenantKey), resp.AppAccessToken, time.Second*time.Duration(resp.Expire))
	if err != nil {
		r.cli.logError(ctx, "[lark] Auth#GetTenantAccessToken set token to store failed: %s", err)
	}

	return &TokenExpire{
		Token:  resp.TenantAccessToken,
		Expire: resp.Expire,
	}, response, nil
}

func (r *Mock) MockGetTenantAccessToken(f func(ctx context.Context) (*TokenExpire, *Response, error)) {
	r.mockGetTenantAccessToken = f
}

func (r *Mock) UnMockGetTenantAccessToken() {
	r.mockGetTenantAccessToken = nil
}

type TokenExpire struct {
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}

type getTenantAccessTokenReq struct {
	AppID          string `json:"app_id,omitempty"`
	AppSecret      string `json:"app_secret,omitempty"`
	AppAccessToken string `json:"app_access_token,omitempty"`
	TenantKey      string `json:"tenant_key,omitempty"`
	AppTicket      string `json:"app_ticket,omitempty"`
}

type getTenantAccessTokenResp struct {
	Code              int    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg               string `json:"msg,omitempty"`  // 错误描述
	TenantAccessToken string `json:"tenant_access_token"`
	AppAccessToken    string `json:"app_access_token"`
	Expire            int    `json:"expire"`
}
