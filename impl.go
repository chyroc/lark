package lark

import (
	"net/http"
	"time"
)

type Lark struct {
	appID             string
	appSecret         string
	encryptKey        string
	verificationToken string
	helpdeskID        string
	helpdeskToken     string

	timeout    time.Duration
	httpClient *http.Client

	mock         *Mock
	eventHandler *eventHandler
}

func (r *Lark) Chat() *ChatAPI {
	return &ChatAPI{
		cli: r,
	}
}

type ChatAPI struct {
	cli *Lark
}

func (r *Lark) Message() *MessageAPI {
	return &MessageAPI{
		cli: r,
	}
}

type MessageAPI struct {
	cli *Lark
}

func (r *Lark) Bot() *BotAPI {
	return &BotAPI{
		cli: r,
	}
}

type BotAPI struct {
	cli *Lark
}

func (r *Lark) Calendar() *CalendarAPI {
	return &CalendarAPI{
		cli: r,
	}
}

type CalendarAPI struct {
	cli *Lark
}

func (r *Lark) Token() *TokenAPI {
	return &TokenAPI{
		cli: r,
	}
}

type TokenAPI struct {
	cli *Lark
}

func (r *Lark) Contact() *ContactAPI {
	return &ContactAPI{
		cli: r,
	}
}

type ContactAPI struct {
	cli *Lark
}

func (r *Lark) Approval() *ApprovalAPI {
	return &ApprovalAPI{
		cli: r,
	}
}

type ApprovalAPI struct {
	cli *Lark
}

func (r *Lark) Helpdesk() *HelpdeskAPI {
	return &HelpdeskAPI{
		cli: r,
	}
}

type HelpdeskAPI struct {
	cli *Lark
}

func (r *Lark) File() *FileAPI {
	return &FileAPI{
		cli: r,
	}
}

type FileAPI struct {
	cli *Lark
}

func (r *Lark) Attendance() *AttendanceAPI {
	return &AttendanceAPI{
		cli: r,
	}
}

type AttendanceAPI struct {
	cli *Lark
}

func (r *Lark) Mail() *MailAPI {
	return &MailAPI{
		cli: r,
	}
}

type MailAPI struct {
	cli *Lark
}

func (r *Lark) MeetingRoom() *MeetingRoomAPI {
	return &MeetingRoomAPI{
		cli: r,
	}
}

type MeetingRoomAPI struct {
	cli *Lark
}

func (r *Lark) EventCallback() *EventCallbackAPI {
	return &EventCallbackAPI{
		cli: r,
	}
}

type VCAPI struct {
	cli *Lark
}

func (r *Lark) VC() *VCAPI {
	return &VCAPI{
		cli: r,
	}
}

type TenantAPI struct {
	cli *Lark
}

func (r *Lark) Tenant() *TenantAPI {
	return &TenantAPI{
		cli: r,
	}
}

type EHRAPI struct {
	cli *Lark
}

func (r *Lark) EHR() *EHRAPI {
	return &EHRAPI{
		cli: r,
	}
}

type AIAPI struct {
	cli *Lark
}

func (r *Lark) AI() *AIAPI {
	return &AIAPI{
		cli: r,
	}
}

type HumanAuthAPI struct {
	cli *Lark
}

func (r *Lark) HumanAuth() *HumanAuthAPI {
	return &HumanAuthAPI{
		cli: r,
	}
}

type AdminAPI struct {
	cli *Lark
}

func (r *Lark) Admin() *AdminAPI {
	return &AdminAPI{
		cli: r,
	}
}

type EventCallbackAPI struct {
	cli *Lark
}
