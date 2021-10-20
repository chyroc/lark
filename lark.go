package lark

import (
	"net/http"
	"strings"
	"time"
)

func New(options ...ClientOptionFunc) *Lark {
	return newClient("", options)
}

var AppLink = New().AppLink

type ClientOptionFunc func(*Lark)

func WithAppCredential(appID, appSecret string) ClientOptionFunc {
	return func(r *Lark) {
		r.appID = appID
		r.appSecret = appSecret
	}
}

func WithCustomBot(customBotWebHookURL, customBotSecret string) ClientOptionFunc {
	return func(r *Lark) {
		r.customBotWebHookURL = customBotWebHookURL
		r.customBotSecret = customBotSecret
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

func WithOpenBaseURL(baseURL string) ClientOptionFunc {
	return func(lark *Lark) {
		lark.openBaseURL = strings.TrimRight(baseURL, "/")
	}
}

func WithWWWBaseURL(baseURL string) ClientOptionFunc {
	return func(lark *Lark) {
		lark.wwwBaseURL = strings.TrimRight(baseURL, "/")
	}
}

type MethodOptionFunc func(*MethodOption)

func WithUserAccessToken(token string) MethodOptionFunc {
	return func(option *MethodOption) {
		option.userAccessToken = token
	}
}

type MethodOption struct {
	userAccessToken string
}

func newMethodOption(options []MethodOptionFunc) *MethodOption {
	opt := new(MethodOption)
	for _, v := range options {
		v(opt)
	}
	return opt
}

func newClient(tenantKey string, options []ClientOptionFunc) *Lark {
	r := &Lark{
		timeout:      time.Second * 3,
		isISV:        false,
		tenantKey:    tenantKey,
		store:        NewStoreMemory(),
		mock:         new(Mock),
		eventHandler: new(eventHandler),
		openBaseURL:  "https://open.feishu.cn",
		wwwBaseURL:   "https://www.feishu.cn",
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
