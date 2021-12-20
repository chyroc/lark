package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

// Bitable Bitable client
type Bitable struct {
	larkClient *lark.Lark
	appToken   string
}

// NewBitable new Bitable client
func NewBitable(larkClient *lark.Lark, appToken string) *Bitable {
	return &Bitable{larkClient: larkClient, appToken: appToken}
}

// Meta get bitable meta
func (r *Bitable) Meta(ctx context.Context) (*lark.GetBitableMetaRespApp, error) {
	return r.meta(ctx)
}
