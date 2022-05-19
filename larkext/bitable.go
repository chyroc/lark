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

// Bitable Bitable client
type Bitable struct {
	larkClient *lark.Lark
	token      string
	url        string
	typ        string
}

// NewBitable new Bitable client
func NewBitable(larkClient *lark.Lark, appToken string) *Bitable {
	return newBitable(larkClient, appToken, "")
}

func newBitable(larkClient *lark.Lark, token, url string) *Bitable {
	return &Bitable{larkClient: larkClient, token: token, url: url, typ: "bitable"}
}

// Meta get bitable meta
func (r *Bitable) Meta(ctx context.Context) (*lark.GetBitableMetaRespApp, error) {
	return r.meta(ctx)
}

// Copy copy bitable file
func (r *Bitable) Copy(ctx context.Context, folderToken, name string) (*Bitable, error) {
	return r.copy(ctx, folderToken, name)
}

func (r *Bitable) Move(ctx context.Context, folderToken string) error {
	_, err := moveFile(ctx, r.larkClient, folderToken, r.token, r.typ)
	return err
}

func (r *Bitable) Delete(ctx context.Context) error {
	_, err := deleteFile(ctx, r.larkClient, r.token, r.typ)
	return err
}

// Permission grant bitable permission
func (r *Bitable) Permission() *Permission {
	return newPermission(r.larkClient, r.token, r.typ)
}
