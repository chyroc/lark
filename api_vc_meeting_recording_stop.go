package lark

import (
	"context"
)

// StopMeetingRecording
//
// > 在会议中停止录制
// 会议正在录制中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting.recording/stop
func (r *VCAPI) StopMeetingRecording(ctx context.Context, request *StopMeetingRecordingReq) (*StopMeetingRecordingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/stop",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(stopMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "StopMeetingRecording", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type StopMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议id，示例值："6911188411932033028"
}

type stopMeetingRecordingResp struct {
	Code int                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *StopMeetingRecordingResp `json:"data,omitempty"`
}

type StopMeetingRecordingResp struct{}
