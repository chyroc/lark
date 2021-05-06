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

	tenantKey string

	timeout    time.Duration
	httpClient *http.Client

	mock         *Mock
	eventHandler *eventHandler
}

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

func New(options ...ClientOptionFunc) *Lark {
	r := new(Lark)
	r.mock = new(Mock)
	r.eventHandler = new(eventHandler)
	if r.timeout == 0 {
		r.timeout = time.Second * 3
	}
	r.httpClient = &http.Client{
		Timeout: r.timeout,
	}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	return r
}

// 每个回调都创建一个，没有复用，搞个 sync.Pool ?
// 让用户自己去维护 tenantCli ？
// 搞个 map，存 tenant-key，然后 readme 说明让用户可以用内置的，也可以选择自己维护（也就是不用 .UseTenant 模式）
// 如果是每个 租户一个 cli，如何共用一个 cli  的 app-ticket ？
// 所以本质，必须一个 cli？
// clone 的模式 和 app-ticket 只有一个矛盾了。。。。
// tenant 模式，只有 tenant-key 不同，lark 结构体是复制过来的指针，保持全局同步？
// 那如何做不同tenant下的不同mock？
// 那如何做不同tenant下的不同 event handler？
//
// 换个思路，做 func-opt
// lark 实例搞个 mao，存 tenant-key，和各自 tenant-token
// 所以 app-ticket 天然共存
// 不同的 mock ？tenant-key 是函数参数，天然在 mock 函数的签名中
// 不同的 event-handler？
func (r *Lark) WithTenant(tenantKey string) *Lark {
	cli := r.clone()
	cli.tenantKey = tenantKey
	return cli
}

// 接口不是很多，用 func-opt，
func (r *Lark) WithUser(tenantKey string) *Lark {
	cli := r.clone()
	cli.tenantKey = tenantKey
	return cli
}

func (r *Lark) clone() *Lark {
	cli := &Lark{
		appID:             r.appID,
		appSecret:         r.appSecret,
		encryptKey:        r.encryptKey,
		verificationToken: r.verificationToken,
		helpdeskID:        r.helpdeskID,
		helpdeskToken:     r.helpdeskToken,
		timeout:           r.timeout,
		httpClient:        &http.Client{Timeout: r.timeout},
		mock:              r.mock.clone(),
		eventHandler:      r.eventHandler.clone(),
	}
	return cli
}

type MethodOptionFunc func(*Lark)

func WithTenant(tenantKey string) MethodOptionFunc {
	return func(r *Lark) {
	}
}

func WithUser() MethodOptionFunc {
	return func(r *Lark) {
	}
}
