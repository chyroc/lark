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

func WithTimeout(timeout time.Duration) ClientOptionFunc {
	return func(r *Lark) {
		r.timeout = timeout
	}
}

func WithLogger(logger Logger, level LogLevel) ClientOptionFunc {
	return func(lark *Lark) {
		lark.logger = logger
		lark.logLevel = level
	}
}

func WithISV(isISV bool) ClientOptionFunc {
	return func(lark *Lark) {
		lark.isISV = isISV
	}
}

func WithStore(store Store) ClientOptionFunc {
	return func(lark *Lark) {
		lark.store = store
	}
}

func New(options ...ClientOptionFunc) *Lark {
	return newClient("", options)
}

func newClient(tenantKey string, options []ClientOptionFunc) *Lark {
	r := &Lark{
		timeout:      time.Second * 3,
		isISV:        false,
		tenantKey:    tenantKey,
		store:        NewStoreMemory(),
		mock:         new(Mock),
		eventHandler: new(eventHandler),
	}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	r.httpClient = &http.Client{
		Timeout: r.timeout,
	}

	r.initService()

	return r
}
