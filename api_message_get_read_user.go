package lark

import (
	"context"
)

// GetMessageReadUser 查询消息的已读用户信息。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 只能查询机器人自己发送，且发送时间不超过7天的消息
// - 查询消息已读信息时机器人仍需要在会话内
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/read_users
func (r *MessageAPI) GetMessageReadUser(ctx context.Context, request *GetMessageReadUserReq) (*GetMessageReadUserResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id/read_users",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMessageReadUserResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "GetMessageReadUser", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMessageReadUserReq struct {
	UserIDType IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	PageSize   *int    `query:"page_size" json:"-"`    // 此次调用中使用的分页的大小,**示例值**：20,**数据校验规则**：,- 取值范围：`1` ～ `100`
	PageToken  *string `query:"page_token" json:"-"`   // 下一页分页的token,**示例值**："GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ=="
	MessageID  string  `path:"message_id" json:"-"`    // 待查询的消息的ID,**示例值**："om_dc13264520392913993dd051dba21dcf"
}

type getMessageReadUserResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetMessageReadUserResp `json:"data,omitempty"` // \-
}

type GetMessageReadUserResp struct {
	Items     []*GetMessageReadUserRespItem `json:"items,omitempty"`      // -
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有下一页
	PageToken string                        `json:"page_token,omitempty"` // 下一页分页的token
}

type GetMessageReadUserRespItem struct {
	UserIDType IDType `json:"user_id_type,omitempty"` // 用户id类型
	UserID     string `json:"user_id,omitempty"`      // 用户id
	Timestamp  string `json:"timestamp,omitempty"`    // 阅读时间
}
