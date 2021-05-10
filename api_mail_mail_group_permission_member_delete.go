package lark

import (
	"context"
)

// DeleteMailGroupPermissionMember 从自定义成员中删除单个成员，删除后该成员无法发送邮件到该邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/delete
func (r *MailAPI) DeleteMailGroupPermissionMember(ctx context.Context, request *DeleteMailGroupPermissionMemberReq) (*DeleteMailGroupPermissionMemberResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/:permission_member_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteMailGroupPermissionMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "DeleteMailGroupPermissionMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteMailGroupPermissionMemberReq struct {
	MailgroupID        string `path:"mailgroup_id" json:"-"`         // The unique ID or email address of a mail group, 示例值："xxxxxxxxxxxxxxx or test_mail_group@xxx.xx"
	PermissionMemberID string `path:"permission_member_id" json:"-"` // The unique ID of a member in this permission group, 示例值："xxxxxxxxxxxxxxx"
}

type deleteMailGroupPermissionMemberResp struct {
	Code int                                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailGroupPermissionMemberResp `json:"data,omitempty"`
}

type DeleteMailGroupPermissionMemberResp struct{}
