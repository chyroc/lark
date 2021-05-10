package lark

import (
	"context"
)

// GetMailGroupList 分页批量获取邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/list
func (r *MailAPI) GetMailGroupList(ctx context.Context, request *GetMailGroupListReq) (*GetMailGroupListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/mailgroups",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMailGroupListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "GetMailGroupList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMailGroupListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："xxx"
	PageSize  *int    `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`200`
}

type getMailGroupListResp struct {
	Code int                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupListResp `json:"data,omitempty"` //
}

type GetMailGroupListResp struct {
	HasMore   bool                        `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                      `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetMailGroupListRespItem `json:"items,omitempty"`      // 邮件组列表
}

type GetMailGroupListRespItem struct {
	MailgroupID             string `json:"mailgroup_id,omitempty"`               // 邮件组ID
	Email                   string `json:"email,omitempty"`                      // 邮件组地址
	Name                    string `json:"name,omitempty"`                       // 邮件组名称
	Description             string `json:"description,omitempty"`                // 邮件组描述
	DirectMembersCount      string `json:"direct_members_count,omitempty"`       // 邮件组成员数量
	IncludeExternalMember   bool   `json:"include_external_member,omitempty"`    // 是否包含外部成员
	IncludeAllCompanyMember bool   `json:"include_all_company_member,omitempty"` // 是否是全员邮件组
	WhoCanSendMail          string `json:"who_can_send_mail,omitempty"`          // 谁可发送邮件到此邮件组, 可选值有: `ANYONE`：任何人, `ALL_INTERNAL_USERS`：仅组织内部成员, `ALL_GROUP_MEMBERS`：仅邮件组成员, `CUSTOM_MEMBERS`：自定义成员
}
