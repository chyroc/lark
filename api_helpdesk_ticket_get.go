package lark

import (
	"context"
)

// GetTicket 该接口用于获取单个服务台工单详情。仅支持自建应用。
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/get
func (r *HelpdeskAPI) GetTicket(ctx context.Context, request *GetTicketReq) (*GetTicketResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets/:ticket_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(getTicketResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Helpdesk", "GetTicket", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetTicketReq struct {
	TicketID string `path:"ticket_id" json:"-"` // ticket id,**示例值**："123456"
}

type getTicketResp struct {
	Code int            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *GetTicketResp `json:"data,omitempty"` // \-
}

type GetTicketResp struct {
	Ticket *GetTicketRespTicket `json:"ticket,omitempty"` // 工单详情
}

type GetTicketRespTicket struct {
	TicketID      string                             `json:"ticket_id,omitempty"`     // 工单ID
	HelpdeskID    string                             `json:"helpdesk_id,omitempty"`   // 服务台ID
	Guest         *GetTicketRespTicketGuest          `json:"guest,omitempty"`         // 工单创建用户
	Stage         int                                `json:"stage,omitempty"`         // 工单阶段，1：bot，2：人工
	Status        int                                `json:"status,omitempty"`        // 工单状态，1：已创建 2: 处理中 3: 排队中 4：待定 5：待用户响应 50: 被机器人关闭 51: 被人工关闭
	Score         int                                `json:"score,omitempty"`         // 工单评分，1：不满意，2:一般，3:满意
	CreatedAt     int                                `json:"created_at,omitempty"`    // 工单创建时间
	UpdatedAt     int                                `json:"updated_at,omitempty"`    // 工单更新时间，没有值时为-1
	ClosedAt      int                                `json:"closed_at,omitempty"`     // 工单结束时间
	Agents        []*GetTicketRespTicketAgent        `json:"agents,omitempty"`        // 工单客服
	Channel       int                                `json:"channel,omitempty"`       // 工单渠道
	Solve         int                                `json:"solve,omitempty"`         // 工单是否解决 1:没解决 2:已解决
	ClosedBy      *GetTicketRespTicketClosedBy       `json:"closed_by,omitempty"`     // 关单用户ID
	Collaborators []*GetTicketRespTicketCollaborator `json:"collaborators,omitempty"` // 工单协作者
}

type GetTicketRespTicketGuest struct {
	Id        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketRespTicketAgent struct {
	Id        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketRespTicketClosedBy struct {
	Id        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketRespTicketCollaborator struct {
	Id        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}
