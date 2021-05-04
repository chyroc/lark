package lark

import (
	"context"
)

// StartService 该接口用于创建服务台对话。
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/start_service
func (r *HelpDeskAPI) StartService(ctx context.Context, request *StartServiceReq) (*StartServiceResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/start_service",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(startServiceResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("HelpDesk", "StartService", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type StartServiceReq struct {
	HumanService    *bool    `json:"human_service,omitempty"`    // 是否直接进入人工,**示例值**：false
	AppointedAgents []string `json:"appointed_agents,omitempty"` // 客服 open ids (获取方式参考[获取单个用户信息](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get))，human_service需要为true
	OpenID          string   `json:"open_id,omitempty"`          // 用户 open id,(获取方式参考[获取单个用户信息](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/get)),**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	CustomizedInfo  *string  `json:"customized_info,omitempty"`  // 工单来源自定义信息，长度限制1024字符，如设置，[获取工单详情](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get)会返回此信息,**示例值**："test customized info"
}

type startServiceResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *StartServiceResp `json:"data,omitempty"` // \-
}

type StartServiceResp struct {
	ChatID   string `json:"chat_id,omitempty"`   // 客服群open ID
	TicketID string `json:"ticket_id,omitempty"` // 工单ID
}
