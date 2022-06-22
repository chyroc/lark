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
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Type(t *testing.T) {
	as := assert.New(t)

	t.Run("SheetContent", func(t *testing.T) {
		t.Run("children", func(t *testing.T) {
			children := []*lark.SheetContent{{String: ptr.String("str")}}
			res := lark.SheetContent{
				Children: &children,
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("str", func(t *testing.T) {
			res := lark.SheetContent{
				String: ptr.String("str"),
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("int", func(t *testing.T) {
			res := lark.SheetContent{
				Int: ptr.Int64(1),
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("link", func(t *testing.T) {
			res := lark.SheetContent{
				Link: &lark.SheetValueLink{
					Text: "text",
					Link: "link",
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("user", func(t *testing.T) {
			res := lark.SheetContent{
				AtUser: &lark.SheetValueAtUser{
					Text:                "email",
					TextType:            "email",
					Notify:              true,
					GrantReadPermission: true,
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("formula", func(t *testing.T) {
			res := lark.SheetContent{
				Formula: &lark.SheetValueFormula{
					Text: "formula",
				},
				// AtDoc: &lark.SheetValueAtDoc{
				// 	TextType: "sheet",
				// 	Text:     "sheet",
				// },
				// MultiValue: &lark.SheetValueMultiValue{
				// 	Values: []interface{}{"1", "2"},
				// },
				// EmbedImage: &lark.SheetValueEmbedImage{
				// 	FileToken: "token",
				// 	Height:    100,
				// 	Link:      "link",
				// 	Text:      "text",
				// 	Type:      "image",
				// 	Width:     100,
				// },
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("doc", func(t *testing.T) {
			res := lark.SheetContent{
				AtDoc: &lark.SheetValueAtDoc{
					TextType: "sheet",
					Text:     "sheet",
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("multi", func(t *testing.T) {
			res := lark.SheetContent{
				MultiValue: &lark.SheetValueMultiValue{
					Values: []interface{}{"1", "2"},
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("image", func(t *testing.T) {
			res := lark.SheetContent{
				EmbedImage: &lark.SheetValueEmbedImage{
					FileToken: "token",
					Height:    100,
					Link:      "link",
					Text:      "text",
					Type:      "image",
					Width:     100,
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})
	})

	t.Run("DocContent", func(t *testing.T) {
		res := lark.DocContent{
			Title: &lark.DocParagraph{
				Elements: []*lark.DocParagraphElement{
					{
						Type:             "",
						TextRun:          nil,
						DocsLink:         nil,
						Person:           nil,
						Equation:         nil,
						Reminder:         nil,
						File:             nil,
						Jira:             nil,
						UndefinedElement: nil,
					},
				},
				Location: &lark.DocLocation{
					ZoneID:     "id",
					StartIndex: 1,
					EndIndex:   2,
				},
				LineID: "3",
			},
			Body: &lark.DocBody{
				Blocks: []*lark.DocBlock{
					{
						Type:           "",
						Paragraph:      nil,
						Gallery:        nil,
						File:           nil,
						ChatGroup:      nil,
						Table:          nil,
						HorizontalLine: nil,
						EmbeddedPage:   nil,
						Sheet:          nil,
						Bitable:        nil,
						Diagram:        nil,
						Jira:           nil,
						Poll:           nil,
						Code:           nil,
						DocsApp:        nil,
						Callout:        nil,
						UndefinedBlock: nil,
					},
				},
			},
		}
		bs, err := json.Marshal(res)
		as.Nil(err)

		obj := new(lark.DocContent)
		as.Nil(json.Unmarshal(bs, obj))

		as.Equal(res, *obj)
	})

	t.Run("md-builder", func(t *testing.T) {
		as.Equal(`<at id=all></at>`, lark.MdBuilder.AtAll())
		as.Equal(`[title](url)`, lark.MdBuilder.Link("url", "title"))
		as.Equal(`[&gt;](&&#114;egion)`, lark.MdBuilder.LinkOrigin("&region", ">"))
		as.Equal(`![hover](key)`, lark.MdBuilder.Image("key", "hover"))
		as.Equal(`<at email=email></at>`, lark.MdBuilder.AtUserEmail("email"))
		as.Equal(`**bold**`, lark.MdBuilder.Bold("bold"))
		as.Equal(`<at id=user></at>`, lark.MdBuilder.AtUserID("user"))
		as.Equal(`*italic*`, lark.MdBuilder.Italic("italic"))
		as.Equal(`~~Strikethrough~~`, lark.MdBuilder.Strikethrough("Strikethrough"))
	})

	t.Run("at-builder", func(t *testing.T) {
		as.Equal(`<at user_id="1"></at>`, lark.AtBuilder.AtOpenID("1"))
		as.Equal(`<at user_id="all"></at>`, lark.AtBuilder.AtAll())
	})

	t.Run("doc-content", func(t *testing.T) {
		tests := []struct {
			name       string
			arg        *lark.DocContent
			want       string
			errContain string
		}{
			{name: "0", arg: &lark.DocContent{}, want: `{}`},
			{name: "1", arg: &lark.DocContent{Title: &lark.DocParagraph{}}, want: `{"title":{}}`},
		}
		for _, tt := range tests {
			t.Run(tt.name+" marshal", func(t *testing.T) {
				got, err := json.Marshal(tt.arg)
				if tt.errContain != "" {
					as.NotNil(err, tt.name)
					as.Empty(got, tt.name)
				} else {
					as.Nil(err, tt.name)
					as.NotEmpty(got, tt.name)
					as.Equal(tt.want, string(got), tt.name)
				}
			})
			t.Run(tt.name+" unmarshal", func(t *testing.T) {
				resp := new(lark.DocContent)
				err := json.Unmarshal([]byte(tt.want), resp)
				if tt.errContain != "" {
					as.NotNil(err, tt.name)
				} else {
					as.Nil(err, tt.name)
					a, _ := json.Marshal(tt.arg)
					b, _ := json.Marshal(resp)
					as.EqualValues(string(b), string(a), tt.name)
				}
			})
		}
	})

	t.Run("type-ptr", func(t *testing.T) {
		as.Equal(lark.IDTypeOpenID, *lark.IDTypePtr(lark.IDTypeOpenID))
		as.Equal(lark.DepartmentIDTypeOpenDepartmentID, *lark.DepartmentIDTypePtr(lark.DepartmentIDTypeOpenDepartmentID))
		as.Equal(lark.AddMemberPermissionOnlyOwner, *lark.AddMemberPermissionPtr(lark.AddMemberPermissionOnlyOwner))
		as.Equal(lark.AtAllPermissionOnlyOwner, *lark.AtAllPermissionPtr(lark.AtAllPermissionOnlyOwner))
		as.Equal(lark.EditPermissionAllMembers, *lark.EditPermissionPtr(lark.EditPermissionAllMembers))
		as.Equal(lark.ModerationPermissionAllMembers, *lark.ModerationPermissionPtr(lark.ModerationPermissionAllMembers))
		as.Equal(lark.ShareCardPermissionNotAllowed, *lark.ShareCardPermissionPtr(lark.ShareCardPermissionNotAllowed))
		as.Equal(lark.MembershipApprovalNoApprovalRequired, *lark.MembershipApprovalPtr(lark.MembershipApprovalNoApprovalRequired))
		as.Equal(lark.MessageVisibilityAllMembers, *lark.MessageVisibilityPtr(lark.MessageVisibilityAllMembers))
	})
}
