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

func NewRootFolder(ctx context.Context, larkClient *lark.Lark) (*Folder, error) {
	f := NewFolder(larkClient, "")
	meta, err := f.meta(ctx)
	if err != nil {
		return nil, err
	}
	f.folderToken = meta.Token

	return f, nil
}

func (r *Folder) Meta(ctx context.Context) (*FolderMeta, error) {
	return r.meta(ctx)
}

// ListFiles 列出文件下的文件
func (r *Folder) ListFiles(ctx context.Context) (map[string]*lark.GetDriveFolderChildrenRespChildren, error) {
	if r.folderToken == "" {
		meta, err := r.meta(ctx)
		if err != nil {
			return nil, err
		}
		r.folderToken = meta.Token
	}
	return r.listFiles(ctx)
}

func (r *Folder) NewFolder(ctx context.Context, title string) (*Folder, error) {
	return r.newFolder(ctx, title)
}

// NewSheet 创建表格
func (r *Folder) NewSheet(ctx context.Context, title string) (*Sheet, error) {
	return r.newSheet(ctx, title)
}

func (r *Folder) NewDoc(ctx context.Context, title string) (*Doc, error) {
	return r.newDoc(ctx, title)
}

func (r *Folder) DeleteSheet(ctx context.Context, sheetToken string) error {
	return r.deleteSheet(ctx, sheetToken)
}

func (r *Folder) DeleteDoc(ctx context.Context, docToken string) error {
	return r.deleteDoc(ctx, docToken)
}

type FolderMeta struct {
	ID        string `json:"id,omitempty"`        // 文件夹的 id
	Name      string `json:"name,omitempty"`      // 文件夹的标题
	Token     string `json:"token,omitempty"`     // 文件夹的 token
	CreateUid string `json:"createUid,omitempty"` // 文件夹的创建者 id
	EditUid   string `json:"editUid,omitempty"`   // 文件夹的最后编辑者 id
	ParentID  string `json:"parentId,omitempty"`  // 文件夹的上级目录 id
	OwnUid    string `json:"ownUid,omitempty"`    // 文件夹为个人文件夹时，为文件夹的所有者 id；文件夹为共享文件夹时，为文件夹树id
}
