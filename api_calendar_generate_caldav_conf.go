package lark

import (
	"context"
)

// GenerateCaldavConf 用于为当前用户生成一个CalDAV账号密码。
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf
func (r *CalendarAPI) GenerateCaldavConf(ctx context.Context, request *GenerateCaldavConfReq) (*GenerateCaldavConfResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/settings/generate_caldav_conf",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(generateCaldavConfResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "GenerateCaldavConf", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GenerateCaldavConfReq struct {
	DeviceName *string `json:"device_name,omitempty"` // 需要同步日历的设备名，在日历中展示用来管理密码,**示例值**："iPhone",**数据校验规则**：,- 最大长度：`100` 字符
}

type generateCaldavConfResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GenerateCaldavConfResp `json:"data,omitempty"` // \-
}

type GenerateCaldavConfResp struct {
	Password      string `json:"password,omitempty"`       // caldav密码
	UserName      string `json:"user_name,omitempty"`      // caldav用户名
	ServerAddress string `json:"server_address,omitempty"` // 服务器地址
	DeviceName    string `json:"device_name,omitempty"`    // 设备名
}
