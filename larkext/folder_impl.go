/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
		Name:        title,
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

// file：box开头云文档类型
// docx：docx文档类型
// bitable：多维表格类型
// folder：文件夹类型(新版云空间下可用)
// doc：doc文档类型
// sheet：电子表格类型
// mindnote：思维笔记类型
// shortcut：快捷方式类型(新版云空间下可用)
func (r *Folder) deleteFile(ctx context.Context, fileToken, typ string) error {
	_, _, err := r.larkClient.Drive.DeleteDriveFile(ctx, &lark.DeleteDriveFileReq{
		Type:      typ,
		FileToken: fileToken,
	})
	return err
}
