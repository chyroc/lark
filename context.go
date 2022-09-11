package lark

import (
	"context"
)

type HandlerFunc func(*RequestContext)

type RequestContext struct {
	handlers []HandlerFunc
	index    int8

	Context      context.Context
	Request      *rawHttpRequest
	RealResponse interface{}
	Resp         *Response
	Err          error
}

func (c *RequestContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func newRequestContext(ctx context.Context, req *rawHttpRequest, realResp interface{}, handlerList ...HandlerFunc) *RequestContext {
	return &RequestContext{
		Request:      req,
		Context:      ctx,
		handlers:     handlerList,
		index:        -1,
		RealResponse: realResp,
	}
}
