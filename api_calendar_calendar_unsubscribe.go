package lark

import (
	"context"
)

// UnsubscribeCalendar
//
// 该接口用于以当前身份（应用 / 用户）取消对某日历的订阅状态。
// 身份由 Header Authorization 的 Token 类型决定。
// 仅可操作已经被当前身份订阅的日历。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/unsubscribe
func (r *CalendarAPI) UnsubscribeCalendar(ctx context.Context, request *UnsubscribeCalendarReq) (*UnsubscribeCalendarResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(unsubscribeCalendarResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Calendar", "UnsubscribeCalendar", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UnsubscribeCalendarReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type unsubscribeCalendarResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UnsubscribeCalendarResp `json:"data,omitempty"`
}

type UnsubscribeCalendarResp struct{}
