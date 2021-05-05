package lark

import (
	"net/http"
	"time"
)

type ClientOptionFunc func(*Lark)

func WithAppCredential(appID, appSecret string) ClientOptionFunc {
	return func(r *Lark) {
		r.appID = appID
		r.appSecret = appSecret
	}
}

func WithEventCallbackVerify(encryptKey, verificationToken string) ClientOptionFunc {
	return func(r *Lark) {
		r.encryptKey = encryptKey
		r.verificationToken = verificationToken
	}
}

func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(r *Lark) {
		r.helpdeskID = helpdeskID
		r.helpdeskToken = helpdeskToken
	}
}

func New(options ...ClientOptionFunc) *Lark {
	r := new(Lark)
	r.mock = new(Mock)
	r.eventHandler = new(eventHandler)
	r.httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	return r
}
