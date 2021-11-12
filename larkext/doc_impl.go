package larkext

import (
	"context"
	"encoding/json"

	"github.com/chyroc/lark"
)

func (r *Doc) meta(ctx context.Context) (*lark.GetDriveDocMetaResp, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocMeta(ctx, &lark.GetDriveDocMetaReq{
		DocToken: r.docToken,
	})
	return resp, err
}

func (r *Doc) delete(ctx context.Context) error {
	_, _, err := r.larkClient.Drive.DeleteDriveDocFile(ctx, &lark.DeleteDriveDocFileReq{
		DocToken: r.docToken,
	})
	return err
}

func (r *Doc) rawContent(ctx context.Context) (string, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocRawContent(ctx, &lark.GetDriveDocRawContentReq{
		DocToken: r.docToken,
	})
	return resp.Content, err
}

func (r *Doc) content(ctx context.Context) (*lark.DocContent, error) {
	resp, _, err := r.larkClient.Drive.GetDriveDocContent(ctx, &lark.GetDriveDocContentReq{
		DocToken: r.docToken,
	})
	doc := &lark.DocContent{}
	err = json.Unmarshal([]byte(resp.Content), doc)
	if err != nil {
		return nil, err
	}
	return doc, err
}

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

// 编辑
