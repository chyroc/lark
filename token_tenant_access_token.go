package lark

import (
	"context"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
func (r *TokenAPI) GetTenantAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	req := &requestParam{
		Method: "POST",
		URL:    "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/",
		Body: getTenantAccessTokenReq{
			AppID:     r.cli.appID,
			AppSecret: r.cli.appSecret,
		},
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

type TokenExpire struct {
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}

type getTenantAccessTokenReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type getTenantAccessTokenResp struct {
	Code              int    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg               string `json:"msg,omitempty"`  // 错误描述
	TenantAccessToken string `json:"tenant_access_token"`
	AppAccessToken    string `json:"app_access_token"`
	Expire            int    `json:"expire"`
}
