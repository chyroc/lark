package lark

import (
	"context"
	"fmt"
)

type EventType string

const (
	EventTypeIMMessageReceiveV1 EventType = "im.message.receive_v1"
	EventTypeIMMessageReadV1    EventType = "im.message.message_read_v1"
)

type eventHandler struct {
	eventIMMessageReceiveV1Handler eventIMMessageReceiveV1Handler
	eventIMMessageReadV1Handler    eventIMMessageReadV1Handler
}

type eventBody struct {
	eventIMMessageReceiveV1 *EventIMMessageReceiveV1
	eventIMMessageReadV1    *EventIMMessageReadV1
}

func (r *EventCallbackAPI) parserEventV2(req *eventReq) error {
	if req.Header == nil {
		return fmt.Errorf("get schema=2.0, but header is nil")
	}

	switch req.Header.EventType {
	case EventTypeIMMessageReceiveV1:
		event := new(EventIMMessageReceiveV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMMessageReceiveV1 = event
	case EventTypeIMMessageReadV1:
		event := new(EventIMMessageReadV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMMessageReadV1 = event
	}

	return nil
}

func (r *EventCallbackAPI) handlerEventV2(ctx context.Context, req *eventReq) (handled bool, s string, err error) {
	switch {
	case req.eventIMMessageReceiveV1 != nil:
		if r.cli.eventHandler.eventIMMessageReceiveV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMMessageReceiveV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMMessageReceiveV1)
		}
		return true, s, err
	case req.eventIMMessageReadV1 != nil:
		if r.cli.eventHandler.eventIMMessageReadV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMMessageReadV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMMessageReadV1)
		}
		return true, s, err
	}
	return false, "", nil
}
