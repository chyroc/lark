package lark

import (
	"context"
	"fmt"
)

type EventType string

const (
	EventTypeIMMessageReceiveV1          EventType = "im.message.receive_v1"
	EventTypeIMMessageReadV1             EventType = "im.message.message_read_v1"
	EventTypeIMChatDisbandedV1           EventType = "im.chat.disbanded_v1"
	EventTypeIMChatUpdatedV1             EventType = "im.chat.updated_v1"
	EventTypeIMChatMemberBotAddedV1      EventType = "im.chat.member.bot.added_v1"
	EventTypeIMChatMemberBotDeletedV1    EventType = "im.chat.member.bot.deleted_v1"
	EventTypeIMChatMemberUserAddedV1     EventType = "im.chat.member.user.added_v1"
	EventTypeIMChatMemberUserWithdrawnV1 EventType = "im.chat.member.user.withdrawn_v1"
	EventTypeIMChatMemberUserDeletedV1   EventType = "im.chat.member.user.deleted_v1"
	EventTypeVCMeetingMeetingStartedV1   EventType = "vc.meeting.meeting_started_v1"
	EventTypeVCMeetingMeetingEndedV1     EventType = "vc.meeting.meeting_ended_v1"
	EventTypeVCMeetingJoinMeetingV1      EventType = "vc.meeting.join_meeting_v1"
	EventTypeVCMeetingLeaveMeetingV1     EventType = "vc.meeting.leave_meeting_v1"
	EventTypeVCMeetingRecordingStartedV1 EventType = "vc.meeting.recording_started_v1"
	EventTypeVCMeetingRecordingEndedV1   EventType = "vc.meeting.recording_ended_v1"
	EventTypeVCMeetingRecordingReadyV1   EventType = "vc.meeting.recording_ready_v1"
	EventTypeVCMeetingShareStartedV1     EventType = "vc.meeting.share_started_v1"
	EventTypeVCMeetingShareEndedV1       EventType = "vc.meeting.share_ended_v1"
)

type eventHandler struct {
	eventIMMessageReceiveV1Handler          eventIMMessageReceiveV1Handler
	eventIMMessageReadV1Handler             eventIMMessageReadV1Handler
	eventIMChatDisbandedV1Handler           eventIMChatDisbandedV1Handler
	eventIMChatUpdatedV1Handler             eventIMChatUpdatedV1Handler
	eventIMChatMemberBotAddedV1Handler      eventIMChatMemberBotAddedV1Handler
	eventIMChatMemberBotDeletedV1Handler    eventIMChatMemberBotDeletedV1Handler
	eventIMChatMemberUserAddedV1Handler     eventIMChatMemberUserAddedV1Handler
	eventIMChatMemberUserWithdrawnV1Handler eventIMChatMemberUserWithdrawnV1Handler
	eventIMChatMemberUserDeletedV1Handler   eventIMChatMemberUserDeletedV1Handler
	eventVCMeetingMeetingStartedV1Handler   eventVCMeetingMeetingStartedV1Handler
	eventVCMeetingMeetingEndedV1Handler     eventVCMeetingMeetingEndedV1Handler
	eventVCMeetingJoinMeetingV1Handler      eventVCMeetingJoinMeetingV1Handler
	eventVCMeetingLeaveMeetingV1Handler     eventVCMeetingLeaveMeetingV1Handler
	eventVCMeetingRecordingStartedV1Handler eventVCMeetingRecordingStartedV1Handler
	eventVCMeetingRecordingEndedV1Handler   eventVCMeetingRecordingEndedV1Handler
	eventVCMeetingRecordingReadyV1Handler   eventVCMeetingRecordingReadyV1Handler
	eventVCMeetingShareStartedV1Handler     eventVCMeetingShareStartedV1Handler
	eventVCMeetingShareEndedV1Handler       eventVCMeetingShareEndedV1Handler
}

