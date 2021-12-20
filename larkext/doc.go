package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

// Doc is doc client
type Doc struct {
	larkClient *lark.Lark
	docToken   string
}

// NewDoc new doc client
func NewDoc(larkClient *lark.Lark, docToken string) *Doc {
	r := new(Doc)
	r.larkClient = larkClient
	r.docToken = docToken
	return r
}

// DocToken get doc Token
func (r *Doc) DocToken() string {
	return r.docToken
}

// Meta get doc meta
func (r *Doc) Meta(ctx context.Context) (*lark.GetDriveDocMetaResp, error) {
	return r.meta(ctx)
}

// Delete delete doc
func (r *Doc) Delete(ctx context.Context) error {
	return r.delete(ctx)
}

// RawContent get doc raw content
func (r *Doc) RawContent(ctx context.Context) (string, error) {
	return r.rawContent(ctx)
}

// RawContent get doc rich content
func (r *Doc) Content(ctx context.Context) (*lark.DocContent, error) {
	return r.content(ctx)
}
