package lark

import (
	"context"
)

// SetHostMeeting
//
// > 设置会议的主持人
// 操作者必须具有相应的权限（如果操作者为用户，必须是会中当前主持人；如果操作者为应用，需传入会中当前主持人，并通过CAS校验）。如果操作失败可使用返回的当前主持人重试
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/set_host
func (r *VCAPI) SetHostMeeting(ctx context.Context, request *SetHostMeetingReq) (*SetHostMeetingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/set_host",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(setHostMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "SetHostMeeting", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SetHostMeetingReq struct {
	UserIDType  IDType                        `query:"user_id_type" json:"-"`  // 用户ID类型，可用值：【open_id，union_id，user_id】，默认值：open_id
	MeetingID   string                        `path:"meeting_id" json:"-"`     // 会议id，示例值："6911188411932033028"
	HostUser    *SetHostMeetingReqHostUser    `json:"host_user,omitempty"`     // 将要设置的主持人，必选字段
	OldHostUser *SetHostMeetingReqOldHostUser `json:"old_host_user,omitempty"` // 当前主持人，使用user_access_token可不填；使用tenant_access_token必填（CAS并发安全：如果该字段和会中当前主持人不符则会设置失败，可使用返回的最新数据重试）
}

type SetHostMeetingReqHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户id
	UserType int    `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type SetHostMeetingReqOldHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户id
	UserType int    `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type setHostMeetingResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *SetHostMeetingResp `json:"data,omitempty"` // -
}

type SetHostMeetingResp struct {
	HostUser *SetHostMeetingRespHostUser `json:"host_user,omitempty"` // 会中当前主持人
}

type SetHostMeetingRespHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户id
	UserType int    `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}
