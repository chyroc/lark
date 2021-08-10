package lark

import (
	"context"
	"io"
	"net/http"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b

func (r *EventCallbackService) ListenCallback(ctx context.Context, reader io.Reader, writer http.ResponseWriter) {
	r.listenCallback(ctx, false, http.Header{}, reader, writer)
}

func (r *EventCallbackService) ListenSecurityCallback(ctx context.Context, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	r.listenCallback(ctx, true, header, reader, writer)
}

// ListenCardCallback 卡片消息回调监听
func (r *EventCallbackService) ListenCardCallback(ctx context.Context, checkSecurity bool, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	r.listenCardCallback(ctx, checkSecurity, header, reader, writer)
}
