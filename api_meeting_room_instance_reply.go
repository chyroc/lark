package lark

import (
	"context"
)

// ReplyInstance 该接口用于回复会议室日程实例，包括未签到释放和提前结束释放。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYzN4UjL2cDO14iN3gTN
func (r *MeetingRoomAPI) ReplyInstance(ctx context.Context, request *ReplyInstanceReq) (*ReplyInstanceResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/instance/reply",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(replyInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "ReplyInstance", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type ReplyInstanceReq struct {
	RoomID       string `json:"room_id,omitempty"`       // 会议室的 ID
	Uid          string `json:"uid,omitempty"`           // 会议室的日程 ID
	OriginalTime int    `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
	Status       string `json:"status,omitempty"`        // 回复状态，NOT_CHECK_IN 表示未签到，ENDED_BEFORE_DUE 表示提前结束
}

type replyInstanceResp struct {
	Code int                `json:"code,omitempty"` // 返回码，非 0 表示失败。105003表示 original_time 非法，此时可能是重复日程的整个开始时间被修改，建议应用重新查询会议室日程实例列表，获取最新的 original_time。
	Msg  string             `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *ReplyInstanceResp `json:"data,omitempty"`
}

type ReplyInstanceResp struct{}
