package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

func (r *Bitable) meta(ctx context.Context) (*lark.GetBitableMetaRespApp, error) {
	resp, _, err := r.larkClient.Bitable.GetBitableMeta(ctx, &lark.GetBitableMetaReq{
		AppToken: r.appToken,
	})
	if err != nil {
		return nil, err
	}
	return resp.App, err
}
