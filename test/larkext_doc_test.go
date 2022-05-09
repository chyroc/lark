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

	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_LarkExt_Doc(t *testing.T) {
	as := assert.New(t)

	larkClient := AppAllPermission.Ins()
	folderClient, err := larkext.NewRootFolder(ctx, larkClient)
	as.Nil(err)

	docClient, err := folderClient.NewDoc(ctx, "doc title")
	as.Nil(err)
	defer func() {
		as.Nil(docClient.Delete(ctx))
	}()

	t.Run("meta", func(t *testing.T) {
		meta, err := docClient.Meta(ctx)
		as.Nil(err)
		as.NotNil(meta)
	})

	t.Run("raw-content", func(t *testing.T) {
		text, err := docClient.RawContent(ctx)
		as.Nil(err)
		as.Equal("", strings.TrimSpace(text))
	})

	t.Run("content", func(t *testing.T) {
		content, err := docClient.Content(ctx)
		as.Nil(err)
		bs, err := json.Marshal(content)
		as.Nil(err)
		as.Equal(`{"title":{"location":{"zoneId":"0"}},"body":{"blocks":[{"type":"paragraph","paragraph":{"location":{"zoneId":"0","startIndex":1,"endIndex":1}}}]}}`, string(bs))
		fmt.Println(string(bs))
	})
}
