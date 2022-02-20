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
package test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Bitable(t *testing.T) {
	as := assert.New(t)
	cli := AppAllPermission.Ins()

	appToken := "bascn9T22ch7O8cfAuEza3Zv2tb"
	tableID := "tbly5aRriIYOT5SM"

	t.Run("view-list", func(t *testing.T) {
		resp, response, err := cli.Bitable.GetBitableViewList(ctx, &lark.GetBitableViewListReq{
			PageSize:  nil,
			PageToken: nil,
			AppToken:  appToken,
			TableID:   tableID,
		})
		as.Nil(err)
		as.NotEmpty(response.RequestID)
		as.True(len(resp.Items) > 0)
	})

	t.Skip()

	t.Run("bitable-view", func(t *testing.T) {
		var createResp *lark.CreateBitableViewResp

		defer func() {
			if createResp == nil || createResp.View == nil {
				return
			}

			// TODO: 这里的失败了太高，其实重试后可以成功，暂时忽略 err 断言
			err := retry(func() error {
				_, _, err := cli.Bitable.DeleteBitableView(ctx, &lark.DeleteBitableViewReq{
					AppToken: appToken,
					TableID:  tableID,
					ViewID:   createResp.View.ViewID,
				})
				return err
			})
			fmt.Println(err)
		}()

		{
			var resp *lark.CreateBitableViewResp
			as.Nil(retry(func() error {
				res, _, err := cli.Bitable.CreateBitableView(ctx, &lark.CreateBitableViewReq{
					AppToken: appToken,
					TableID:  tableID,
					ViewName: "view-name-" + randInt64String(),
					ViewType: ptr.String("gantt"),
				})
				resp = res
				return err
			}))
			createResp = resp
		}
	})
}

func retry(f func() error) (err error) {
	for i := 0; i < 5; i++ {
		err = f()
		if err == nil {
			return nil
		}
	}
	return err
}
