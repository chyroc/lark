package lark

import (
	"context"
)

// EventVCMeetingMeetingStartedV1
//
// 发生在会议开始时
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](/ssl:ttdoc/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/events/meeting_started
func (r *EventCallbackAPI) HandlerEventVCMeetingMeetingStartedV1(f eventVCMeetingMeetingStartedV1Handler) {
	r.cli.eventHandler.eventVCMeetingMeetingStartedV1Handler = f
}

type eventVCMeetingMeetingStartedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventVCMeetingMeetingStartedV1) (string, error)

type EventVCMeetingMeetingStartedV1 struct {
	Meeting  *EventVCMeetingMeetingStartedV1Meeting  `json:"meeting,omitempty"`  // 会议数据
	Operator *EventVCMeetingMeetingStartedV1Operator `json:"operator,omitempty"` // 事件操作人
}

type EventVCMeetingMeetingStartedV1Meeting struct {
	ID        string                                         `json:"id,omitempty"`         // 会议id
	Topic     string                                         `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                         `json:"meeting_no,omitempty"` // 9位会议号
	StartTime string                                         `json:"start_time,omitempty"` // 会议开始时间（unix时间，单位sec）
	EndTime   string                                         `json:"end_time,omitempty"`   // 会议结束时间（unix时间，单位sec）
	HostUser  *EventVCMeetingMeetingStartedV1MeetingHostUser `json:"host_user,omitempty"`  // 会议主持人
	Owner     *EventVCMeetingMeetingStartedV1MeetingOwner    `json:"owner,omitempty"`      // 会议拥有者
}

type EventVCMeetingMeetingStartedV1MeetingHostUser struct {
	ID       *EventVCMeetingMeetingStartedV1MeetingHostUserID `json:"id,omitempty"`        // 用户ID
	UserRole int                                              `json:"user_role,omitempty"` // 用户会中角色，可用值：【1（普通参会人），2（主持人），3（联席主持人）】
	UserType int                                              `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type EventVCMeetingMeetingStartedV1MeetingHostUserID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}

type EventVCMeetingMeetingStartedV1MeetingOwner struct {
	ID       *EventVCMeetingMeetingStartedV1MeetingOwnerID `json:"id,omitempty"`        // 用户ID
	UserRole int                                           `json:"user_role,omitempty"` // 用户会中角色，可用值：【1（普通参会人），2（主持人），3（联席主持人）】
	UserType int                                           `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type EventVCMeetingMeetingStartedV1MeetingOwnerID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}

type EventVCMeetingMeetingStartedV1Operator struct {
	ID       *EventVCMeetingMeetingStartedV1OperatorID `json:"id,omitempty"`        // 用户ID
	UserRole int                                       `json:"user_role,omitempty"` // 用户会中角色，可用值：【1（普通参会人），2（主持人），3（联席主持人）】
	UserType int                                       `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type EventVCMeetingMeetingStartedV1OperatorID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}
