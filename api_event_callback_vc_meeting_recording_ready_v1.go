package lark

import (
	"context"
)

// EventVCMeetingRecordingReadyV1
//
// 发生在录制文件上传完毕时。收到该事件后，方可进行录制文件获取、授权等操作
//
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](/ssl:ttdoc/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/events/recording_ready
func (r *EventCallbackAPI) HandlerEventVCMeetingRecordingReadyV1(f eventVCMeetingRecordingReadyV1Handler) {
	r.cli.eventHandler.eventVCMeetingRecordingReadyV1Handler = f
}

type eventVCMeetingRecordingReadyV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventVCMeetingRecordingReadyV1) (string, error)

type EventVCMeetingRecordingReadyV1 struct {
	Meeting  *EventVCMeetingRecordingReadyV1Meeting `json:"meeting,omitempty"`  // 会议数据
	Url      string                                 `json:"url,omitempty"`      // 会议录制链接
	Duration string                                 `json:"duration,omitempty"` // 录制总时长（单位msec）
}

type EventVCMeetingRecordingReadyV1Meeting struct {
	ID        string                                      `json:"id,omitempty"`         // 会议id
	Topic     string                                      `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                      `json:"meeting_no,omitempty"` // 9位会议号
	Owner     *EventVCMeetingRecordingReadyV1MeetingOwner `json:"owner,omitempty"`      // 会议拥有者
}

type EventVCMeetingRecordingReadyV1MeetingOwner struct {
	ID *EventVCMeetingRecordingReadyV1MeetingOwnerID `json:"id,omitempty"` // 用户ID
}

type EventVCMeetingRecordingReadyV1MeetingOwnerID struct {
	UserID  string `json:"user_id,omitempty"`  // 用户的user_id
	OpenID  string `json:"open_id,omitempty"`  // 用户的open_id
	UnionID string `json:"union_id,omitempty"` // 用户的union_id
}
