package larkext

import (
	"context"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
)

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
