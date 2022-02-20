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
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_SheetContent(t *testing.T) {
	as := assert.New(t)

	content := [][]lark.SheetContent{
		{
			{String: ptr.String("text")},
			{String: ptr.String("123.456")},
			{Link: &lark.SheetValueLink{Text: "text", Link: "https://github.com/"}},
			{String: ptr.String("a@b.com")},
			{AtUser: &lark.SheetValueAtUser{Text: "a@b.com", TextType: "email", Notify: false, GrantReadPermission: true}},
		},
		{
			{Formula: &lark.SheetValueFormula{Text: "=A1"}},
			{AtDoc: &lark.SheetValueAtDoc{ObjType: "doc", Text: "doc-token"}},
			{MultiValue: &lark.SheetValueMultiValue{Values: []interface{}{"1", 123}}},
		},
	}
	bs, err := json.Marshal(content)
	as.Nil(err)
	as.Contains(string(bs), `"text"`)
	as.Contains(string(bs), `123.456`)
	as.Contains(string(bs), `"link":"https://github.com/"`)
	as.Contains(string(bs), `"a@b.com"`)
	as.Contains(string(bs), `"grantReadPermission":true`)
	as.Contains(string(bs), `"text":"=A1"`)
	as.Contains(string(bs), `"values":["1",123]`)
}
