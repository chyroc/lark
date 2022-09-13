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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

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
	resp, _, err := r.larkClient.Drive.CreateSpreadsheet(ctx, &lark.CreateSpreadsheetReq{
		Title:       &title,
		FolderToken: ptr.StringNoNonePtr(r.folderToken),
	})
	if err != nil {
		return nil, err
	}
	return NewSheet(r.larkClient, resp.Spreadsheet.SpreadSheetToken), nil
}

func (r *Folder) newDoc(ctx context.Context, title string, blocks ...*lark.DocBlock) (*Doc, error) {
	b := lark.DocContent{Title: &lark.DocParagraph{Elements: []*lark.DocParagraphElement{{Type: lark.DocParagraphElementTypeTextRun, TextRun: &lark.DocTextRun{Text: title}}}}}
	if blocks != nil {
		b.Body = &lark.DocBody{Blocks: blocks}
	}
	body, _ := json.Marshal(b)
	resp, _, err := r.larkClient.Drive.CreateDriveDoc(ctx, &lark.CreateDriveDocReq{
		Content:     ptr.StringNoNonePtr(string(body)),
		FolderToken: ptr.StringNoNonePtr(r.folderToken),
	})
	if err != nil {
		return nil, err
	}
	return newDoc(r.larkClient, resp.ObjToken, resp.URL), nil
}

func (r *Folder) newFile(ctx context.Context, title string, typ string) (*lark.CreateDriveFileResp, error) {
	resp, _, err := r.larkClient.Drive.CreateDriveFile(ctx, &lark.CreateDriveFileReq{
		FolderToken: r.folderToken,
		Title:       title,
		Type:        typ,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Folder) uploadFile(ctx context.Context, file *FileInfo) (*File, error) {
	var token string
	if file.Size < 20*1024 {
		resp, _, err := r.larkClient.Drive.UploadDriveFile(ctx, &lark.UploadDriveFileReq{
			ParentType: "explorer",
			ParentNode: r.folderToken,
			FileName:   file.FileName,
			Size:       file.Size,
			Checksum:   file.Checksum,
			File:       file.File,
		})
		if err != nil {
			return nil, err
		}
		token = resp.FileToken
	} else {
		prepareResp, _, err := r.larkClient.Drive.PrepareUploadDriveFile(ctx, &lark.PrepareUploadDriveFileReq{
			FileName:   file.FileName,
			ParentType: "explorer",
			ParentNode: r.folderToken,
			Size:       file.Size,
		})
		if err != nil {
			return nil, err
		}

		for seq := 0; ; seq++ {
			bs, err := ioutil.ReadAll(io.LimitReader(file.File, prepareResp.BlockSize))
			if err != nil {
				return nil, err
			}
			if len(bs) == 0 {
				break
			}
			_, _, err = r.larkClient.Drive.PartUploadDriveFile(ctx, &lark.PartUploadDriveFileReq{
				UploadID: prepareResp.UploadID,
				Seq:      int64(seq),
				Size:     int64(len(bs)),
				Checksum: nil,
				File:     bytes.NewReader(bs),
			})
			if err != nil {
				return nil, err
			}
			if int64(len(bs)) < prepareResp.BlockSize {
				break
			}
		}
		resp, _, err := r.larkClient.Drive.FinishUploadDriveFile(ctx, &lark.FinishUploadDriveFileReq{
			UploadID: prepareResp.UploadID,
			BlockNum: prepareResp.BlockNum,
		})
		if err != nil {
			return nil, err
		}
		token = resp.FileToken
	}

	return newFile(r.larkClient, token, ""), nil
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
