package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

func (r *Doc) uploadMedia(ctx context.Context, parentType, parentToken string) {
	r.larkClient.Drive.UploadDriveMedia(ctx, &lark.UploadDriveMediaReq{
		FileName:   "",
		ParentType: parentType,
		ParentNode: parentToken,
		Size:       0,
		Checksum:   nil,
		Extra:      nil,
		File:       nil,
	})
}
