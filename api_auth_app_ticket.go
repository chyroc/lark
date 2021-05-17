package lark

import (
	"context"
	"time"
)

// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/g

func (r *Lark) WithTenant(tenantKey string) *Lark {
	return &Lark{
		appID:             r.appID,
		appSecret:         r.appSecret,
		encryptKey:        r.encryptKey,
		verificationToken: r.verificationToken,
		helpdeskID:        r.helpdeskID,
		helpdeskToken:     r.helpdeskToken,
		timeout:           r.timeout,
		isISV:             r.isISV,
		tenantKey:         tenantKey,
		httpClient:        r.httpClient,
		logger:            r.logger,
		logLevel:          r.logLevel,
		store:             r.store,
		mock:              r.mock,
		eventHandler:      r.eventHandler,
	}
}

func (r *AuthService) GetAppTicket(ctx context.Context) (string, error) {
	s, _, err := r.cli.store.Get(ctx, genISVAppTicketKey(r.cli.appSecret))
	return s, err
}

func (r *AuthService) SetAppTicket(ctx context.Context, appTicket string) error {
	return r.cli.store.Set(ctx, genISVAppTicketKey(r.cli.appID), appTicket, time.Hour*2)
}
