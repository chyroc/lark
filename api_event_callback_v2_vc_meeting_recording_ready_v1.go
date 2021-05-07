package lark

import (
	"context"
)

// EventV2VCMeetingRecordingReadyV1
//
// 发生在录制文件上传完毕时。收到该事件后，方可进行录制文件获取、授权等操作
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](/ssl:ttdoc/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/events/recording_ready
func (r *EventCallbackAPI) HandlerEventV2VCMeetingRecordingReadyV1(f eventV2VCMeetingRecordingReadyV1Handler) {
	r.cli.eventHandler.eventV2VCMeetingRecordingReadyV1Handler = f
}

type eventV2VCMeetingRecordingReadyV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCMeetingRecordingReadyV1) (string, error)

type EventV2VCMeetingRecordingReadyV1 struct {
	Meeting  *EventV2VCMeetingRecordingReadyV1Meeting `json:"meeting,omitempty"`  // 会议数据
	Url      string                                   `json:"url,omitempty"`      // 会议录制链接
	Duration string                                   `json:"duration,omitempty"` // 录制总时长（单位msec）
}

type EventV2VCMeetingRecordingReadyV1Meeting struct {
	ID        string                                        `json:"id,omitempty"`         // 会议id
	Topic     string                                        `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                        `json:"meeting_no,omitempty"` // 9位会议号
	Owner     *EventV2VCMeetingRecordingReadyV1MeetingOwner `json:"owner,omitempty"`      // 会议拥有者
}

type EventV2VCMeetingRecordingReadyV1MeetingOwner struct {
	ID *EventV2VCMeetingRecordingReadyV1MeetingOwnerID `json:"id,omitempty"` // 用户ID
}

type EventV2VCMeetingRecordingReadyV1MeetingOwnerID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}
