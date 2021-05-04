package lark

import (
	"context"
)

type Mock struct {
	mockGetTenantAccessToken func(ctx context.Context) (*TokenExpire, *Response, error)
}

func (r *Lark) Mock() *Mock {
	return r.mock
}
