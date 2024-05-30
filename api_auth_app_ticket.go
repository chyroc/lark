/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
	if err != nil && err != ErrStoreNotFound {
		r.cli.Log(ctx, LogLevelError, "[lark] Auth#GetAppTicket get_ticket_cache failed, app_id=%s, err=%s", r.cli.appID, err)
	}
	if err == ErrStoreNotFound && r.cli.getAppTicketFunc != nil {
		ticket, err := r.cli.getAppTicketFunc(ctx, r.cli, r.cli.appID)
		if err != nil {
			r.cli.Log(ctx, LogLevelError, "[lark] Auth#GetAppTicket get_ticket failed, app_id=%s, err=%s", r.cli.appID, err)
			return "", err
		}
		if err = r.SetAppTicket(ctx, ticket); err != nil {
			r.cli.Log(ctx, LogLevelError, "[lark] Auth#GetAppTicket set_ticket_cache failed, app_id=%s, err=%s", r.cli.appID, err)
		}
		return ticket, nil
	}
	return s, err
}

// SetAppTicket ...
func (r *AuthService) SetAppTicket(ctx context.Context, appTicket string) error {
	return r.cli.store.Set(ctx, genISVAppTicketKey(r.cli.appID), appTicket, time.Hour*2)
}
