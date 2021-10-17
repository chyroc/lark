package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

type Folder struct {
	larkClient  *lark.Lark
	folderToken string
}

func NewFolder(larkClient *lark.Lark, folderToken string) *Folder {
	r := new(Folder)
	r.larkClient = larkClient
	r.folderToken = folderToken

	return r
}

func (r *Folder) NewSheet(ctx context.Context, title string) (*Sheet, error) {
	return r.newSheet(ctx, title)
}
