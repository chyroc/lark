package lark

import (
	"context"
	"encoding/json"
	"fmt"
)

type EventType string

const (
	EventTypeV2IMMessageReceiveV1          EventType = "im.message.receive_v1"
	EventTypeV2IMMessageReadV1             EventType = "im.message.message_read_v1"
	EventTypeV2IMChatDisbandedV1           EventType = "im.chat.disbanded_v1"
	EventTypeV2IMChatUpdatedV1             EventType = "im.chat.updated_v1"
	EventTypeV2IMChatMemberBotAddedV1      EventType = "im.chat.member.bot.added_v1"
	EventTypeV2IMChatMemberBotDeletedV1    EventType = "im.chat.member.bot.deleted_v1"
	EventTypeV2IMChatMemberUserAddedV1     EventType = "im.chat.member.user.added_v1"
	EventTypeV2IMChatMemberUserWithdrawnV1 EventType = "im.chat.member.user.withdrawn_v1"
	EventTypeV2IMChatMemberUserDeletedV1   EventType = "im.chat.member.user.deleted_v1"
	EventTypeV2VCMeetingMeetingStartedV1   EventType = "vc.meeting.meeting_started_v1"
	EventTypeV2VCMeetingMeetingEndedV1     EventType = "vc.meeting.meeting_ended_v1"
	EventTypeV2VCMeetingJoinMeetingV1      EventType = "vc.meeting.join_meeting_v1"
	EventTypeV2VCMeetingLeaveMeetingV1     EventType = "vc.meeting.leave_meeting_v1"
	EventTypeV2VCMeetingRecordingStartedV1 EventType = "vc.meeting.recording_started_v1"
	EventTypeV2VCMeetingRecordingEndedV1   EventType = "vc.meeting.recording_ended_v1"
	EventTypeV2VCMeetingRecordingReadyV1   EventType = "vc.meeting.recording_ready_v1"
	EventTypeV2VCMeetingShareStartedV1     EventType = "vc.meeting.share_started_v1"
	EventTypeV2VCMeetingShareEndedV1       EventType = "vc.meeting.share_ended_v1"
	EventTypeV1P2PChatCreate               EventType = "p2p_chat_create"
	EventTypeV1AddUserToChat               EventType = "add_user_to_chat"
)

type eventHandler struct {
	eventV2IMMessageReceiveV1Handler          eventV2IMMessageReceiveV1Handler
	eventV2IMMessageReadV1Handler             eventV2IMMessageReadV1Handler
	eventV2IMChatDisbandedV1Handler           eventV2IMChatDisbandedV1Handler
	eventV2IMChatUpdatedV1Handler             eventV2IMChatUpdatedV1Handler
	eventV2IMChatMemberBotAddedV1Handler      eventV2IMChatMemberBotAddedV1Handler
	eventV2IMChatMemberBotDeletedV1Handler    eventV2IMChatMemberBotDeletedV1Handler
	eventV2IMChatMemberUserAddedV1Handler     eventV2IMChatMemberUserAddedV1Handler
	eventV2IMChatMemberUserWithdrawnV1Handler eventV2IMChatMemberUserWithdrawnV1Handler
	eventV2IMChatMemberUserDeletedV1Handler   eventV2IMChatMemberUserDeletedV1Handler
	eventV2VCMeetingMeetingStartedV1Handler   eventV2VCMeetingMeetingStartedV1Handler
	eventV2VCMeetingMeetingEndedV1Handler     eventV2VCMeetingMeetingEndedV1Handler
	eventV2VCMeetingJoinMeetingV1Handler      eventV2VCMeetingJoinMeetingV1Handler
	eventV2VCMeetingLeaveMeetingV1Handler     eventV2VCMeetingLeaveMeetingV1Handler
	eventV2VCMeetingRecordingStartedV1Handler eventV2VCMeetingRecordingStartedV1Handler
	eventV2VCMeetingRecordingEndedV1Handler   eventV2VCMeetingRecordingEndedV1Handler
	eventV2VCMeetingRecordingReadyV1Handler   eventV2VCMeetingRecordingReadyV1Handler
	eventV2VCMeetingShareStartedV1Handler     eventV2VCMeetingShareStartedV1Handler
	eventV2VCMeetingShareEndedV1Handler       eventV2VCMeetingShareEndedV1Handler
	eventV1P2PChatCreateHandler               eventV1P2PChatCreateHandler
	eventV1AddUserToChatHandler               eventV1AddUserToChatHandler
}