func (r *eventHandler) clone() *eventHandler {
	return &eventHandler{
		eventIMMessageReceiveV1Handler:          r.eventIMMessageReceiveV1Handler,
		eventIMMessageReadV1Handler:             r.eventIMMessageReadV1Handler,
		eventIMChatDisbandedV1Handler:           r.eventIMChatDisbandedV1Handler,
		eventIMChatUpdatedV1Handler:             r.eventIMChatUpdatedV1Handler,
		eventIMChatMemberBotAddedV1Handler:      r.eventIMChatMemberBotAddedV1Handler,
		eventIMChatMemberBotDeletedV1Handler:    r.eventIMChatMemberBotDeletedV1Handler,
		eventIMChatMemberUserAddedV1Handler:     r.eventIMChatMemberUserAddedV1Handler,
		eventIMChatMemberUserWithdrawnV1Handler: r.eventIMChatMemberUserWithdrawnV1Handler,
		eventIMChatMemberUserDeletedV1Handler:   r.eventIMChatMemberUserDeletedV1Handler,
		eventVCMeetingMeetingStartedV1Handler:   r.eventVCMeetingMeetingStartedV1Handler,
		eventVCMeetingMeetingEndedV1Handler:     r.eventVCMeetingMeetingEndedV1Handler,
		eventVCMeetingJoinMeetingV1Handler:      r.eventVCMeetingJoinMeetingV1Handler,
		eventVCMeetingLeaveMeetingV1Handler:     r.eventVCMeetingLeaveMeetingV1Handler,
		eventVCMeetingRecordingStartedV1Handler: r.eventVCMeetingRecordingStartedV1Handler,
		eventVCMeetingRecordingEndedV1Handler:   r.eventVCMeetingRecordingEndedV1Handler,
		eventVCMeetingRecordingReadyV1Handler:   r.eventVCMeetingRecordingReadyV1Handler,
		eventVCMeetingShareStartedV1Handler:     r.eventVCMeetingShareStartedV1Handler,
		eventVCMeetingShareEndedV1Handler:       r.eventVCMeetingShareEndedV1Handler,
	}
}

