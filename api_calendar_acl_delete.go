package lark

import (
	"context"
)

// DeleteCalendarACL
//
// 该接口用于以当前身份（应用 / 用户）删除日历的控制权限，即日历成员。
// 身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.acl&method=delete)
// 当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/delete
func (r *CalendarAPI) DeleteCalendarACL(ctx context.Context, request *DeleteCalendarACLReq) (*DeleteCalendarACLResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteCalendarACLResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "DeleteCalendarACL", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteCalendarACLReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	ACLID      string `path:"acl_id" json:"-"`      // acl资源ID,**示例值**："user_xxxxxx"
}

type deleteCalendarACLResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarACLResp `json:"data,omitempty"`
}

type DeleteCalendarACLResp struct{}
