package lark

import (
	"context"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
func (r *TokenAPI) GetAppAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	req := &requestParam{
		Method: "POST",
		URL:    "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/",
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
		return nil, response, newError("Token", "GetAppAccessToken", resp.Code, resp.Msg)
	}

	return &TokenExpire{
		Token:  resp.AppAccessToken,
		Expire: resp.Expire,
	}, response, nil
}
