package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

type Bitable struct {
	larkClient *lark.Lark
	appToken   string
}

func NewBitable(larkClient *lark.Lark, appToken string) *Bitable {
	return &Bitable{larkClient: larkClient, appToken: appToken}
}

func (r *Bitable) Meta(ctx context.Context) (*lark.GetBitableMetaRespApp, error) {
	return r.meta(ctx)
}
