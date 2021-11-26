package lark

import (
	"context"
	"time"
)

// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/g

func (r *Lark) WithTenant(tenantKey string) *Lark {
	return r.clone(tenantKey)
}

func (r *AuthService) GetAppTicket(ctx context.Context) (string, error) {
	s, _, err := r.cli.store.Get(ctx, genISVAppTicketKey(r.cli.appID))
	return s, err
}

func (r *AuthService) SetAppTicket(ctx context.Context, appTicket string) error {
	return r.cli.store.Set(ctx, genISVAppTicketKey(r.cli.appID), appTicket, time.Hour*2)
}
