package lark

import (
	"context"
	"time"
)

// docs: https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token_internal
// docs: https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
func (r *AuthService) GetAppAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	if r.cli.mock.mockGetAppAccessToken != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Auth#GetAppAccessToken mock enable")
		return r.cli.mock.mockGetAppAccessToken(ctx)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Auth#GetAppAccessToken call api")

	val, ttl, err := r.cli.store.Get(ctx, genAppTokenKey(r.cli.isISV, r.cli.appID))
	if err != nil && err != ErrStoreNotFound {
		r.cli.log(ctx, LogLevelError, "[lark] Auth#GetAppAccessToken get token from store failed: %s", err)
	} else if val != "" && ttl > 0 {
		return &TokenExpire{Token: val, Expire: int64(ttl.Seconds())}, &Response{}, nil
	}

	uri := r.cli.openBaseURL + "/open-apis/auth/v3/app_access_token/internal"
	appTicket := ""
	if r.cli.isISV {
		uri = r.cli.openBaseURL + "/open-apis/auth/v3/app_access_token"
		s, err := r.GetAppTicket(ctx)
		if err != nil {
			return nil, nil, err
		}
		appTicket = s
	}

	req := &RawRequestReq{
		Scope:  "Auth",
		API:    "GetAppAccessToken",
		Method: "POST",
		URL:    uri,
		Body: getTenantAccessTokenReq{
			AppID:     r.cli.appID,
			AppSecret: r.cli.appSecret,
			AppTicket: appTicket,
		},
	}
	resp := new(getTenantAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Auth#GetAppAccessToken POST %s failed: %s", uri, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Auth#GetAppAccessToken POST %s failed, code: %d, msg: %s", uri, resp.Code, resp.Msg)
		return nil, response, NewError("Token", "GetAppAccessToken", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Auth#GetAppAccessToken request_id: %s, response: %s", response.RequestID, jsonString(resp))

	err = r.cli.store.Set(ctx, genAppTokenKey(r.cli.isISV, r.cli.appID), resp.AppAccessToken, time.Second*time.Duration(resp.Expire))
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Auth#GetAppAccessToken set token to store failed: %s", err)
	}

	return &TokenExpire{
		Token:  resp.AppAccessToken,
		Expire: resp.Expire,
	}, response, nil
}

// MockGetAppAccessToken ...
func (r *Mock) MockGetAppAccessToken(f func(ctx context.Context) (*TokenExpire, *Response, error)) {
	r.mockGetAppAccessToken = f
}

// UnMockGetAppAccessToken ...
func (r *Mock) UnMockGetAppAccessToken() {
	r.mockGetTenantAccessToken = nil
}

// lark 中有下面这些 token，每个 token 的决定因素不一样，决定了各自的 store key 组成方式
// tenant-token：app_id
// app-token：app_id
// isv-tenant-token：app-id + tenant_key
// isv app-token：app_id
// app-ticket: app_id
// user-token，不需要 store，通过 method option 传入

func genTenantTokenKey(isISV bool, appID, tenantKey string) string {
	if isISV {
		return "isv-tenant-token:" + appID + "-" + tenantKey
	}
	return "internal-tenant-token:" + appID
}

func genAppTokenKey(isISV bool, appID string) string {
	if isISV {
		return "isv-app-token:" + appID
	}
	return "internal-app-token:" + appID
}

func genISVAppTicketKey(appID string) string {
	return "isv-app-ticket:" + appID
}
