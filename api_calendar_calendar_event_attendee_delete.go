package lark

import (
	"context"
)

// DeleteCalendarEventAttendee 批量删除日程的参与人。
//
// - 当前身份需要有日历的 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
// - 当前身份需要是日程的组织者。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/batch_delete
func (r *CalendarAPI) DeleteCalendarEventAttendee(ctx context.Context, request *DeleteCalendarEventAttendeeReq) (*DeleteCalendarEventAttendeeResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteCalendarEventAttendeeResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "DeleteCalendarEventAttendee", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteCalendarEventAttendeeReq struct {
	CalendarID  string   `path:"calendar_id" json:"-"`   // 日历 ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID     string   `path:"event_id" json:"-"`      // 日程 ID,**示例值**："xxxxxxxxx_0"
	AttendeeIDs []string `json:"attendee_ids,omitempty"` // 要移除的参与人 ID 列表
}

type deleteCalendarEventAttendeeResp struct {
	Code int                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarEventAttendeeResp `json:"data,omitempty"`
}

type DeleteCalendarEventAttendeeResp struct{}
