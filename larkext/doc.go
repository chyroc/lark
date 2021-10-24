package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

type Doc struct {
	larkClient *lark.Lark
	docToken   string
}

func NewDoc(larkClient *lark.Lark, docToken string) *Doc {
	r := new(Doc)
	r.larkClient = larkClient
	r.docToken = docToken
	return r
}

// Meta 获取文档元信息
func (r *Doc) DocToken() string {
	return r.docToken
}

// Meta 获取文档元信息
func (r *Doc) Meta(ctx context.Context) (*lark.GetDriveDocMetaResp, error) {
	return r.meta(ctx)
}

// Delete 删除云文档
func (r *Doc) Delete(ctx context.Context) error {
	return r.delete(ctx)
}

// RawContent 获取文档文本内容
func (r *Doc) RawContent(ctx context.Context) (string, error) {
	return r.rawContent(ctx)
}

// RawContent 获取文档文本内容
func (r *Doc) Content(ctx context.Context) (*lark.DocContent, error) {
	return r.content(ctx)
}
