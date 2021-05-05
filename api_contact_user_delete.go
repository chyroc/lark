package lark

import (
	"context"
)

// DeleteUser 该接口向通讯录删除一个用户信息。
//
// 应用需要待删除用户的所有部门的通讯录权限才能删除该用户。应用商店应用无权限调用接口。用户可以在删除员工时设置删除员工数据的接收者，如果不设置则由其leader接受，如果该员工没有leader，则会将该员工的数据删除。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/delete
func (r *ContactAPI) DeleteUser(ctx context.Context, request *DeleteUserReq) (*DeleteUserResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/users/:user_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteUserResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Contact", "DeleteUser", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteUserReq struct {
	UserIDType                   *IDType `query:"user_id_type" json:"-"`                     // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	UserID                       string  `path:"user_id" json:"-"`                           // 用户ID，需要与查询参数中的user_id_type类型保持一致。,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	DepartmentChatAcceptorUserID *string `json:"department_chat_acceptor_user_id,omitempty"` // 用户部门群聊接收者,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	ExternalChatAcceptorUserID   *string `json:"external_chat_acceptor_user_id,omitempty"`   // 用户外部群聊接收者,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	DocsAcceptorUserID           *string `json:"docs_acceptor_user_id,omitempty"`            // 用户文档接收者,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	CalendarAcceptorUserID       *string `json:"calendar_acceptor_user_id,omitempty"`        // 用户日历接收者,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	ApplicationAcceptorUserID    *string `json:"application_acceptor_user_id,omitempty"`     // 用户应用接收者,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	HelpdeskAcceptorUserID       *string `json:"helpdesk_acceptor_user_id,omitempty"`        // 用户超文本帮助信息接收者,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
}

type deleteUserResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteUserResp `json:"data,omitempty"`
}

type DeleteUserResp struct{}
