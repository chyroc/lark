package lark

import (
	"net/http"
	"time"
)

func New(appID, appSecret string) API {
	return &apiImpl{
		appID:     appID,
		appSecret: appSecret,
		httpClient: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

type API interface {
	Message() *MessageAPI
	Chat() *ChatAPI
}
