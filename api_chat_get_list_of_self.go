package lark

import (
	"context"
)

// GetChatListOfSelf 获取用户或者机器人所在群列表。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/list
func (r *ChatAPI) GetChatListOfSelf(ctx context.Context, request *GetChatListOfSelfReq) (*GetChatListOfSelfResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(getChatListOfSelfResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "GetChatListOfSelf", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetChatListOfSelfReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果,**示例值**："dmJCRHhpd3JRbGV1VEVNRFFyTitRWDY5ZFkybmYrMEUwMUFYT0VMMWdENEtuYUhsNUxGMDIwemtvdE5ORjBNQQ=="
	PageSize   *int    `query:"page_size" json:"-"`    // 分页大小,**示例值**：10,**数据校验规则**：,- 最大值：`100`
}

type getChatListOfSelfResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetChatListOfSelfResp `json:"data,omitempty"` // \-
}

type GetChatListOfSelfResp struct {
	Items     []*GetChatListOfSelfRespItem `json:"items,omitempty"`      // chat 列表
	PageToken string                       `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
}

type GetChatListOfSelfRespItem struct {
	ChatID      string `json:"chat_id,omitempty"`       // 群组 ID
	Avatar      string `json:"avatar,omitempty"`        // 群头像 URL
	Name        string `json:"name,omitempty"`          // 群名称
	Description string `json:"description,omitempty"`   // 群描述
	OwnerID     string `json:"owner_id,omitempty"`      // 群主 ID
	OwnerIDType IDType `json:"owner_id_type,omitempty"` // 群主 ID 类型
}
