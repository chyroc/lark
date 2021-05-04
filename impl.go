package lark

import (
	"net/http"
)

type Lark struct {
	appID      string
	appSecret  string
	httpClient *http.Client
	mock       *Mock
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

func (r *Lark) HelpDesk() *HelpDeskAPI {
	return &HelpDeskAPI{
		cli: r,
	}
}

type HelpDeskAPI struct {
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
