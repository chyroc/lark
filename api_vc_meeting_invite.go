package lark

import (
	"context"
)

// InviteMeeting
//
// > 邀请用户进入会议
// 操作者必须具有相应的权限（如果操作者为用户，则必须在会中），每次最多邀请10人；如果会议被锁定、或参会人数如果达到上限，则会邀请失败
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/invite
func (r *VCAPI) InviteMeeting(ctx context.Context, request *InviteMeetingReq) (*InviteMeetingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/invite",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(inviteMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "InviteMeeting", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type InviteMeetingReq struct {
	UserIDType IDType                     `query:"user_id_type" json:"-"` // 用户ID类型，可用值：【open_id，union_id，user_id】，默认值：open_id
	MeetingID  string                     `path:"meeting_id" json:"-"`    // 会议id，示例值："6911188411932033028"
	Invitees   []*InviteMeetingReqInvitee `json:"invitees,omitempty"`     // 被邀请的用户列表，必选字段
}

type InviteMeetingReqInvitee struct {
	ID       string `json:"id,omitempty"`        // 用户id
	UserType int    `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type inviteMeetingResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *InviteMeetingResp `json:"data,omitempty"` // -
}

type InviteMeetingResp struct {
	InviteResults []*InviteMeetingRespInviteResult `json:"invite_results,omitempty"` // 邀请结果
}

type InviteMeetingRespInviteResult struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int    `json:"user_type,omitempty"` // 用户类型，可用值：【1（lark用户），2（rooms用户），3（文档用户），4（neo单品用户），5（neo单品游客用户），6（pstn用户），7（sip用户）】
	Status   int    `json:"status,omitempty"`    // 邀请结果，可用值：【1（邀请成功），2（邀请失败）】
}
