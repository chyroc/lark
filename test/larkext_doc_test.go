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
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_LarkExt_Doc(t *testing.T) {
	SkipRealAPITest(t)
	as := assert.New(t)

	larkClient := AppAllPermission.Ins()
	folderClient, err := larkext.NewRootFolder(ctx, larkClient)
	if !as.Nil(err) {
		return
	}

	docClient, err := folderClient.NewDoc(ctx, "doc title")
	if err != nil && lark.GetErrorCode(err) == 95054 {
		t.Skipf("skip because CreateDriveDoc is deprecated: %v", err)
		return
	}
	if !as.Nil(err) || !as.NotNil(docClient) {
		return
	}
	defer func() {
		_, err := docClient.Delete(ctx)
		as.Nil(err)
	}()

	t.Run("meta", func(t *testing.T) {
		meta, err := docClient.Meta(ctx)
		as.Nil(err)
		if !as.NotNil(meta) {
			return
		}
	})

	t.Run("raw-content", func(t *testing.T) {
		text, err := docClient.RawContent(ctx)
		as.Nil(err)
		as.Equal("doc title", strings.TrimSpace(text))
	})

	t.Run("content", func(t *testing.T) {
		content, err := docClient.Content(ctx)
		as.Nil(err)
		bs, err := json.Marshal(content)
		as.Nil(err)
		as.NotEmpty(bs)
		fmt.Println(string(bs))
	})
}
