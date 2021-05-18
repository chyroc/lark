// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UnsubscribeEvent 用于取消订阅服务台事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/event/unsubscribe
func (r *HelpdeskService) UnsubscribeEvent(ctx context.Context, request *UnsubscribeEventReq, options ...MethodOptionFunc) (*UnsubscribeEventResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUnsubscribeEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UnsubscribeEvent mock enable")
		return r.cli.mock.mockHelpdeskUnsubscribeEvent(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Helpdesk#UnsubscribeEvent call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UnsubscribeEvent request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/events/unsubscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedHelpdeskAuth: true,
	}
	resp := new(unsubscribeEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#UnsubscribeEvent POST https://open.feishu.cn/open-apis/helpdesk/v1/events/unsubscribe failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Helpdesk#UnsubscribeEvent POST https://open.feishu.cn/open-apis/helpdesk/v1/events/unsubscribe failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "UnsubscribeEvent", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UnsubscribeEvent success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskUnsubscribeEvent(f func(ctx context.Context, request *UnsubscribeEventReq, options ...MethodOptionFunc) (*UnsubscribeEventResp, *Response, error)) {
	r.mockHelpdeskUnsubscribeEvent = f
}

func (r *Mock) UnMockHelpdeskUnsubscribeEvent() {
	r.mockHelpdeskUnsubscribeEvent = nil
}

type UnsubscribeEventReq struct {
	Events []*UnsubscribeEventReqEvent `json:"events,omitempty"` // event list to unsubscribe
}

type UnsubscribeEventReqEvent struct {
	Type    string `json:"type,omitempty"`    // 事件类型, 示例值："helpdesk.ticket_message"
	Subtype string `json:"subtype,omitempty"` // 事件子类型, 示例值："ticket_message.created_v1"
}

type unsubscribeEventResp struct {
	Code int                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *UnsubscribeEventResp `json:"data,omitempty"`
}

type UnsubscribeEventResp struct{}
