package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

func (r *Drive) newFolder(ctx context.Context) {
	r.larkClient.Drive.CreateDriveFolder(ctx, &lark.CreateDriveFolderReq{
		// FolderToken: r.,
		// Title:       "",
	})
}