func (r *eventHandler) clone() *eventHandler {
	return &eventHandler{
		eventV2IMMessageReceiveV1Handler:          r.eventV2IMMessageReceiveV1Handler,
		eventV2IMMessageReadV1Handler:             r.eventV2IMMessageReadV1Handler,
		eventV2IMChatDisbandedV1Handler:           r.eventV2IMChatDisbandedV1Handler,
		eventV2IMChatUpdatedV1Handler:             r.eventV2IMChatUpdatedV1Handler,
		eventV2IMChatMemberBotAddedV1Handler:      r.eventV2IMChatMemberBotAddedV1Handler,
		eventV2IMChatMemberBotDeletedV1Handler:    r.eventV2IMChatMemberBotDeletedV1Handler,
		eventV2IMChatMemberUserAddedV1Handler:     r.eventV2IMChatMemberUserAddedV1Handler,
		eventV2IMChatMemberUserWithdrawnV1Handler: r.eventV2IMChatMemberUserWithdrawnV1Handler,
		eventV2IMChatMemberUserDeletedV1Handler:   r.eventV2IMChatMemberUserDeletedV1Handler,
		eventV2VCMeetingMeetingStartedV1Handler:   r.eventV2VCMeetingMeetingStartedV1Handler,
		eventV2VCMeetingMeetingEndedV1Handler:     r.eventV2VCMeetingMeetingEndedV1Handler,
		eventV2VCMeetingJoinMeetingV1Handler:      r.eventV2VCMeetingJoinMeetingV1Handler,
		eventV2VCMeetingLeaveMeetingV1Handler:     r.eventV2VCMeetingLeaveMeetingV1Handler,
		eventV2VCMeetingRecordingStartedV1Handler: r.eventV2VCMeetingRecordingStartedV1Handler,
		eventV2VCMeetingRecordingEndedV1Handler:   r.eventV2VCMeetingRecordingEndedV1Handler,
		eventV2VCMeetingRecordingReadyV1Handler:   r.eventV2VCMeetingRecordingReadyV1Handler,
		eventV2VCMeetingShareStartedV1Handler:     r.eventV2VCMeetingShareStartedV1Handler,
		eventV2VCMeetingShareEndedV1Handler:       r.eventV2VCMeetingShareEndedV1Handler,
		eventV1P2PChatCreateHandler:               r.eventV1P2PChatCreateHandler,
		eventV1AddUserToChatHandler:               r.eventV1AddUserToChatHandler,
	}
}

type eventBody struct {
	eventV2IMMessageReceiveV1          *EventV2IMMessageReceiveV1
	eventV2IMMessageReadV1             *EventV2IMMessageReadV1
	eventV2IMChatDisbandedV1           *EventV2IMChatDisbandedV1
	eventV2IMChatUpdatedV1             *EventV2IMChatUpdatedV1
	eventV2IMChatMemberBotAddedV1      *EventV2IMChatMemberBotAddedV1
	eventV2IMChatMemberBotDeletedV1    *EventV2IMChatMemberBotDeletedV1
	eventV2IMChatMemberUserAddedV1     *EventV2IMChatMemberUserAddedV1
	eventV2IMChatMemberUserWithdrawnV1 *EventV2IMChatMemberUserWithdrawnV1
	eventV2IMChatMemberUserDeletedV1   *EventV2IMChatMemberUserDeletedV1
	eventV2VCMeetingMeetingStartedV1   *EventV2VCMeetingMeetingStartedV1
	eventV2VCMeetingMeetingEndedV1     *EventV2VCMeetingMeetingEndedV1
	eventV2VCMeetingJoinMeetingV1      *EventV2VCMeetingJoinMeetingV1
	eventV2VCMeetingLeaveMeetingV1     *EventV2VCMeetingLeaveMeetingV1
	eventV2VCMeetingRecordingStartedV1 *EventV2VCMeetingRecordingStartedV1
	eventV2VCMeetingRecordingEndedV1   *EventV2VCMeetingRecordingEndedV1
	eventV2VCMeetingRecordingReadyV1   *EventV2VCMeetingRecordingReadyV1
	eventV2VCMeetingShareStartedV1     *EventV2VCMeetingShareStartedV1
	eventV2VCMeetingShareEndedV1       *EventV2VCMeetingShareEndedV1
	eventV1P2PChatCreate               *EventV1P2PChatCreate
	eventV1AddUserToChat               *EventV1AddUserToChat
}

