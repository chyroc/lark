package lark

import (
	"context"
)

// UpdateRoom 该接口用于更新会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMTNwYjLzUDM24yM1AjN
func (r *MeetingRoomAPI) UpdateRoom(ctx context.Context, request *UpdateRoomReq) (*UpdateRoomResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/update",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "UpdateRoom", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
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
