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

func (r *Mock) clone() *Mock {
	return &Mock{
		mockGetTenantAccessToken: r.mockGetTenantAccessToken,
	}
}
