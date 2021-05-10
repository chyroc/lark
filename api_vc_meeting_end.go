package lark

import (
	"context"
)

// EndMeeting
//
// > 结束一个进行中的会议
// 该会议正在进行中，且操作者具有相应的权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/end
func (r *VCAPI) EndMeeting(ctx context.Context, request *EndMeetingReq) (*EndMeetingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/end",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(endMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "EndMeeting", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type EndMeetingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议id，示例值："6911188411932033028"
}

type endMeetingResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *EndMeetingResp `json:"data,omitempty"`
}

type EndMeetingResp struct{}
