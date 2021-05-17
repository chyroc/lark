// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreatePublicMailboxMember 向公共邮箱添加单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/create
func (r *MailService) CreatePublicMailboxMember(ctx context.Context, request *CreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*CreatePublicMailboxMemberResp, *Response, error) {
	if r.cli.mock.mockMailCreatePublicMailboxMember != nil {
		r.cli.logDebug(ctx, "[lark] Mail#CreatePublicMailboxMember mock enable")
		return r.cli.mock.mockMailCreatePublicMailboxMember(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Mail#CreatePublicMailboxMember call api")
	r.cli.logDebug(ctx, "[lark] Mail#CreatePublicMailboxMember request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createPublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Mail#CreatePublicMailboxMember POST https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Mail#CreatePublicMailboxMember POST https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Mail", "CreatePublicMailboxMember", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Mail#CreatePublicMailboxMember request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMailCreatePublicMailboxMember(f func(ctx context.Context, request *CreatePublicMailboxMemberReq, options ...MethodOptionFunc) (*CreatePublicMailboxMemberResp, *Response, error)) {
	r.mockMailCreatePublicMailboxMember = f
}

func (r *Mock) UnMockMailCreatePublicMailboxMember() {
	r.mockMailCreatePublicMailboxMember = nil
}

type CreatePublicMailboxMemberReq struct {
	UserIDType      *IDType       `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	PublicMailboxID string        `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	UserID          *string       `json:"user_id,omitempty"`          // 租户内用户的唯一标识（当成员类型是USER时有值）, 示例值："xxxxxxxxxx"
	Type            *MailUserType `json:"type,omitempty"`             // 成员类型, 示例值："USER", 可选值有: `USER`：内部用户
}

type createPublicMailboxMemberResp struct {
	Code int                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *CreatePublicMailboxMemberResp `json:"data,omitempty"` //
}

type CreatePublicMailboxMemberResp struct {
	MemberID string       `json:"member_id,omitempty"` // 公共邮箱内成员唯一标识
	UserID   string       `json:"user_id,omitempty"`   // 租户内用户的唯一标识（当成员类型是USER时有值）
	Type     MailUserType `json:"type,omitempty"`      // 成员类型, 可选值有: `USER`：内部用户
}
