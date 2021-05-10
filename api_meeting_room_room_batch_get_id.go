package lark

import (
	"context"
)

// BatchGetRoomID 该接口用于根据租户自定义会议室ID查询会议室ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYzMxYjL2MTM24iNzEjN
func (r *MeetingRoomAPI) BatchGetRoomID(ctx context.Context, request *BatchGetRoomIDReq) (*BatchGetRoomIDResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/batch_get_id?custom_room_ids=test01&custom_room_ids=test02",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetRoomIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "BatchGetRoomID", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetRoomIDReq struct {
	CustomRoomIDs string `query:"custom_room_ids" json:"-"` // 用于查询指定会议室的租户自定义会议室ID
}

type batchGetRoomIDResp struct {
	Code int                 `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetRoomIDResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetRoomIDResp struct {
	Rooms *BatchGetRoomIDRespRooms `json:"rooms,omitempty"` // 会议室列表
}

type BatchGetRoomIDRespRooms struct {
	RoomID       string `json:"room_id,omitempty"`        // 会议室 ID
	CustomRoomID string `json:"custom_room_id,omitempty"` // 租户自定义会议室 ID
}
