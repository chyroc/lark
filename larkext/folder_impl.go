package larkext

import (
	"context"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
)

func (r *Folder) meta(ctx context.Context) (*lark.GetDriveFolderMetaResp, error) {
	resp, _, err := r.larkClient.Drive.GetDriveFolderMeta(ctx, &lark.GetDriveFolderMetaReq{
		FolderToken: r.folderToken,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Folder) newFolder(ctx context.Context, title string) (*Folder, error) {
	resp, _, err := r.larkClient.Drive.CreateDriveFolder(ctx, &lark.CreateDriveFolderReq{
		FolderToken: r.folderToken,
		Title:       title,
	})
	if err != nil {
		return nil, err
	}
	return NewFolder(r.larkClient, resp.Token), nil
}

func (r *Folder) newSheet(ctx context.Context, title string) (*Sheet, error) {
	resp, _, err := r.larkClient.Drive.CreateSheet(ctx, &lark.CreateSheetReq{
		Title:       &title,
		FolderToken: ptr.StringNoNonePtr(r.folderToken),
	})
	if err != nil {
		return nil, err
	}
	return NewSheet(r.larkClient, resp.Spreadsheet.SpreadSheetToken), nil
}

func (r *Folder) newDoc(ctx context.Context, title string) (*Doc, error) {
	resp, _, err := r.larkClient.Drive.CreateDriveDoc(ctx, &lark.CreateDriveDocReq{
		Content:     ptr.StringNoNonePtr(""), // TODO:
		FolderToken: ptr.StringNoNonePtr(r.folderToken),
	})
	if err != nil {
		return nil, err
	}
	return NewDoc(r.larkClient, resp.ObjToken), nil
}
