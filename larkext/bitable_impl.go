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

func (r *Bitable) meta(ctx context.Context) (*lark.GetBitableMetaRespApp, error) {
	resp, _, err := r.larkClient.Bitable.GetBitableMeta(ctx, &lark.GetBitableMetaReq{
		AppToken: r.token,
	})
	if err != nil {
		return nil, err
	}
	return resp.App, err
}

func (r *Bitable) updateMeta(ctx context.Context, name *string, isAdvanced *bool) error {
	_, _, err := r.larkClient.Bitable.UpdateBitableMeta(ctx, &lark.UpdateBitableMetaReq{
		AppToken:   r.token,
		Name:       name,
		IsAdvanced: isAdvanced,
	})
	return err
}

func (r *Bitable) copy(ctx context.Context, folderToken, name string) (*Bitable, error) {
	res, err := copyFile(ctx, r.larkClient, folderToken, r.token, r.typ, name)
	if err != nil {
		return nil, err
	}
	return newBitable(r.larkClient, res.Token, res.URL), nil
}

func (r *Bitable) createBitableTable(ctx context.Context, name string) (*BitableTable, error) {
	resp, _, err := r.larkClient.Bitable.CreateBitableTable(ctx, &lark.CreateBitableTableReq{
		AppToken: r.token,
		Table: &lark.CreateBitableTableReqTable{
			Name: &name,
		},
	})
	if err != nil {
		return nil, err
	}
	return NewBitableTable(r.larkClient, r.token, resp.TableID), nil
}
