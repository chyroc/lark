package lark

import (
	"context"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
func (r *TokenAPI) GetAppAccessToken(ctx context.Context) (*TokenExpire, *Response, error) {
	url := "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/"
	body := getTenantAccessTokenReq{
		AppID:     r.cli.appID,
		AppSecret: r.cli.appSecret,
	}

	if r.cli.tenantKey != "" {
		url = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/"
		body = getTenantAccessTokenReq{
			AppID:     r.cli.appID,
			AppSecret: r.cli.appSecret,
			AppTicket: "",
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
		return nil, response, newError("Token", "GetAppAccessToken", resp.Code, resp.Msg)
	}

	return &TokenExpire{
		Token:  resp.AppAccessToken,
		Expire: resp.Expire,
	}, response, nil
}