type eventBody struct {
	eventIMMessageReceiveV1          *EventIMMessageReceiveV1
	eventIMMessageReadV1             *EventIMMessageReadV1
	eventIMChatDisbandedV1           *EventIMChatDisbandedV1
	eventIMChatUpdatedV1             *EventIMChatUpdatedV1
	eventIMChatMemberBotAddedV1      *EventIMChatMemberBotAddedV1
	eventIMChatMemberBotDeletedV1    *EventIMChatMemberBotDeletedV1
	eventIMChatMemberUserAddedV1     *EventIMChatMemberUserAddedV1
	eventIMChatMemberUserWithdrawnV1 *EventIMChatMemberUserWithdrawnV1
	eventIMChatMemberUserDeletedV1   *EventIMChatMemberUserDeletedV1
	eventVCMeetingMeetingStartedV1   *EventVCMeetingMeetingStartedV1
	eventVCMeetingMeetingEndedV1     *EventVCMeetingMeetingEndedV1
	eventVCMeetingJoinMeetingV1      *EventVCMeetingJoinMeetingV1
	eventVCMeetingLeaveMeetingV1     *EventVCMeetingLeaveMeetingV1
	eventVCMeetingRecordingStartedV1 *EventVCMeetingRecordingStartedV1
	eventVCMeetingRecordingEndedV1   *EventVCMeetingRecordingEndedV1
	eventVCMeetingRecordingReadyV1   *EventVCMeetingRecordingReadyV1
	eventVCMeetingShareStartedV1     *EventVCMeetingShareStartedV1
	eventVCMeetingShareEndedV1       *EventVCMeetingShareEndedV1
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
	case EventTypeIMChatDisbandedV1:
		event := new(EventIMChatDisbandedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatDisbandedV1 = event
	case EventTypeIMChatUpdatedV1:
		event := new(EventIMChatUpdatedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatUpdatedV1 = event
	case EventTypeIMChatMemberBotAddedV1:
		event := new(EventIMChatMemberBotAddedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatMemberBotAddedV1 = event
	case EventTypeIMChatMemberBotDeletedV1:
		event := new(EventIMChatMemberBotDeletedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatMemberBotDeletedV1 = event
	case EventTypeIMChatMemberUserAddedV1:
		event := new(EventIMChatMemberUserAddedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatMemberUserAddedV1 = event
	case EventTypeIMChatMemberUserWithdrawnV1:
		event := new(EventIMChatMemberUserWithdrawnV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatMemberUserWithdrawnV1 = event
	case EventTypeIMChatMemberUserDeletedV1:
		event := new(EventIMChatMemberUserDeletedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventIMChatMemberUserDeletedV1 = event
	case EventTypeVCMeetingMeetingStartedV1:
		event := new(EventVCMeetingMeetingStartedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingMeetingStartedV1 = event
	case EventTypeVCMeetingMeetingEndedV1:
		event := new(EventVCMeetingMeetingEndedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingMeetingEndedV1 = event
	case EventTypeVCMeetingJoinMeetingV1:
		event := new(EventVCMeetingJoinMeetingV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingJoinMeetingV1 = event
	case EventTypeVCMeetingLeaveMeetingV1:
		event := new(EventVCMeetingLeaveMeetingV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingLeaveMeetingV1 = event
	case EventTypeVCMeetingRecordingStartedV1:
		event := new(EventVCMeetingRecordingStartedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingRecordingStartedV1 = event
	case EventTypeVCMeetingRecordingEndedV1:
		event := new(EventVCMeetingRecordingEndedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingRecordingEndedV1 = event
	case EventTypeVCMeetingRecordingReadyV1:
		event := new(EventVCMeetingRecordingReadyV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingRecordingReadyV1 = event
	case EventTypeVCMeetingShareStartedV1:
		event := new(EventVCMeetingShareStartedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingShareStartedV1 = event
	case EventTypeVCMeetingShareEndedV1:
		event := new(EventVCMeetingShareEndedV1)
		if err := req.unmarshalEvent(event); err != nil {
			return err
		}
		req.eventVCMeetingShareEndedV1 = event
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
	case req.eventIMChatDisbandedV1 != nil:
		if r.cli.eventHandler.eventIMChatDisbandedV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatDisbandedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatDisbandedV1)
		}
		return true, s, err
	case req.eventIMChatUpdatedV1 != nil:
		if r.cli.eventHandler.eventIMChatUpdatedV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatUpdatedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatUpdatedV1)
		}
		return true, s, err
	case req.eventIMChatMemberBotAddedV1 != nil:
		if r.cli.eventHandler.eventIMChatMemberBotAddedV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatMemberBotAddedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatMemberBotAddedV1)
		}
		return true, s, err
	case req.eventIMChatMemberBotDeletedV1 != nil:
		if r.cli.eventHandler.eventIMChatMemberBotDeletedV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatMemberBotDeletedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatMemberBotDeletedV1)
		}
		return true, s, err
	case req.eventIMChatMemberUserAddedV1 != nil:
		if r.cli.eventHandler.eventIMChatMemberUserAddedV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatMemberUserAddedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatMemberUserAddedV1)
		}
		return true, s, err
	case req.eventIMChatMemberUserWithdrawnV1 != nil:
		if r.cli.eventHandler.eventIMChatMemberUserWithdrawnV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatMemberUserWithdrawnV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatMemberUserWithdrawnV1)
		}
		return true, s, err
	case req.eventIMChatMemberUserDeletedV1 != nil:
		if r.cli.eventHandler.eventIMChatMemberUserDeletedV1Handler != nil {
			s, err = r.cli.eventHandler.eventIMChatMemberUserDeletedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventIMChatMemberUserDeletedV1)
		}
		return true, s, err
	case req.eventVCMeetingMeetingStartedV1 != nil:
		if r.cli.eventHandler.eventVCMeetingMeetingStartedV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingMeetingStartedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingMeetingStartedV1)
		}
		return true, s, err
	case req.eventVCMeetingMeetingEndedV1 != nil:
		if r.cli.eventHandler.eventVCMeetingMeetingEndedV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingMeetingEndedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingMeetingEndedV1)
		}
		return true, s, err
	case req.eventVCMeetingJoinMeetingV1 != nil:
		if r.cli.eventHandler.eventVCMeetingJoinMeetingV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingJoinMeetingV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingJoinMeetingV1)
		}
		return true, s, err
	case req.eventVCMeetingLeaveMeetingV1 != nil:
		if r.cli.eventHandler.eventVCMeetingLeaveMeetingV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingLeaveMeetingV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingLeaveMeetingV1)
		}
		return true, s, err
	case req.eventVCMeetingRecordingStartedV1 != nil:
		if r.cli.eventHandler.eventVCMeetingRecordingStartedV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingRecordingStartedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingRecordingStartedV1)
		}
		return true, s, err
	case req.eventVCMeetingRecordingEndedV1 != nil:
		if r.cli.eventHandler.eventVCMeetingRecordingEndedV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingRecordingEndedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingRecordingEndedV1)
		}
		return true, s, err
	case req.eventVCMeetingRecordingReadyV1 != nil:
		if r.cli.eventHandler.eventVCMeetingRecordingReadyV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingRecordingReadyV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingRecordingReadyV1)
		}
		return true, s, err
	case req.eventVCMeetingShareStartedV1 != nil:
		if r.cli.eventHandler.eventVCMeetingShareStartedV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingShareStartedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingShareStartedV1)
		}
		return true, s, err
	case req.eventVCMeetingShareEndedV1 != nil:
		if r.cli.eventHandler.eventVCMeetingShareEndedV1Handler != nil {
			s, err = r.cli.eventHandler.eventVCMeetingShareEndedV1Handler(ctx, r.cli, req.Schema, req.Header, req.eventVCMeetingShareEndedV1)
		}
		return true, s, err
	}
	return false, "", nil
}
