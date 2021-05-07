package lark

import (
	"context"
)

// EventV2VCMeetingShareEndedV1
//
// 发生在屏幕共享停止时
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](/ssl:ttdoc/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/events/share_ended
func (r *EventCallbackAPI) HandlerEventV2VCMeetingShareEndedV1(f eventV2VCMeetingShareEndedV1Handler) {
	r.cli.eventHandler.eventV2VCMeetingShareEndedV1Handler = f
}

type eventV2VCMeetingShareEndedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCMeetingShareEndedV1) (string, error)

type EventV2VCMeetingShareEndedV1 struct {
	Meeting  *EventV2VCMeetingShareEndedV1Meeting  `json:"meeting,omitempty"`  // 会议数据
	Operator *EventV2VCMeetingShareEndedV1Operator `json:"operator,omitempty"` // 事件操作人
}

type EventV2VCMeetingShareEndedV1Meeting struct {
	ID        string                                       `json:"id,omitempty"`         // 会议id
	Topic     string                                       `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                       `json:"meeting_no,omitempty"` // 9位会议号
	StartTime string                                       `json:"start_time,omitempty"` // 会议开始时间（unix时间，单位sec）
	EndTime   string                                       `json:"end_time,omitempty"`   // 会议结束时间（unix时间，单位sec）
	HostUser  *EventV2VCMeetingShareEndedV1MeetingHostUser `json:"host_user,omitempty"`  // 会议主持人
	Owner     *EventV2VCMeetingShareEndedV1MeetingOwner    `json:"owner,omitempty"`      // 会议拥有者
}

type EventV2VCMeetingShareEndedV1MeetingHostUser struct {
	ID       *EventV2VCMeetingShareEndedV1MeetingHostUserID `json:"id,omitempty"`        // 用户ID
	UserRole int                                            `json:"user_role,omitempty"` // 用户会中角色，可用值：【1（普通参会人），2（主持人），3（联席主持人）】
	UserType int                                            `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type EventV2VCMeetingShareEndedV1MeetingHostUserID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}

type EventV2VCMeetingShareEndedV1MeetingOwner struct {
	ID       *EventV2VCMeetingShareEndedV1MeetingOwnerID `json:"id,omitempty"`        // 用户ID
	UserRole int                                         `json:"user_role,omitempty"` // 用户会中角色，可用值：【1（普通参会人），2（主持人），3（联席主持人）】
	UserType int                                         `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type EventV2VCMeetingShareEndedV1MeetingOwnerID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}

type EventV2VCMeetingShareEndedV1Operator struct {
	ID       *EventV2VCMeetingShareEndedV1OperatorID `json:"id,omitempty"`        // 用户ID
	UserRole int                                     `json:"user_role,omitempty"` // 用户会中角色，可用值：【1（普通参会人），2（主持人），3（联席主持人）】
	UserType int                                     `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type EventV2VCMeetingShareEndedV1OperatorID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}
