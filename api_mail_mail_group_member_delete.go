// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteMailGroupMember 删除邮件组单个成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/delete
func (r *MailService) DeleteMailGroupMember(ctx context.Context, request *DeleteMailGroupMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupMemberResp, *Response, error) {
	if r.cli.mock.mockMailDeleteMailGroupMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailGroupMember mock enable")
		return r.cli.mock.mockMailDeleteMailGroupMember(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Mail#DeleteMailGroupMember call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailGroupMember request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteMailGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Mail#DeleteMailGroupMember DELETE https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Mail#DeleteMailGroupMember DELETE https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Mail", "DeleteMailGroupMember", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailGroupMember success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMailDeleteMailGroupMember(f func(ctx context.Context, request *DeleteMailGroupMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupMemberResp, *Response, error)) {
	r.mockMailDeleteMailGroupMember = f
}

func (r *Mock) UnMockMailDeleteMailGroupMember() {
	r.mockMailDeleteMailGroupMember = nil
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
