package lark

import (
	"context"
)

// DeleteMailGroupMember 删除邮件组单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/delete
func (r *MailAPI) DeleteMailGroupMember(ctx context.Context, request *DeleteMailGroupMemberReq) (*DeleteMailGroupMemberResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteMailGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "DeleteMailGroupMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteMailGroupMemberReq struct {
	MailGroupID string `path:"mailgroup_id" json:"-"` // The unique ID or email address of a mail group, 示例值："xxxxxxxxxxxxxxx or test_mail_group@xxx.xx"
	MemberID    string `path:"member_id" json:"-"`    // The unique ID of a member in this mail group, 示例值："xxxxxxxxxxxxxxx"
}

type deleteMailGroupMemberResp struct {
	Code int                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailGroupMemberResp `json:"data,omitempty"`
}

type DeleteMailGroupMemberResp struct{}
