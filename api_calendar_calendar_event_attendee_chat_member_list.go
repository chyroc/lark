package lark

import (
	"context"
)

// GetCalendarEventAttendeeChatMemberList 获取日程的群参与人的群成员列表。
//
// - 当前身份必须有权限查看日程的参与人列表。
// - 当前身份必须在群聊中，或有权限查看群成员列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee-chat_member/list
func (r *CalendarAPI) GetCalendarEventAttendeeChatMemberList(ctx context.Context, request *GetCalendarEventAttendeeChatMemberListReq) (*GetCalendarEventAttendeeChatMemberListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getCalendarEventAttendeeChatMemberListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "GetCalendarEventAttendeeChatMemberList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetCalendarEventAttendeeChatMemberListReq struct {
	PageToken  *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果,**示例值**："23jhysaxxxxsysy"
	PageSize   *int    `query:"page_size" json:"-"`  // 分页大小,**示例值**：10,**数据校验规则**：,- 最大值：`100`
	CalendarID string  `path:"calendar_id" json:"-"` // 日历 ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID    string  `path:"event_id" json:"-"`    // 日程 ID,**示例值**："xxxxxxxxx_0"
	AttendeeID string  `path:"attendee_id" json:"-"` // 参与人 ID,**示例值**："oc_xxxxxxxx"
}

type getCalendarEventAttendeeChatMemberListResp struct {
	Code int                                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                      `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventAttendeeChatMemberListResp `json:"data,omitempty"` // \-
}

type GetCalendarEventAttendeeChatMemberListResp struct {
	Items     []*GetCalendarEventAttendeeChatMemberListRespItem `json:"items,omitempty"`      // 群中的群成员，当type为chat时有效；群成员不支持编辑
	HasMore   bool                                              `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                            `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
}

type GetCalendarEventAttendeeChatMemberListRespItem struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`  // 参与人RSVP状态,**可选值有**：,- `needs_action`：参与人尚未回复状态，或表示会议室预约中,- `accept`：参与人回复接受，或表示会议室预约成功,- `tentative`：参与人回复待定,- `decline`：参与人回复拒绝，或表示会议室预约失败,- `removed`：参与人或会议室已经从日程中被移除
	IsOptional  bool   `json:"is_optional,omitempty"`  // 参与人是否为「可选参加」,**默认值**：`false`
	DisplayName string `json:"display_name,omitempty"` // 参与人名称
	IsOrganizer bool   `json:"is_organizer,omitempty"` // 参与人是否为日程组织者
	IsExternal  bool   `json:"is_external,omitempty"`  // 参与人是否为外部参与人
}
