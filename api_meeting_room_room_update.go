// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateRoom 该接口用于更新会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMTNwYjLzUDM24yM1AjN
func (r *MeetingRoomService) UpdateRoom(ctx context.Context, request *UpdateRoomReq, options ...MethodOptionFunc) (*UpdateRoomResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomUpdateRoom != nil {
		r.cli.logDebug(ctx, "[lark] MeetingRoom#UpdateRoom mock enable")
		return r.cli.mock.mockMeetingRoomUpdateRoom(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] MeetingRoom#UpdateRoom call api")
	r.cli.logDebug(ctx, "[lark] MeetingRoom#UpdateRoom request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] MeetingRoom#UpdateRoom POST https://open.feishu.cn/open-apis/meeting_room/room/update failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] MeetingRoom#UpdateRoom POST https://open.feishu.cn/open-apis/meeting_room/room/update failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("MeetingRoom", "UpdateRoom", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] MeetingRoom#UpdateRoom request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomUpdateRoom(f func(ctx context.Context, request *UpdateRoomReq, options ...MethodOptionFunc) (*UpdateRoomResp, *Response, error)) {
	r.mockMeetingRoomUpdateRoom = f
}

func (r *Mock) UnMockMeetingRoomUpdateRoom() {
	r.mockMeetingRoomUpdateRoom = nil
}

type UpdateRoomReq struct {
	RoomID       string  `json:"room_id,omitempty"`        // 要更新的会议室ID
	Name         *string `json:"name,omitempty"`           // 会议室名称
	Capacity     *int    `json:"capacity,omitempty"`       // 容量
	IsDisabled   *bool   `json:"is_disabled,omitempty"`    // 是否禁用
	CustomRoomID *string `json:"custom_room_id,omitempty"` // 租户自定义会议室ID
}

type updateRoomResp struct {
	Code int             `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *UpdateRoomResp `json:"data,omitempty"`
}

type UpdateRoomResp struct{}
