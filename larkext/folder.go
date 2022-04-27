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

	"github.com/chyroc/lark"
)

// Folder is folder client
type Folder struct {
	larkClient  *lark.Lark
	folderToken string
}

// NewFolder new folder client
func NewFolder(larkClient *lark.Lark, folderToken string) *Folder {
	r := new(Folder)
	r.larkClient = larkClient
	r.folderToken = folderToken

	return r
}

// NewRootFolder new root folder client
func NewRootFolder(ctx context.Context, larkClient *lark.Lark) (*Folder, error) {
	f := NewFolder(larkClient, "")
	meta, err := f.meta(ctx)
	if err != nil {
		return nil, err
	}
	f.folderToken = meta.Token

	return f, nil
}

// Meta get folder meta
func (r *Folder) Meta(ctx context.Context) (*FolderMeta, error) {
	return r.meta(ctx)
}

// ListFiles list files in folder
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

// NewFolder create new another folder in folder
func (r *Folder) NewFolder(ctx context.Context, title string) (*Folder, error) {
	return r.newFolder(ctx, title)
}

// NewSheet new sheet in folder
func (r *Folder) NewSheet(ctx context.Context, title string) (*Sheet, error) {
	return r.newSheet(ctx, title)
}

// NewDoc new doc in folder
func (r *Folder) NewDoc(ctx context.Context, title string) (*Doc, error) {
	return r.newDoc(ctx, title)
}

// DeleteFile delete file in folder
func (r *Folder) DeleteFile(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "file")
}

// DeleteDocx delete docx in folder
func (r *Folder) DeleteDocx(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "docx")
}

// DeleteBitable delete bitable in folder
func (r *Folder) DeleteBitable(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "bitable")
}

// DeleteFolder delete folder in folder
func (r *Folder) DeleteFolder(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "folder")
}

// DeleteDoc delete doc in folder
func (r *Folder) DeleteDoc(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "doc")
}

// DeleteSheet delete sheet in folder
func (r *Folder) DeleteSheet(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "sheet")
}

// DeleteMindnote delete mindnote in folder
func (r *Folder) DeleteMindnote(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "mindnote")
}

// DeleteShortcut delete shortcut in folder
func (r *Folder) DeleteShortcut(ctx context.Context, fileToken string) error {
	return r.deleteFile(ctx, fileToken, "shortcut")
}

// FolderMeta is folder meta
type FolderMeta struct {
	ID        string `json:"id,omitempty"`        // 文件夹的 id
	Name      string `json:"name,omitempty"`      // 文件夹的标题
	Token     string `json:"token,omitempty"`     // 文件夹的 token
	CreateUid string `json:"createUid,omitempty"` // 文件夹的创建者 id
	EditUid   string `json:"editUid,omitempty"`   // 文件夹的最后编辑者 id
	ParentID  string `json:"parentId,omitempty"`  // 文件夹的上级目录 id
	OwnUid    string `json:"ownUid,omitempty"`    // 文件夹为个人文件夹时，为文件夹的所有者 id；文件夹为共享文件夹时，为文件夹树id
}

type FileMeta struct {
	Token       string `json:"token,omitempty"`        // 文件标识符
	Name        string `json:"name,omitempty"`         // 文件名
	Type        string `json:"type,omitempty"`         // 文件类型
	ParentToken string `json:"parent_token,omitempty"` // 父文件夹标识
	URL         string `json:"url,omitempty"`          // 在浏览器中查看的链接
}

