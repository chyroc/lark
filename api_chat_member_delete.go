package lark

import (
	"context"
)

// DeleteMember 将用户或机器人移出群聊。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 用户或机器人在任何条件下均可移除自己出群（即主动退群）；但仅有群主 及 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可以移除其他用户或者机器人
// - 每次请求，最多移除50个用户或者5个机器人
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/delete
func (r *ChatAPI) DeleteMember(ctx context.Context, request *DeleteMemberReq) (*DeleteMemberResp, *Response, error) {
	req := &requestParam{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
	}
	resp := new(deleteMemberResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "DeleteMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteMemberReq struct {
	MemberIDType *IDType  `query:"member_id_type" json:"-"` // 出群成员 id 类型 open_id/user_id/union_id/app_id,**示例值**："user_id",**可选值有**：,- `user_id`：以 user_id 来识别成员,- `union_id`：以 union_id 来识别成员,- `open_id`：以 open_id 来识别成员,- `app_id`：以 app_id 来识别成员
	ChatID       string   `path:"chat_id" json:"-"`         // 群 ID,**示例值**："oc_a0553eda9014c201e6969b478895c230"
	IDList       []string `json:"id_list,omitempty"`        // 成员列表
}

type deleteMemberResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMemberResp `json:"data,omitempty"` // \-
}

type DeleteMemberResp struct {
	InvalidIDList []string `json:"invalid_id_list,omitempty"` // 无效成员列表
}
