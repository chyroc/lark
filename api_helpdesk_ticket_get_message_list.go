package lark

import (
	"context"
)

// GetTicketMessageList 该接口用于获取服务台工单消息详情。
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket-message/list
func (r *HelpdeskAPI) GetTicketMessageList(ctx context.Context, request *GetTicketMessageListReq) (*GetTicketMessageListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id/messages",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      true,
		IsFile:                false,
	}
	resp := new(getTicketMessageListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Helpdesk", "GetTicketMessageList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetTicketMessageListReq struct {
	TimeStart *int   `query:"time_start" json:"-"` // 起始时间,**示例值**：1617960686000
	TimeEnd   *int   `query:"time_end" json:"-"`   // 结束时间,**示例值**：1617960687000
	Page      *int   `query:"page" json:"-"`       // 页数ID,**示例值**：1
	PageSize  *int   `query:"page_size" json:"-"`  // 消息数量，最大200，默认20,**示例值**：10
	TicketID  string `path:"ticket_id" json:"-"`   // 工单ID,**示例值**："6948728206392295444"
}

type getTicketMessageListResp struct {
	Code int                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetTicketMessageListResp `json:"data,omitempty"` // \-
}

type GetTicketMessageListResp struct {
	Messages []*GetTicketMessageListRespMessage `json:"messages,omitempty"` // 工单消息列表
	Total    int                                `json:"total,omitempty"`    // 消息总数
}

type GetTicketMessageListRespMessage struct {
	ID          string `json:"id,omitempty"`           // 工单消息ID
	MessageID   string `json:"message_id,omitempty"`   // chat消息ID
	MessageType string `json:"message_type,omitempty"` // 消息类型；text：纯文本；post：富文本
	CreatedAt   int    `json:"created_at,omitempty"`   // 创建时间
	Content     string `json:"content,omitempty"`      // 内容
	UserName    string `json:"user_name,omitempty"`    // 用户名
	AvatarUrl   string `json:"avatar_url,omitempty"`   // 用户图片url
	UserID      string `json:"user_id,omitempty"`      // 用户open ID
}
