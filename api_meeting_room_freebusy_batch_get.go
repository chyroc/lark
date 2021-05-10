package lark

import (
	"context"
)

// BatchGetFreebusy 该接口用于获取指定会议室的忙闲日程实例列表。非重复日程只有唯一实例；重复日程可能存在多个实例，依据重复规则和时间范围扩展。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDOyUjLygjM14iM4ITN
func (r *MeetingRoomAPI) BatchGetFreebusy(ctx context.Context, request *BatchGetFreebusyReq) (*BatchGetFreebusyResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/freebusy/batch_get?room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&room_ids=omm_eada1d61a550955240c28757e7dec3af&time_min=2019-09-04T08:45:00%2B08:00&time_max=2019-09-04T09:45:00%2B08:00",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetFreebusyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "BatchGetFreebusy", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetFreebusyReq struct {
	RoomIDs string `query:"room_ids" json:"-"` // 用于查询指定会议室的 ID
	TimeMin string `query:"time_min" json:"-"` // 查询会议室忙闲的起始时间，需要遵循格式 [RFC3339](https://tools.ietf.org/html/rfc3339)，需要进行URL Encode
	TimeMax string `query:"time_max" json:"-"` // 查询会议室忙闲的结束时间，需要遵循格式 [RFC3339](https://tools.ietf.org/html/rfc3339)，需要进行URL Encode
}

type batchGetFreebusyResp struct {
	Code int                   `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetFreebusyResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetFreebusyResp struct {
	TimeMin  string                        `json:"time_min,omitempty"`  // 查询会议室忙闲的起始时间，与请求参数完全相同
	TimeMax  string                        `json:"time_max,omitempty"`  // 查询会议室忙闲的结束时间，与请求参数完全相同
	FreeBusy *BatchGetFreebusyRespFreeBusy `json:"free_busy,omitempty"` // 会议室忙闲列表
}

type BatchGetFreebusyRespFreeBusy struct {
	RoomID string `json:"room_id,omitempty"` // 与请求合法参数相同，表示之后是对应会议室的忙闲状态
}
