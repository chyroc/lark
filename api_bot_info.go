package lark

import (
	"context"
)

// GetBotInfo 获取机器人的基本信息。
//
// 需要启用机器人能力
// https://open.feishu.cn/document/ukTMukTMukTM/uAjMxEjLwITMx4CMyETM
func (r *BotAPI) GetBotInfo(ctx context.Context, request *GetBotInfoReq) (*GetBotInfoResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/bot/v3/info",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getBotInfoResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Bot", "GetBotInfo", resp.Code, resp.Msg)
	}

	return resp.Bot, response, nil
}

type GetBotInfoReq struct{}

type getBotInfoResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Bot  *GetBotInfoResp `json:"bot,omitempty"`  // 机器人信息
}

type GetBotInfoResp struct {
	ActivateStatus int      `json:"activate_status,omitempty"` // app 当前状态。,0: 初始化，租户待安装,1: 租户停用,2: 租户启用,3: 安装后待启用,4: 升级待启用,5: license过期停用,6: Lark套餐到期或降级停用
	AppName        string   `json:"app_name,omitempty"`        // app 名称
	AvatarUrl      string   `json:"avatar_url,omitempty"`      // app 图像地址
	IpWhiteList    []string `json:"ip_white_list,omitempty"`   // app 的 IP 白名单地址
	OpenID         string   `json:"open_id,omitempty"`         // 机器人的open_id
}
