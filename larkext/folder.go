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

func (r *Folder) Meta(ctx context.Context) (*lark.GetDriveFolderMetaResp, error) {
	return r.meta(ctx)
}

func (r *Folder) NewFolder(ctx context.Context, title string) (*Folder, error) {
	return r.newFolder(ctx, title)
}

func (r *Folder) NewSheet(ctx context.Context, title string) (*Sheet, error) {
	return r.newSheet(ctx, title)
}

func (r *Folder) NewDoc(ctx context.Context, title string) (*Doc, error) {
	return r.newDoc(ctx, title)
}
