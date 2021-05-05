package lark

import (
	"context"
)

// UpdateCalendarEvent
//
// 该接口用于以当前身份（应用 / 用户）更新日历上的一个日程。
// 身份由 Header Authorization 的 Token 类型决定。{尝试一下}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&method=patch)
// 当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
// 当前身份为日程组织者时，可修改所有可编辑字段。
// 当前身份为日程参与者时，仅可编辑部分字段。（如：visibility, free_busy_status, color, reminders）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/patch
func (r *CalendarAPI) UpdateCalendarEvent(ctx context.Context, request *UpdateCalendarEventReq) (*UpdateCalendarEventResp, *Response, error) {
	req := &requestParam{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateCalendarEventResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "UpdateCalendarEvent", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateCalendarEventReq struct {
	CalendarID      string                            `path:"calendar_id" json:"-"`       // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID         string                            `path:"event_id" json:"-"`          // 日程ID,**示例值**："xxxxxxxxx_0"
	Summary         *string                           `json:"summary,omitempty"`          // 日程标题,**示例值**："日程标题",**数据校验规则**：,- 最大长度：`1000` 字符
	Description     *string                           `json:"description,omitempty"`      // 日程描述,**示例值**："日程描述",**数据校验规则**：,- 最大长度：`8192` 字符
	StartTime       *UpdateCalendarEventReqStartTime  `json:"start_time,omitempty"`       // 日程开始时间
	EndTime         *UpdateCalendarEventReqEndTime    `json:"end_time,omitempty"`         // 日程结束时间
	Vchat           *UpdateCalendarEventReqVchat      `json:"vchat,omitempty"`            // 视频会议信息，仅当日程至少有一位attendee时生效
	Visibility      *string                           `json:"visibility,omitempty"`       // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效,**示例值**："default",**可选值有**：,- `default`：默认权限，仅向他人显示是否“忙碌”,- `public`：公开，显示日程详情,- `private`：私密，仅自己可见
	AttendeeAbility *string                           `json:"attendee_ability,omitempty"` // 参与人权限,**示例值**："can_see_others",**可选值有**：,- `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表,- `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表,- `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表,- `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus  *string                           `json:"free_busy_status,omitempty"` // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效,**示例值**："busy",**可选值有**：,- `busy`：忙碌,- `free`：空闲
	Location        *UpdateCalendarEventReqLocation   `json:"location,omitempty"`         // 日程地点
	Color           *int                              `json:"color,omitempty"`            // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。,**示例值**：-1
	Reminders       []*UpdateCalendarEventReqReminder `json:"reminders,omitempty"`        // 日程提醒列表
	Recurrence      *string                           `json:"recurrence,omitempty"`       // 重复日程的重复性规则,**示例值**："xxxxx",**数据校验规则**：,- 最大长度：`2000` 字符
	Schemas         []*UpdateCalendarEventReqSchema   `json:"schemas,omitempty"`          // 日程自定义信息
}

type UpdateCalendarEventReqStartTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定,**示例值**：" 2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区),**示例值**："1605024000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai,**示例值**："Asia/Shanghai"
}

type UpdateCalendarEventReqEndTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定,**示例值**：" 2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区),**示例值**："1605024000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai,**示例值**："Asia/Shanghai"
}

type UpdateCalendarEventReqVchat struct {
	MeetingUrl *string `json:"meeting_url,omitempty"` // 视频会议URL,**示例值**："https://example.com",**数据校验规则**：,- 长度范围：`1` ～ `2000` 字符
}

type UpdateCalendarEventReqLocation struct {
	Name      *string  `json:"name,omitempty"`      // 地点名称,**示例值**："地点名称",**数据校验规则**：,- 长度范围：`1` ～ `512` 字符
	Address   *string  `json:"address,omitempty"`   // 地点地址,**示例值**："地点地址",**数据校验规则**：,- 长度范围：`1` ～ `255` 字符
	Latitude  *float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准,**示例值**：xxxxx
	Longitude *float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准,**示例值**：xxxxx
}

