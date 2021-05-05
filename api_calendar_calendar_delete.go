package lark

import (
	"context"
)

// DeleteCalendar
//
// 该接口用于以当前身份（应用 / 用户）删除一个共享日历。
//
// 身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar&method=delete)
//
//
//
//
//
//
//
// 当前身份必须对日历具有 owner 权限。
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/delete
func (r *CalendarAPI) DeleteCalendar(ctx context.Context, request *DeleteCalendarReq) (*DeleteCalendarResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteCalendarResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "DeleteCalendar", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteCalendarReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type deleteCalendarResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarResp `json:"data,omitempty"`
}

type DeleteCalendarResp struct{}
