// Code generated by lark_sdk_gen. DO NOT EDIT.

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
func (r *CalendarService) UnsubscribeCalendar(ctx context.Context, request *UnsubscribeCalendarReq, options ...MethodOptionFunc) (*UnsubscribeCalendarResp, *Response, error) {
	if r.cli.mock.mockCalendarUnsubscribeCalendar != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UnsubscribeCalendar mock enable")
		return r.cli.mock.mockCalendarUnsubscribeCalendar(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Calendar#UnsubscribeCalendar call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UnsubscribeCalendar request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(unsubscribeCalendarResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Calendar#UnsubscribeCalendar POST https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Calendar#UnsubscribeCalendar POST https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Calendar", "UnsubscribeCalendar", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UnsubscribeCalendar success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockCalendarUnsubscribeCalendar(f func(ctx context.Context, request *UnsubscribeCalendarReq, options ...MethodOptionFunc) (*UnsubscribeCalendarResp, *Response, error)) {
	r.mockCalendarUnsubscribeCalendar = f
}

func (r *Mock) UnMockCalendarUnsubscribeCalendar() {
	r.mockCalendarUnsubscribeCalendar = nil
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
