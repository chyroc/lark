package lark

import (
	"context"
	"math"
)

type HandlerFunc func(*RequestContext)

const abortIndex int8 = math.MaxInt8 >> 1

type RequestContext struct {
	handlers []HandlerFunc
	index    int8

	Context      context.Context
	Request      *RawHttpRequest
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

func (c *RequestContext) Abort() {
	c.index = abortIndex
}

func newRequestContext(ctx context.Context, req *RawHttpRequest, realResp interface{}, handlerList ...HandlerFunc) *RequestContext {
	return &RequestContext{
		Request:      req,
		Context:      ctx,
		handlers:     handlerList,
		index:        -1,
		RealResponse: realResp,
	}
}
