package lark

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b

func (r *EventCallbackAPI) HandlerMessageReceive(f eventTypeImageReceiveHandler) {
	r.cli.eventHandler.eventTypeImageReceiveHandler = f
}

func (r *EventCallbackAPI) ListenCallback(ctx context.Context, reader io.Reader, writer io.Writer) error {
	eventReq := new(eventReq)

	bs, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println(string(bs))
	fmt.Println()
	if err := json.Unmarshal(bs, &eventReq); err != nil {
		return err
	}

	spew.Dump(eventReq)

	switch {
	case eventReq.Type == "url_verification":
		_, err = writer.Write([]byte(fmt.Sprintf(`{"challenge":%q}`, eventReq.Challenge)))
		return err
	case eventReq.Schema == "2.0":
		if eventReq.Header == nil {
			return fmt.Errorf("get schema=2.0, but header is nil")
		}
		switch eventReq.Header.EventType {
		case EventTypeMessageReceive:
			bs, err := json.Marshal(eventReq.Event)
			if err != nil {
				return err
			}
			event := new(EventMessageReceive)
			err = json.Unmarshal(bs, event)
			if err != nil {
				return err
			}
			if r.cli.eventHandler.eventTypeImageReceiveHandler != nil {
				r.cli.eventHandler.eventTypeImageReceiveHandler(ctx, r.cli, writer, eventReq.Schema, eventReq.Header, event)
			}
		}
	}

	return nil
}

type eventReq struct {
	// v2 字段
	Schema string       `json:"schema"` // 2.0
	Header *EventHeader `json:"header"`

	// v1 字段
	UUID      string `json:"uuid"`
	Token     string `json:"token"`
	Ts        string `json:"ts"`
	Type      string `json:"type"` // event_callback,
	Challenge string `json:"challenge"`

	// 通用字段
	Event interface{} `json:"event"`
}

type EventHeader struct {
	EventID    string    `json:"event_id,omitempty"`    // 事件 ID
	EventType  EventType `json:"event_type,omitempty"`  // 事件类型
	CreateTime string    `json:"create_time,omitempty"` // 事件创建时间戳（单位：毫秒）
	Token      string    `json:"token,omitempty"`       // 事件 Token
	AppID      string    `json:"app_id,omitempty"`      // 应用 ID
	TenantKey  string    `json:"tenant_key,omitempty"`  // 租户 Key
}
