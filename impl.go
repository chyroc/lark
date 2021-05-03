package lark

import (
	"net/http"
)

type Lark struct {
	appID      string
	appSecret  string
	httpClient *http.Client
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
