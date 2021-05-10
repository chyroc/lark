package lark

import (
	"context"
)

// BatchGetBuildingID 该接口用于删除建筑物（办公大楼）。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN
func (r *MeetingRoomAPI) BatchGetBuildingID(ctx context.Context, request *BatchGetBuildingIDReq) (*BatchGetBuildingIDResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/delete",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetBuildingIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "BatchGetBuildingID", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetBuildingIDReq struct {
	BuildingID string `json:"building_id,omitempty"` // 要删除的建筑ID
}

type batchGetBuildingIDResp struct {
	Code int                     `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetBuildingIDResp `json:"data,omitempty"`
}

type BatchGetBuildingIDResp struct{}
