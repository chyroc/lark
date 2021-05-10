package lark

import (
	"context"
)

// DeleteMailGroup 删除一个邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/delete
func (r *MailAPI) DeleteMailGroup(ctx context.Context, request *DeleteMailGroupReq) (*DeleteMailGroupResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups/:mailgroup_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteMailGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "DeleteMailGroup", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteMailGroupReq struct {
	MailGroupID string `path:"mailgroup_id" json:"-"` // 邮件组ID或者邮件组地址, 示例值："xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
}

type deleteMailGroupResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailGroupResp `json:"data,omitempty"`
}

type DeleteMailGroupResp struct{}
