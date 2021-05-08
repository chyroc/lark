package lark

import (
	"context"
)

// DeleteCalendarTimeoffEvent 删除一个指定的请假日程，请假日程删除，用户个人签名页的请假信息也会消失。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/delete
func (r *CalendarAPI) DeleteCalendarTimeoffEvent(ctx context.Context, request *DeleteCalendarTimeoffEventReq) (*DeleteCalendarTimeoffEventResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/timeoff_events/:timeoff_event_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteCalendarTimeoffEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Calendar", "DeleteCalendarTimeoffEvent", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteCalendarTimeoffEventReq struct {
	TimeoffEventID string `path:"timeoff_event_id" json:"-"` // 休假申请的唯一标识id, 示例值："timeoff:XXXXXX-XXXX-0917-1623-aa493d591a39"
}

type deleteCalendarTimeoffEventResp struct {
	Code int                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarTimeoffEventResp `json:"data,omitempty"`
}

type DeleteCalendarTimeoffEventResp struct{}
