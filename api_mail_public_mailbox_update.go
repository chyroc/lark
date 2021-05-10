package lark

import (
	"context"
)

// UpdatePublicMailbox 更新公共邮箱所有信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update
func (r *MailAPI) UpdatePublicMailbox(ctx context.Context, request *UpdatePublicMailboxReq) (*UpdatePublicMailboxResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updatePublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "UpdatePublicMailbox", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdatePublicMailboxReq struct {
	PublicMailboxID string  `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	Name            *string `json:"name,omitempty"`             // 公共邮箱名称, 示例值："test public mailbox"
}

type updatePublicMailboxResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdatePublicMailboxResp `json:"data,omitempty"` //
}

type UpdatePublicMailboxResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}
