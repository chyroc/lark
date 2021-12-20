package lark

import (
	"context"
	"time"
)

// WithTenant ...
func (r *Lark) WithTenant(tenantKey string) *Lark {
	return r.clone(tenantKey)
}

// GetAppTicket ...
func (r *AuthService) GetAppTicket(ctx context.Context) (string, error) {
	s, _, err := r.cli.store.Get(ctx, genISVAppTicketKey(r.cli.appID))
	return s, err
}

// SetAppTicket ...
func (r *AuthService) SetAppTicket(ctx context.Context, appTicket string) error {
	return r.cli.store.Set(ctx, genISVAppTicketKey(r.cli.appID), appTicket, time.Hour*2)
}
