package lark

import (
	"context"
)

// UpdateMessage 更新应用已发送的消息卡片内容。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 当前仅更新支持卡片消息
// - 只支持对所有人都更新的「共享卡片」。如果你只想更新特定人的消息卡片，必须要用户在卡片操作交互后触发，开发文档参考[「独享卡片」](https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN#49904b71)
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch
func (r *MessageAPI) UpdateMessage(ctx context.Context, request *UpdateMessageReq) (*UpdateMessageResp, *Response, error) {
	req := &requestParam{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateMessageResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "UpdateMessage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待更新的消息的ID,**示例值**："om_dc13264520392913993dd051dba21dcf"
	Content   string `json:"content,omitempty"`   // 消息内容 json 格式，[发送消息 content 说明](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json)，参考文档中的卡片格式|,**示例值**："参考链接"
}

type updateMessageResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateMessageResp `json:"data,omitempty"`
}

type UpdateMessageResp struct{}
