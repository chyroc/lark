package lark

import (
	"context"
)

// DeleteCalendarEvent
//
// 该接口用于以当前身份（应用 / 用户）删除日历上的一个日程。
// 身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=delete)
// 当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
// 当前身份必须是日程的组织者。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/delete
func (r *CalendarAPI) DeleteCalendarEvent(ctx context.Context, request *DeleteCalendarEventReq) (*DeleteCalendarEventResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteCalendarEventResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "DeleteCalendarEvent", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteCalendarEventReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID    string `path:"event_id" json:"-"`    // 日程ID,**示例值**："xxxxxxxxx_0"
}

type deleteCalendarEventResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarEventResp `json:"data,omitempty"`
}

type DeleteCalendarEventResp struct{}
