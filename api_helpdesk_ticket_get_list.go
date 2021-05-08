package lark

import (
	"context"
)

// GetTicketList 该接口用于获取全部工单详情。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list
func (r *HelpdeskAPI) GetTicketList(ctx context.Context, request *GetTicketListReq) (*GetTicketListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      true,
		IsFile:                false,
	}
	resp := new(getTicketListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Helpdesk", "GetTicketList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetTicketListReq struct {
	TicketID         *string                    `query:"ticket_id" json:"-"`         // 搜索条件：工单ID, 示例值："123456"
	AgentID          *string                    `query:"agent_id" json:"-"`          // 搜索条件: 客服id, 示例值："ou_b5de90429xxx"
	ClosedByID       *string                    `query:"closed_by_id" json:"-"`      // 搜索条件: 关单客服id, 示例值："ou_b5de90429xxx"
	Type             *int                       `query:"type" json:"-"`              // 搜索条件: 工单类型 1:bot 2:人工, 示例值：1
	Channel          *int                       `query:"channel" json:"-"`           // 搜索条件: 工单渠道, 示例值：0
	Solved           *int                       `query:"solved" json:"-"`            // 搜索条件: 工单是否解决 1:没解决 2:已解决	, 示例值：1
	Score            *int                       `query:"score" json:"-"`             // 搜索条件: 工单评分, 示例值：1
	StatusList       []int                      `query:"status_list" json:"-"`       // 搜索条件: 工单状态列表
	GuestName        *string                    `query:"guest_name" json:"-"`        // 搜索条件: 用户名称, 示例值："abc"
	GuestID          *string                    `query:"guest_id" json:"-"`          // 搜索条件: 用户id, 示例值："ou_b5de90429xxx"
	CustomizedFields []*HelpdeskCustomizedField `query:"customized_fields" json:"-"` // 搜索条件: 自定义字段列表
	Tags             []string                   `query:"tags" json:"-"`              // 搜索条件: 用户标签列表
	Page             *int                       `query:"page" json:"-"`              // 页数, 从1开始, 默认为1, 示例值：1
	PageSize         *int                       `query:"page_size" json:"-"`         // 当前页大小，最大为200, 默认为20, 示例值：20
	CreateTimeStart  *int                       `query:"create_time_start" json:"-"` // 搜索条件: 工单创建起始时间 ms, 示例值：1616920429000
	CreateTimeEnd    *int                       `query:"create_time_end" json:"-"`   // 搜索条件: 工单创建结束时间 ms, 示例值：1616920429000
	UpdateTimeStart  *int                       `query:"update_time_start" json:"-"` // 搜索条件: 工单修改起始时间 ms, 示例值：1616920429000
	UpdateTimeEnd    *int                       `query:"update_time_end" json:"-"`   // 搜索条件: 工单修改结束时间 ms, 示例值：1616920429000
}

type getTicketListResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *GetTicketListResp `json:"data,omitempty"` //
}

type GetTicketListResp struct {
	Total   int                        `json:"total,omitempty"`   // 工单总数
	Tickets []*GetTicketListRespTicket `json:"tickets,omitempty"` // 工单
}

type GetTicketListRespTicket struct {
	TicketID      string                                 `json:"ticket_id,omitempty"`     // 工单ID
	HelpdeskID    string                                 `json:"helpdesk_id,omitempty"`   // 服务台ID
	Guest         *GetTicketListRespTicketGuest          `json:"guest,omitempty"`         // 工单创建用户
	Stage         int                                    `json:"stage,omitempty"`         // 工单阶段，1：bot，2：人工
	Status        int                                    `json:"status,omitempty"`        // 工单状态，1：已创建 2: 处理中 3: 排队中 4：待定 5：待用户响应 50: 被机器人关闭 51: 被人工关闭
	Score         int                                    `json:"score,omitempty"`         // 工单评分，1：不满意，2:一般，3:满意
	CreatedAt     int                                    `json:"created_at,omitempty"`    // 工单创建时间
	UpdatedAt     int                                    `json:"updated_at,omitempty"`    // 工单更新时间，没有值时为-1
	ClosedAt      int                                    `json:"closed_at,omitempty"`     // 工单结束时间
	Agents        []*GetTicketListRespTicketAgent        `json:"agents,omitempty"`        // 工单客服
	Channel       int                                    `json:"channel,omitempty"`       // 工单渠道
	Solve         int                                    `json:"solve,omitempty"`         // 工单是否解决 1:没解决 2:已解决
	ClosedBy      *GetTicketListRespTicketClosedBy       `json:"closed_by,omitempty"`     // 关单用户ID
	Collaborators []*GetTicketListRespTicketCollaborator `json:"collaborators,omitempty"` // 工单协作者
}

type GetTicketListRespTicketGuest struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketAgent struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketClosedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketListRespTicketCollaborator struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}
