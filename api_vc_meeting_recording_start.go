// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// StartVCMeetingRecording 在会议中开始录制。
//
// 会议正在进行中, 且操作者具有相应权限（如果操作者为用户, 必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/start
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/meeting-recording/start
func (r *VCService) StartVCMeetingRecording(ctx context.Context, request *StartVCMeetingRecordingReq, options ...MethodOptionFunc) (*StartVCMeetingRecordingResp, *Response, error) {
	if r.cli.mock.mockVCStartVCMeetingRecording != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] VC#StartVCMeetingRecording mock enable")
		return r.cli.mock.mockVCStartVCMeetingRecording(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "StartVCMeetingRecording",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/vc/v1/meetings/:meeting_id/recording/start",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(startVCMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCStartVCMeetingRecording mock VCStartVCMeetingRecording method
func (r *Mock) MockVCStartVCMeetingRecording(f func(ctx context.Context, request *StartVCMeetingRecordingReq, options ...MethodOptionFunc) (*StartVCMeetingRecordingResp, *Response, error)) {
	r.mockVCStartVCMeetingRecording = f
}

// UnMockVCStartVCMeetingRecording un-mock VCStartVCMeetingRecording method
func (r *Mock) UnMockVCStartVCMeetingRecording() {
	r.mockVCStartVCMeetingRecording = nil
}

// StartVCMeetingRecordingReq ...
type StartVCMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID（视频会议的唯一标识, 视频会议开始后才会产生）, 示例值: "6911188411932033028"
	Timezone  *int64 `json:"timezone,omitempty"`  // 录制文件时间显示使用的时区[-12, 12], 示例值: 8
}

// StartVCMeetingRecordingResp ...
type StartVCMeetingRecordingResp struct {
}

// startVCMeetingRecordingResp ...
type startVCMeetingRecordingResp struct {
	Code  int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                       `json:"msg,omitempty"`  // 错误描述
	Data  *StartVCMeetingRecordingResp `json:"data,omitempty"`
	Error *ErrorDetail                 `json:"error,omitempty"`
}
