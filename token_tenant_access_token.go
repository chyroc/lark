package lark

import (
	"context"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
func (r *TokenAPI) GetTenantAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	if r.cli.mock.mockGetTenantAccessToken != nil {
		return r.cli.mock.mockGetTenantAccessToken(ctx)
	}

	appToken, res, err := r.GetAppAccessToken(ctx)
	if err != nil {
		return nil, res, err
	}

	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	body := getTenantAccessTokenReq{
		AppID:     r.cli.appID,
		AppSecret: r.cli.appSecret,
	}
	if r.cli.tenantKey != "" {
		url = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/"
		body = getTenantAccessTokenReq{
			TenantKey: r.cli.tenantKey,
			AppSecret: appToken.Token,
		}
	}

	req := &requestParam{
		Method: "POST",
		URL:    url,
		Body:   body,
	}
	resp := new(getTenantAccessTokenResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Token", "GetTenantAccessToken", resp.Code, resp.Msg)
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
	AppTicket      string `json:"app_ticket,omitempty"`       // 平台定时推送给应用的临时凭证，通过事件监听机制获得，详见订阅事件
	TenantKey      string `json:"tenant_key,omitempty"`       // 企业标识，两种获取方式，企业开通应用时由平台方推送给应用 / 用户登录时返回，示例：73658811060f175d
	AppAccessToken string `json:"app_access_token,omitempty"` // app 凭证，示例：a-32bd8551db2f081cbfd26293f27516390b9feb04
}

type getTenantAccessTokenResp struct {
	Code              int    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg               string `json:"msg,omitempty"`  // 错误描述
	TenantAccessToken string `json:"tenant_access_token"`
	AppAccessToken    string `json:"app_access_token"`
	Expire            int    `json:"expire"`
}