//  - SearchDriveFile
//  - GetDriveFileMeta
//  - CreateDriveFile
//  - GetDriveFolderMeta
//  - GetDriveRootFolderMeta
//  - GetDriveFolderChildren
//  - GetDriveFileStatistics
//  - DownloadDriveFile
//  - CopyDriveFile
//  - CreateDriveFolder
//  - MoveDriveFile
//  - UploadDriveFile
//  - PrepareUploadDriveFile
//  - PartUploadDriveFile
//  - FinishUploadDriveFile
//  - DownloadDriveMedia
//  - UploadDriveMedia
//  - PrepareUploadDriveMedia
//  - PartUploadDriveMedia
//  - FinishUploadDriveMedia
//  - CreateDriveMemberPermissionOld
//  - TransferDriveMemberPermission
//  - GetDriveMemberPermissionList
//  - CreateDriveMemberPermission
//  - DeleteDriveMemberPermission
//  - DeleteDriveMemberPermissionOld
//  - UpdateDriveMemberPermissionOld
//  - UpdateDriveMemberPermission
//  - CheckDriveMemberPermission
//  - UpdateDrivePublicPermissionV1Old
//  - UpdateDrivePublicPermissionV2Old
//  - GetDrivePublicPermissionV2
//  - GetDrivePublicPermission
//  - UpdateDrivePublicPermission
//  - BatchGetDriveMediaTmpDownloadURL
//  - GetDriveCommentList
//  - GetDriveComment
//  - CreateDriveComment
//  - UpdateDriveComment
//  - DeleteDriveComment
//  - UpdateDriveCommentPatch
//  - CreateDriveFileSubscription
//  - GetDriveFileSubscription
//  - UpdateDriveFileSubscription
//  - CreateDriveDoc
//  - GetDriveDocContent
//  - GetDriveDocRawContent
//  - GetDriveDocMeta
//  - CreateSheet
//  - GetSheetMeta
//  - UpdateSheetProperty
//  - BatchUpdateSheet
//  - ImportSheet
//  - CreateDriveImportTask
//  - GetDriveImportTask
//  - MoveSheetDimension
//  - PrependSheetValue
//  - AppendSheetValue
//  - InsertSheetDimensionRange
//  - AddSheetDimensionRange
//  - UpdateSheetDimensionRange
//  - DeleteSheetDimensionRange
//  - GetSheetValue
//  - BatchGetSheetValue
//  - SetSheetValue
//  - BatchSetSheetValue
//  - SetSheetStyle
//  - BatchSetSheetStyle
//  - MergeSheetCell
//  - UnmergeSheetCell
//  - SetSheetValueImage
//  - FindSheet
//  - ReplaceSheet
//  - CreateSheetConditionFormat
//  - GetSheetConditionFormat
//  - UpdateSheetConditionFormat
//  - DeleteSheetConditionFormat
//  - CreateSheetProtectedDimension
//  - GetSheetProtectedDimension
//  - UpdateSheetProtectedDimension
//  - DeleteSheetProtectedDimension
//  - CreateSheetDataValidationDropdown
//  - DeleteSheetDataValidationDropdown
//  - UpdateSheetDataValidationDropdown
//  - GetSheetDataValidationDropdown
//  - CreateSheetFilter
//  - DeleteSheetFilter
//  - UpdateSheetFilter
//  - GetSheetFilter
//  - CreateSheetFilterView
//  - DeleteSheetFilterView
//  - UpdateSheetFilterView
//  - GetSheetFilterView
//  - QuerySheetFilterView
//  - CreateSheetFilterViewCondition
//  - DeleteSheetFilterViewCondition
//  - UpdateSheetFilterViewCondition
//  - GetSheetFilterViewCondition
//  - QuerySheetFilterViewCondition
//  - CreateSheetFloatImage
//  - DeleteSheetFloatImage
//  - UpdateSheetFloatImage
//  - GetSheetFloatImage
//  - QuerySheetFloatImage
//  - CreateWikiSpace
//  - GetWikiSpaceList
//  - GetWikiSpace
//  - UpdateWikiSpaceSetting
//  - DeleteWikiSpaceMember
//  - AddWikiSpaceMember
//  - CreateWikiNode
//  - GetWikiNodeList
//  - MoveWikiNode
//  - GetWikiNode
//  - MoveDocsToWiki
//  - GetWikiTask
