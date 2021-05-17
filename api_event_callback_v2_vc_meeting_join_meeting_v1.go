// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2VCMeetingJoinMeetingV1
//
// 发生在有人加入会议时
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/join_meeting
func (r *EventCallbackService) HandlerEventV2VCMeetingJoinMeetingV1(f eventV2VCMeetingJoinMeetingV1Handler) {
	r.cli.eventHandler.eventV2VCMeetingJoinMeetingV1Handler = f
}

type eventV2VCMeetingJoinMeetingV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCMeetingJoinMeetingV1) (string, error)

type EventV2VCMeetingJoinMeetingV1 struct {
	Meeting  *EventV2VCMeetingJoinMeetingV1Meeting  `json:"meeting,omitempty"`  // 会议数据
	Operator *EventV2VCMeetingJoinMeetingV1Operator `json:"operator,omitempty"` // 事件操作人
}

type EventV2VCMeetingJoinMeetingV1Meeting struct {
	ID        string                                        `json:"id,omitempty"`         // 会议ID
	Topic     string                                        `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                        `json:"meeting_no,omitempty"` // 9位会议号
	StartTime string                                        `json:"start_time,omitempty"` // 会议开始时间（unix时间，单位sec）
	EndTime   string                                        `json:"end_time,omitempty"`   // 会议结束时间（unix时间，单位sec）
	HostUser  *EventV2VCMeetingJoinMeetingV1MeetingHostUser `json:"host_user,omitempty"`  // 会议主持人
	Owner     *EventV2VCMeetingJoinMeetingV1MeetingOwner    `json:"owner,omitempty"`      // 会议拥有者
}

type EventV2VCMeetingJoinMeetingV1MeetingHostUser struct {
	ID       *EventV2VCMeetingJoinMeetingV1MeetingHostUserID `json:"id,omitempty"`        // 用户 ID
	UserRole int                                             `json:"user_role,omitempty"` // 用户会中角色, 可选值有: `1`：普通参会人, `2`：主持人, `3`：联席主持人
	UserType int                                             `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type EventV2VCMeetingJoinMeetingV1MeetingHostUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2VCMeetingJoinMeetingV1MeetingOwner struct {
	ID       *EventV2VCMeetingJoinMeetingV1MeetingOwnerID `json:"id,omitempty"`        // 用户 ID
	UserRole int                                          `json:"user_role,omitempty"` // 用户会中角色, 可选值有: `1`：普通参会人, `2`：主持人, `3`：联席主持人
	UserType int                                          `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type EventV2VCMeetingJoinMeetingV1MeetingOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2VCMeetingJoinMeetingV1Operator struct {
	ID       *EventV2VCMeetingJoinMeetingV1OperatorID `json:"id,omitempty"`        // 用户 ID
	UserRole int                                      `json:"user_role,omitempty"` // 用户会中角色, 可选值有: `1`：普通参会人, `2`：主持人, `3`：联席主持人
	UserType int                                      `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type EventV2VCMeetingJoinMeetingV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
