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

// Doc is doc client
type Doc struct {
	larkClient *lark.Lark
	docToken   string
}

// NewDoc new doc client
func NewDoc(larkClient *lark.Lark, docToken string) *Doc {
	r := new(Doc)
	r.larkClient = larkClient
	r.docToken = docToken
	return r
}

// DocToken get doc Token
func (r *Doc) DocToken() string {
	return r.docToken
}

// Meta get doc meta
func (r *Doc) Meta(ctx context.Context) (*lark.GetDriveDocMetaResp, error) {
	return r.meta(ctx)
}

// Delete delete doc
func (r *Doc) Delete(ctx context.Context) error {
	return r.delete(ctx)
}

// Copy copy doc file
func (r *Doc) Copy(ctx context.Context, folderToken, name string) (*FileMeta, error) {
	return copyFile(ctx, r.larkClient, folderToken, r.docToken, "doc", name)
}

// RawContent get doc raw content
func (r *Doc) RawContent(ctx context.Context) (string, error) {
	return r.rawContent(ctx)
}

// RawContent get doc rich content
func (r *Doc) Content(ctx context.Context) (*lark.DocContent, error) {
	return r.content(ctx)
}