func (r *EventCallbackAPI) parserEventV2(req *eventReq) error {
	if req.Header == nil {
		return fmt.Errorf("get schema=2.0, but header is nil")
	}

	switch req.Header.EventType {
	case EventTypeV2IMMessageReceiveV1:
		event := new(EventV2IMMessageReceiveV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMMessageReceiveV1 = event
	case EventTypeV2IMMessageReadV1:
		event := new(EventV2IMMessageReadV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMMessageReadV1 = event
	case EventTypeV2IMChatDisbandedV1:
		event := new(EventV2IMChatDisbandedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatDisbandedV1 = event
	case EventTypeV2IMChatUpdatedV1:
		event := new(EventV2IMChatUpdatedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatUpdatedV1 = event
	case EventTypeV2IMChatMemberBotAddedV1:
		event := new(EventV2IMChatMemberBotAddedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatMemberBotAddedV1 = event
	case EventTypeV2IMChatMemberBotDeletedV1:
		event := new(EventV2IMChatMemberBotDeletedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatMemberBotDeletedV1 = event
	case EventTypeV2IMChatMemberUserAddedV1:
		event := new(EventV2IMChatMemberUserAddedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatMemberUserAddedV1 = event
	case EventTypeV2IMChatMemberUserWithdrawnV1:
		event := new(EventV2IMChatMemberUserWithdrawnV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatMemberUserWithdrawnV1 = event
	case EventTypeV2IMChatMemberUserDeletedV1:
		event := new(EventV2IMChatMemberUserDeletedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2IMChatMemberUserDeletedV1 = event
	case EventTypeV2VCMeetingMeetingStartedV1:
		event := new(EventV2VCMeetingMeetingStartedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingMeetingStartedV1 = event
	case EventTypeV2VCMeetingMeetingEndedV1:
		event := new(EventV2VCMeetingMeetingEndedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingMeetingEndedV1 = event
	case EventTypeV2VCMeetingJoinMeetingV1:
		event := new(EventV2VCMeetingJoinMeetingV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingJoinMeetingV1 = event
	case EventTypeV2VCMeetingLeaveMeetingV1:
		event := new(EventV2VCMeetingLeaveMeetingV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingLeaveMeetingV1 = event
	case EventTypeV2VCMeetingRecordingStartedV1:
		event := new(EventV2VCMeetingRecordingStartedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingRecordingStartedV1 = event
	case EventTypeV2VCMeetingRecordingEndedV1:
		event := new(EventV2VCMeetingRecordingEndedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingRecordingEndedV1 = event
	case EventTypeV2VCMeetingRecordingReadyV1:
		event := new(EventV2VCMeetingRecordingReadyV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingRecordingReadyV1 = event
	case EventTypeV2VCMeetingShareStartedV1:
		event := new(EventV2VCMeetingShareStartedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingShareStartedV1 = event
	case EventTypeV2VCMeetingShareEndedV1:
		event := new(EventV2VCMeetingShareEndedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventV2VCMeetingShareEndedV1 = event
	}

	return nil
}

// https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/bot-events
func (r *EventCallbackAPI) parserEventV1(req *eventReq) error {
	if req.UUID == "" {
		return fmt.Errorf("get schema=1.0, but uuid is nil")
	}

	bs, err := json.Marshal(req.Event)
	if err != nil {
		return err
	}

	v1type := new(v1type)
	if err = json.Unmarshal(bs, v1type); err != nil {
		return err
	}

	switch v1type.Type {
	case EventTypeV1P2PChatCreate:
		event := new(EventV1P2PChatCreate)
		if err := json.Unmarshal(bs, event); err != nil {
			return fmt.Errorf("lark event unmarshal event %s failed", bs)
		}

		req.eventV1P2PChatCreate = event
	case EventTypeV1AddUserToChat:
		event := new(EventV1AddUserToChat)
		if err := json.Unmarshal(bs, event); err != nil {
			return fmt.Errorf("lark event unmarshal event %s failed", bs)
		}

		req.eventV1AddUserToChat = event
	}

	return nil
}

type v1type struct {
	Type EventType `json:"type"`
}

func (r *EventCallbackAPI) handlerEventV2(ctx context.Context, req *eventReq) (handled bool, s string, err error) {
	switch {
	case req.eventV2IMMessageReceiveV1 != nil:
		if r.cli.eventHandler.eventV2IMMessageReceiveV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMMessageReceiveV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMMessageReceiveV1)
		}
		return true, s, err
	case req.eventV2IMMessageReadV1 != nil:
		if r.cli.eventHandler.eventV2IMMessageReadV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMMessageReadV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMMessageReadV1)
		}
		return true, s, err
	case req.eventV2IMChatDisbandedV1 != nil:
		if r.cli.eventHandler.eventV2IMChatDisbandedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatDisbandedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatDisbandedV1)
		}
		return true, s, err
	case req.eventV2IMChatUpdatedV1 != nil:
		if r.cli.eventHandler.eventV2IMChatUpdatedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatUpdatedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatUpdatedV1)
		}
		return true, s, err
	case req.eventV2IMChatMemberBotAddedV1 != nil:
		if r.cli.eventHandler.eventV2IMChatMemberBotAddedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatMemberBotAddedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatMemberBotAddedV1)
		}
		return true, s, err
	case req.eventV2IMChatMemberBotDeletedV1 != nil:
		if r.cli.eventHandler.eventV2IMChatMemberBotDeletedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatMemberBotDeletedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatMemberBotDeletedV1)
		}
		return true, s, err
	case req.eventV2IMChatMemberUserAddedV1 != nil:
		if r.cli.eventHandler.eventV2IMChatMemberUserAddedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatMemberUserAddedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatMemberUserAddedV1)
		}
		return true, s, err
	case req.eventV2IMChatMemberUserWithdrawnV1 != nil:
		if r.cli.eventHandler.eventV2IMChatMemberUserWithdrawnV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatMemberUserWithdrawnV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatMemberUserWithdrawnV1)
		}
		return true, s, err
	case req.eventV2IMChatMemberUserDeletedV1 != nil:
		if r.cli.eventHandler.eventV2IMChatMemberUserDeletedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2IMChatMemberUserDeletedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2IMChatMemberUserDeletedV1)
		}
		return true, s, err
	case req.eventV2VCMeetingMeetingStartedV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingMeetingStartedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingMeetingStartedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingMeetingStartedV1)
		}
		return true, s, err
	case req.eventV2VCMeetingMeetingEndedV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingMeetingEndedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingMeetingEndedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingMeetingEndedV1)
		}
		return true, s, err
	case req.eventV2VCMeetingJoinMeetingV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingJoinMeetingV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingJoinMeetingV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingJoinMeetingV1)
		}
		return true, s, err
	case req.eventV2VCMeetingLeaveMeetingV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingLeaveMeetingV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingLeaveMeetingV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingLeaveMeetingV1)
		}
		return true, s, err
	case req.eventV2VCMeetingRecordingStartedV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingRecordingStartedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingRecordingStartedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingRecordingStartedV1)
		}
		return true, s, err
	case req.eventV2VCMeetingRecordingEndedV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingRecordingEndedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingRecordingEndedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingRecordingEndedV1)
		}
		return true, s, err
	case req.eventV2VCMeetingRecordingReadyV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingRecordingReadyV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingRecordingReadyV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingRecordingReadyV1)
		}
		return true, s, err
	case req.eventV2VCMeetingShareStartedV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingShareStartedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingShareStartedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingShareStartedV1)
		}
		return true, s, err
	case req.eventV2VCMeetingShareEndedV1 != nil:
		if r.cli.eventHandler.eventV2VCMeetingShareEndedV1Handler != nil {
			s, err = r.cli.eventHandler.eventV2VCMeetingShareEndedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventV2VCMeetingShareEndedV1)
		}
		return true, s, err
	case req.eventV1P2PChatCreate != nil:
		if r.cli.eventHandler.eventV1P2PChatCreateHandler != nil {
			s, err = r.cli.eventHandler.eventV1P2PChatCreateHandler(ctx, r.cli, req.Schema, req.headerV1(EventTypeV1P2PChatCreate), req.eventV1P2PChatCreate)
		}
		return true, s, err
	case req.eventV1AddUserToChat != nil:
		if r.cli.eventHandler.eventV1AddUserToChatHandler != nil {
			s, err = r.cli.eventHandler.eventV1AddUserToChatHandler(ctx, r.cli, req.Schema, req.headerV1(EventTypeV1AddUserToChat), req.eventV1AddUserToChat)
		}
		return true, s, err
	}
	return false, "", nil
}
