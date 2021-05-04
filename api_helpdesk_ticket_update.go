package lark

import (
	"context"
)

// UpdateTicket 该接口用于更新服务台工单详情。只会更新数据，不会触发相关操作。如修改工单状态到关单，不会关闭聊天页面。仅支持自建应用。
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/update
func (r *HelpDeskAPI) UpdateTicket(ctx context.Context, request *UpdateTicketReq) (*UpdateTicketResp, *Response, error) {
	req := &requestParam{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(updateTicketResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("HelpDesk", "UpdateTicket", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateTicketReq struct {
	TicketID         string                            `path:"ticket_id" json:"-"`          // 工单ID,**示例值**："6945345902185807891"
	Status           *int                              `json:"status,omitempty"`            // new status, 1: 已创建, 2: 处理中, 3: 排队中, 5: 待定, 50: 机器人关闭工单, 51: 关闭工单,**示例值**：1
	TagNames         []string                          `json:"tag_names,omitempty"`         // 新标签名
	Comment          *string                           `json:"comment,omitempty"`           // 新评论,**示例值**："good"
	CustomizedFields []*UpdateTicketReqCustomizedField `json:"customized_fields,omitempty"` // 自定义字段
	TicketType       *int                              `json:"ticket_type,omitempty"`       // ticket stage,**示例值**：1
	Solved           *int                              `json:"solved,omitempty"`            // 工单是否解决，1: 未解决, 2: 已解决,**示例值**：1
	Channel          *int                              `json:"channel,omitempty"`           // 工单来源渠道ID,**示例值**：1
}

type UpdateTicketReqCustomizedField struct {
	Id      *string `json:"id,omitempty"`       // id,**示例值**："123"
	Value   *string `json:"value,omitempty"`    // value,**示例值**："value"
	KeyName *string `json:"key_name,omitempty"` // key name,**示例值**："key"
}

type updateTicketResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *UpdateTicketResp `json:"data,omitempty"`
}

type UpdateTicketResp struct{}
