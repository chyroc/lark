// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteCalendarACL
//
// 该接口用于以当前身份（应用 / 用户）删除日历的控制权限，即日历成员。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/delete
func (r *CalendarService) DeleteCalendarACL(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error) {
	if r.cli.mock.mockCalendarDeleteCalendarACL != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarACL mock enable")
		return r.cli.mock.mockCalendarDeleteCalendarACL(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Calendar#DeleteCalendarACL call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarACL request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(deleteCalendarACLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Calendar#DeleteCalendarACL DELETE https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Calendar#DeleteCalendarACL DELETE https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Calendar", "DeleteCalendarACL", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarACL success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockCalendarDeleteCalendarACL(f func(ctx context.Context, request *DeleteCalendarACLReq, options ...MethodOptionFunc) (*DeleteCalendarACLResp, *Response, error)) {
	r.mockCalendarDeleteCalendarACL = f
}

func (r *Mock) UnMockCalendarDeleteCalendarACL() {
	r.mockCalendarDeleteCalendarACL = nil
}

type DeleteCalendarACLReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	ACLID      string `path:"acl_id" json:"-"`      // acl资源ID, 示例值："user_xxxxxx"
}

type deleteCalendarACLResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarACLResp `json:"data,omitempty"`
}

type DeleteCalendarACLResp struct{}
