package lark

import (
	"context"
)

type apiImpl struct{}

type req struct {
	method string
	url    string
	body   interface{}
}

func (r *apiImpl) request(ctx context.Context, req req, resp interface{}) error {
	return nil
}
