package lark

import (
	"context"
)

// GetMemberList 如果用户在群中，则返回该群的成员列表。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 该接口不会返回群内的机器人成员
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/get
func (r *ChatAPI) GetMemberList(ctx context.Context, request *GetMemberListReq) (*GetMemberListResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMemberListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "GetMemberList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMemberListReq struct {
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 群成员 id 类型 open_id/user_id/union_id, 示例值："user_id", 可选值有: `user_id`：以 user_id 来识别成员, `union_id`：以 union_id 来识别成员, `open_id`：以 open_id 来识别成员
	PageToken    *string `query:"page_token" json:"-"`     // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ=="
	PageSize     *int    `query:"page_size" json:"-"`      // 分页大小, 示例值：10, 最大值：`100`
	ChatID       string  `path:"chat_id" json:"-"`         // 群 ID, 示例值："oc_a0553eda9014c201e6969b478895c230"
}

type getMemberListResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *GetMemberListResp `json:"data,omitempty"` //
}

type GetMemberListResp struct {
	Items     []*GetMemberListRespItem `json:"items,omitempty"`      // member 列表
	PageToken string                   `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                     `json:"has_more,omitempty"`   // 是否还有更多项
}

type GetMemberListRespItem struct {
	MemberIDType IDType `json:"member_id_type,omitempty"` // member id 类型
	MemberID     string `json:"member_id,omitempty"`      // member id
	Name         string `json:"name,omitempty"`           // 名字
}
