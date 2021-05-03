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

func New(options ...ClientOptionFunc) *Lark {
	r := new(Lark)
	r.httpClient = &http.Client{
		Timeout: time.Second * 3,
	}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	return r
}
