package lark

import (
	"context"
	"encoding/json"
)

// HandlerEventCard
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
func (r *EventCallbackService) HandlerEventCard(f eventCardHandler) {
	r.cli.eventHandler.eventCardHandler = f
}

type eventCardHandler func(ctx context.Context, cli *Lark, event *EventCardCallback) (string, error)

type EventCardCallback struct {
	RefreshToken string `json:"refresh_token"` // header: X-Refresh-Token

	OpenID        string                   `json:"open_id"`
	UserID        string                   `json:"user_id"`
	OpenMessageID string                   `json:"open_message_id"`
	TenantKey     string                   `json:"tenant_key"`
	Token         string                   `json:"token"` // 更新卡片用的token(凭证)
	Action        *EventCardCallbackAction `json:"action"`
}

type EventCardCallbackAction struct {
	Value  json.RawMessage `json:"value"`  // 交互元素的value字段值
	Tag    string          `json:"tag"`    // 交互元素的tag字段值
	Option string          `json:"option"` // 选中option的value（button元素不适用）
}
