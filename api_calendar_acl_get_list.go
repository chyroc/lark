package lark

import (
	"context"
)

// GetCalendarACLList
//
// 该接口用于以当前身份（应用 / 用户）获取日历的控制权限列表。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/list
func (r *CalendarAPI) GetCalendarACLList(ctx context.Context, request *GetCalendarACLListReq) (*GetCalendarACLListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getCalendarACLListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "GetCalendarACLList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetCalendarACLListReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果,**示例值**："xxx"
	PageSize   *int    `query:"page_size" json:"-"`    // 分页大小,**示例值**：10，小于10取10,**数据校验规则**：,- 最大值：`50`
	CalendarID string  `path:"calendar_id" json:"-"`   // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type getCalendarACLListResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarACLListResp `json:"data,omitempty"` // \-
}

type GetCalendarACLListResp struct {
	Acls      []*GetCalendarACLListRespAcl `json:"acls,omitempty"`       // 入参日历对应的acl列表
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                       `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
}

type GetCalendarACLListRespAcl struct {
	ACLID string                          `json:"acl_id,omitempty"` // acl资源ID
	Role  CalendarRole                    `json:"role,omitempty"`   // 对日历的访问权限,**可选值有**：,- `unknown`：未知权限,- `free_busy_reader`：游客，只能看到忙碌/空闲信息,- `reader`：订阅者，查看所有日程详情,- `writer`：编辑者，创建及修改日程,- `owner`：管理员，管理日历及共享设置
	Scope *GetCalendarACLListRespAclScope `json:"scope,omitempty"`  // 权限范围
}

type GetCalendarACLListRespAclScope struct {
	Type   string `json:"type,omitempty"`    // 权限类型，当type为User时，值为open_id/user_id/union_id,**可选值有**：,- `user`：用户
	UserID string `json:"user_id,omitempty"` // 用户ID
}
