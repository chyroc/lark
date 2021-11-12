package larkext

import (
	"context"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
)

func (r *Folder) meta(ctx context.Context) (*FolderMeta, error) {
	if r.folderToken == "" {
		resp, _, err := r.larkClient.Drive.GetDriveRootFolderMeta(ctx, &lark.GetDriveRootFolderMetaReq{})
		if err != nil {
			return nil, err
		}
		return &FolderMeta{
			ID:     resp.ID,
			Token:  resp.Token,
			OwnUid: resp.UserID,
		}, nil
	}
	resp, _, err := r.larkClient.Drive.GetDriveFolderMeta(ctx, &lark.GetDriveFolderMetaReq{
		FolderToken: r.folderToken,
	})
	if err != nil {
		return nil, err
	}
	return &FolderMeta{
		ID:        resp.ID,
		Name:      resp.Name,
		Token:     resp.Token,
		CreateUid: resp.CreateUid,
		EditUid:   resp.EditUid,
		ParentID:  resp.ParentID,
		OwnUid:    resp.OwnUid,
	}, nil
}

func (r *Folder) listFiles(ctx context.Context) (map[string]*lark.GetDriveFolderChildrenRespChildren, error) {
	resp, _, err := r.larkClient.Drive.GetDriveFolderChildren(ctx, &lark.GetDriveFolderChildrenReq{
		FolderToken: r.folderToken,
	})
	if err != nil {
		return nil, err
	}
	return resp.Children, nil
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

func (r *Folder) deleteSheet(ctx context.Context, sheetToken string) error {
	_, _, err := r.larkClient.Drive.DeleteDriveSheetFile(ctx, &lark.DeleteDriveSheetFileReq{
		SpreadSheetToken: sheetToken,
	})
	return err
}

func (r *Folder) deleteDoc(ctx context.Context, docToken string) error {
	_, _, err := r.larkClient.Drive.DeleteDriveDocFile(ctx, &lark.DeleteDriveDocFileReq{
		DocToken: docToken,
	})
	return err
}
