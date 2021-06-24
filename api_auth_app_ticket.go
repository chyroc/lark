package lark

import (
	"context"
	"time"
)

// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/g

func (r *Lark) WithTenant(tenantKey string) *Lark {
	r2 := r.clone()
	r2.tenantKey = tenantKey
	return r2
}

func (r *AuthService) GetAppTicket(ctx context.Context) (string, error) {
	s, _, err := r.cli.store.Get(ctx, genISVAppTicketKey(r.cli.appSecret))
	return s, err
}

func (r *AuthService) SetAppTicket(ctx context.Context, appTicket string) error {
	return r.cli.store.Set(ctx, genISVAppTicketKey(r.cli.appID), appTicket, time.Hour*2)
}
