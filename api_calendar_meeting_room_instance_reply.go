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

// ReplyCalendarMeetingRoomInstance 该接口用于回复会议室日程实例, 包括未签到释放和提前结束释放。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYzN4UjL2cDO14iN3gTN
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/meeting-room-event/reply-meeting-room-event-instance
func (r *CalendarService) ReplyCalendarMeetingRoomInstance(ctx context.Context, request *ReplyCalendarMeetingRoomInstanceReq, options ...MethodOptionFunc) (*ReplyCalendarMeetingRoomInstanceResp, *Response, error) {
	if r.cli.mock.mockCalendarReplyCalendarMeetingRoomInstance != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#ReplyCalendarMeetingRoomInstance mock enable")
		return r.cli.mock.mockCalendarReplyCalendarMeetingRoomInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "ReplyCalendarMeetingRoomInstance",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/meeting_room/instance/reply",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(replyCalendarMeetingRoomInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarReplyCalendarMeetingRoomInstance mock CalendarReplyCalendarMeetingRoomInstance method
func (r *Mock) MockCalendarReplyCalendarMeetingRoomInstance(f func(ctx context.Context, request *ReplyCalendarMeetingRoomInstanceReq, options ...MethodOptionFunc) (*ReplyCalendarMeetingRoomInstanceResp, *Response, error)) {
	r.mockCalendarReplyCalendarMeetingRoomInstance = f
}

// UnMockCalendarReplyCalendarMeetingRoomInstance un-mock CalendarReplyCalendarMeetingRoomInstance method
func (r *Mock) UnMockCalendarReplyCalendarMeetingRoomInstance() {
	r.mockCalendarReplyCalendarMeetingRoomInstance = nil
}

// ReplyCalendarMeetingRoomInstanceReq ...
type ReplyCalendarMeetingRoomInstanceReq struct {
	RoomID       string `json:"room_id,omitempty"`       // 会议室的 ID
	Uid          string `json:"uid,omitempty"`           // 会议室的日程 ID
	OriginalTime int64  `json:"original_time,omitempty"` // 日程实例原始时间, 非重复日程必为0。重复日程若为0则表示回复其所有实例, 否则表示回复单个实例。
	Status       string `json:"status,omitempty"`        // 回复状态, NOT_CHECK_IN 表示未签到, ENDED_BEFORE_DUE 表示提前结束, ACCEPTED_BY_ADMIN 表示被管理员置为接受, DECLINED_BY_ADMIN 表示被管理员置为拒绝
}

// ReplyCalendarMeetingRoomInstanceResp ...
type ReplyCalendarMeetingRoomInstanceResp struct {
}

// replyCalendarMeetingRoomInstanceResp ...
type replyCalendarMeetingRoomInstanceResp struct {
	Code  int64                                 `json:"code,omitempty"` // 返回码, 非 0 表示失败。105003表示 original_time 非法, 此时可能是重复日程的整个开始时间被修改, 建议应用重新查询会议室日程实例列表, 获取最新的 original_time。
	Msg   string                                `json:"msg,omitempty"`  // 返回码的描述, "success" 表示成功, 其他为错误提示信息
	Data  *ReplyCalendarMeetingRoomInstanceResp `json:"data,omitempty"`
	Error *ErrorDetail                          `json:"error,omitempty"`
}