type UpdateCalendarEventReqReminder struct {
	Minutes *int `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效,**示例值**：5,**数据校验规则**：,- 取值范围：`-20160` ～ `20160`
}

type UpdateCalendarEventReqSchema struct {
	UiName   *string `json:"ui_name,omitempty"`   // UI项名称 TODO文档,**示例值**："ForwardIcon"
	UiStatus *string `json:"ui_status,omitempty"` // UI项自定义状态,**示例值**："hide",**可选值有**：,- `hide`：隐藏显示,- `readonly`：只读,- `editable`：可编辑,- `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  *string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接,**示例值**："xxxxx",**数据校验规则**：,- 最大长度：`2000` 字符
}

type updateCalendarEventResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCalendarEventResp `json:"data,omitempty"` // \-
}

type UpdateCalendarEventResp struct {
	Event *UpdateCalendarEventRespEvent `json:"event,omitempty"` // 更新后的日程实体
}

type UpdateCalendarEventRespEvent struct {
	EventID          string                                  `json:"event_id,omitempty"`           // 日程ID
	Summary          string                                  `json:"summary,omitempty"`            // 日程标题,**数据校验规则**：,- 最大长度：`1000` 字符
	Description      string                                  `json:"description,omitempty"`        // 日程描述,**数据校验规则**：,- 最大长度：`8192` 字符
	StartTime        *UpdateCalendarEventRespEventStartTime  `json:"start_time,omitempty"`         // 日程开始时间
	EndTime          *UpdateCalendarEventRespEventEndTime    `json:"end_time,omitempty"`           // 日程结束时间
	Vchat            *UpdateCalendarEventRespEventVchat      `json:"vchat,omitempty"`              // 视频会议信息，仅当日程至少有一位attendee时生效
	Visibility       string                                  `json:"visibility,omitempty"`         // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效,**可选值有**：,- `default`：默认权限，仅向他人显示是否“忙碌”,- `public`：公开，显示日程详情,- `private`：私密，仅自己可见
	AttendeeAbility  string                                  `json:"attendee_ability,omitempty"`   // 参与人权限,**可选值有**：,- `none`：无法编辑日程、无法邀请其它参与人、无法查看参与人列表,- `can_see_others`：无法编辑日程、无法邀请其它参与人、可以查看参与人列表,- `can_invite_others`：无法编辑日程、可以邀请其它参与人、可以查看参与人列表,- `can_modify_event`：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   string                                  `json:"free_busy_status,omitempty"`   // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效,**可选值有**：,- `busy`：忙碌,- `free`：空闲
	Location         *UpdateCalendarEventRespEventLocation   `json:"location,omitempty"`           // 日程地点
	Color            int                                     `json:"color,omitempty"`              // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders        []*UpdateCalendarEventRespEventReminder `json:"reminders,omitempty"`          // 日程提醒列表
	Recurrence       string                                  `json:"recurrence,omitempty"`         // 重复日程的重复性规则,**数据校验规则**：,- 最大长度：`2000` 字符
	Status           string                                  `json:"status,omitempty"`             // 日程状态,**可选值有**：,- `tentative`：未回应,- `confirmed`：已确认,- `cancelled`：日程已取消
	IsException      bool                                    `json:"is_exception,omitempty"`       // 日程是否是一个重复日程的例外日程
	RecurringEventID string                                  `json:"recurring_event_id,omitempty"` // 例外日程的原重复日程的event_id
	Schemas          []*UpdateCalendarEventRespEventSchema   `json:"schemas,omitempty"`            // 日程自定义信息
}

type UpdateCalendarEventRespEventStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type UpdateCalendarEventRespEventEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

type UpdateCalendarEventRespEventVchat struct {
	MeetingUrl string `json:"meeting_url,omitempty"` // 视频会议URL,**数据校验规则**：,- 长度范围：`1` ～ `2000` 字符
}

type UpdateCalendarEventRespEventLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称,**数据校验规则**：,- 长度范围：`1` ～ `512` 字符
	Address   string  `json:"address,omitempty"`   // 地点地址,**数据校验规则**：,- 长度范围：`1` ～ `255` 字符
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
}

type UpdateCalendarEventRespEventReminder struct {
	Minutes int `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒,新建或更新日程时传入该字段，仅对当前身份生效,**数据校验规则**：,- 取值范围：`-20160` ～ `20160`
}

type UpdateCalendarEventRespEventSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI项名称 TODO文档
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态,**可选值有**：,- `hide`：隐藏显示,- `readonly`：只读,- `editable`：可编辑,- `unknown`：未知UI项自定义状态，仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接,**数据校验规则**：,- 最大长度：`2000` 字符
}
