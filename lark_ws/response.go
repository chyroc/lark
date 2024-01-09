package lark_ws

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type httpResponse struct {
	c      *Client
	ctx    context.Context
	frame  Frame
	start  int64
	status int
}

func (c *Client) newHTTPResponse(ctx context.Context, frame Frame, start int64) *httpResponse {
	return &httpResponse{
		c:      c,
		ctx:    ctx,
		frame:  frame,
		start:  start,
		status: http.StatusOK,
	}
}

func (r *httpResponse) Header() http.Header {
	return http.Header{}
}

func (r *httpResponse) Write(bytes []byte) (int, error) {
	return r.writeToWebsocket(bytes)
}

func (r *httpResponse) WriteHeader(statusCode int) {
	r.status = statusCode
}

func (r *httpResponse) writeToWebsocket(bs []byte) (int, error) {
	hs := wsHeader(r.frame.Headers)
	msgID := hs.GetString("message_id")
	traceID := hs.GetString("trace_id")
	type_ := hs.GetString("type")

	frame := r.frame
	end := time.Now().UnixMilli()
	hs.Add("biz_rt", strconv.FormatInt(end-r.start, 10)) // 业务处理耗时, ms

	resp := &responseMessage{
		StatusCode: r.status,
		Data:       bs,
	}
	p, _ := json.Marshal(resp)
	frame.Payload = p
	frame.Headers = hs
	bs, _ = frame.Marshal()
	err := r.c.writeMessage(websocket.BinaryMessage, bs)
	if err != nil {
		r.c.logError(r.ctx, "response message failed, type: %s, message_id: %s, trace_id: %s, err: %v", type_, msgID, traceID, err)
		return 0, err
	}
	return len(bs), nil
}
