package lark

import (
	"net/http"
)

type apiImpl struct {
	appID      string
	appSecret  string
	httpClient *http.Client
}

func (r *apiImpl) Chat() *ChatAPI {
	return &ChatAPI{
		cli: r,
	}
}

type ChatAPI struct {
	cli *apiImpl
}

func (r *apiImpl) Message() *MessageAPI {
	return &MessageAPI{
		cli: r,
	}
}

type MessageAPI struct {
	cli *apiImpl
}

func (r *apiImpl) Token() *TokenAPI {
	return &TokenAPI{
		cli: r,
	}
}

type TokenAPI struct {
	cli *apiImpl
}

func (r *apiImpl) Contact() *ContactAPI {
	return &ContactAPI{
		cli: r,
	}
}

type ContactAPI struct {
	cli *apiImpl
}
