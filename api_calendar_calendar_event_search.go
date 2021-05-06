package lark

import (
	"context"
)

// SearchCalendarEvent
//
// 该接口用于以用户身份搜索某日历下的相关日程。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份必须对日历有访问权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search
func (r *CalendarAPI) SearchCalendarEvent(ctx context.Context, request *SearchCalendarEventReq) (*SearchCalendarEventResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/search",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(searchCalendarEventResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "SearchCalendarEvent", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SearchCalendarEventReq struct {
	UserIDType *IDType                       `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	PageToken  *string                       `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："xxxxx"
	PageSize   *int                          `query:"page_size" json:"-"`    // 分页大小, 示例值：10, 最大值：`xxxxx`
	CalendarID string                        `path:"calendar_id" json:"-"`   // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	Query      string                        `json:"query,omitempty"`        // 搜索关键字, 示例值："query words", 长度范围：`1` ～ `200` 字符
	Filter     *SearchCalendarEventReqFilter `json:"filter,omitempty"`       // 搜索过滤器
}

type SearchCalendarEventReqFilter struct {
	StartTime *SearchCalendarEventReqFilterStartTime `json:"start_time,omitempty"` // 搜索过滤项，日程搜索区间的开始时间，被搜索日程的事件必须与搜索区间有交集
	EndTime   *SearchCalendarEventReqFilterEndTime   `json:"end_time,omitempty"`   // 搜索过滤项，日程搜索区间的结束时间，被搜索日程的事件必须与搜索区间有交集
	UserIDs   []string                               `json:"user_ids,omitempty"`   // 搜索过滤项，参与人的用户ID列表，被搜索日程中必须包含至少一个其中的参与人
	RoomIDs   []string                               `json:"room_ids,omitempty"`   // 搜索过滤项，会议室ID列表，被搜索日程中必须包含至少一个其中的会议室
	ChatIDs   []string                               `json:"chat_ids,omitempty"`   // 搜索过滤项，群ID列表，被搜索日程的参与人中必须包含至少一个其中的群
}

type SearchCalendarEventReqFilterStartTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定, 示例值：" 2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区), 示例值："1605024000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai, 示例值："Asia/Shanghai"
}

type SearchCalendarEventReqFilterEndTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定, 示例值：" 2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区), 示例值："1605024000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai, 示例值："Asia/Shanghai"
}

type searchCalendarEventResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *SearchCalendarEventResp `json:"data,omitempty"` //
}

type SearchCalendarEventResp struct {
	Items     []*SearchCalendarEventRespItem `json:"items,omitempty"`      // 搜索命中的日程列表
	PageToken string                         `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
}

type SearchCalendarEventRespItem struct {
	EventID          string                                 `json:"event_id,omitempty"`           // 日程ID
	Summary          string                                 `json:"summary,omitempty"`            // 日程标题, 最大长度：`1000` 字符
	Description      string                                 `json:"description,omitempty"`        // 日程描述, 最大长度：`8192` 字符
	StartTime        *SearchCalendarEventRespItemStartTime  `json:"start_time,omitempty"`         // 日程开始时间
	EndTime          *SearchCalendarEventRespItemEndTime    `json:"end_time,omitempty"`           // 日程结束时间
	Vchat            *SearchCalendarEventRespItemVchat      `json:"vchat,omitempty"`              // 视频会议信息，仅当日程至少有一位attendee时生效
	Visibility       string                                 `json:"visibility,omitempty"`         // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `default`：默认权限，仅向他人显示是否“忙碌”, `public`：公开，显示日程详情, `private`：私密，仅自己可见
	AttendeeAbility  string                                 `json:"attendee_ability,omitempty"`   // 参与人权限, 可选值有: `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表, `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表, `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表, `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   string                                 `json:"free_busy_status,omitempty"`   // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效, 可选值有: `busy`：忙碌, `free`：空闲
	Location         *SearchCalendarEventRespItemLocation   `json:"location,omitempty"`           // 日程地点
	Color            int                                    `json:"color,omitempty"`              // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders        []*SearchCalendarEventRespItemReminder `json:"reminders,omitempty"`          // 日程提醒列表
	Recurrence       string                                 `json:"recurrence,omitempty"`         // 重复日程的重复性规则, 最大长度：`2000` 字符
	Status           string                                 `json:"status,omitempty"`             // 日程状态, 可选值有: `tentative`：未回应, `confirmed`：已确认, `cancelled`：日程已取消
	IsException      bool                                   `json:"is_exception,omitempty"`       // 日程是否是一个重复日程的例外日程
	RecurringEventID string                                 `json:"recurring_event_id,omitempty"` // 例外日程的原重复日程的event_id
	Schemas          []*SearchCalendarEventRespItemSchema   `json:"schemas,omitempty"`            // 日程自定义信息
}

type SearchCalendarEventRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type SearchCalendarEventRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type SearchCalendarEventRespItemVchat struct {
	MeetingUrl string `json:"meeting_url,omitempty"` // 视频会议URL, 长度范围：`1` ～ `2000` 字符
}

type SearchCalendarEventRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称, 长度范围：`1` ～ `512` 字符
	Address   string  `json:"address,omitempty"`   // 地点地址, 长度范围：`1` ～ `255` 字符
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
}

type SearchCalendarEventRespItemReminder struct {
	Minutes int `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效, 取值范围：`-20160` ～ `20160`
}

type SearchCalendarEventRespItemSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI项名称 TODO文档
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态, 可选值有: `hide`：隐藏显示, `readonly`：只读, `editable`：可编辑, `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接, 最大长度：`2000` 字符
}
