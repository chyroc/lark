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
package lark

func (r *DocContent) GetParagraphElementLocation(pe *DocParagraphElement) *DocLocation {
	for _, ele := range r.Title.Elements {
		if ele.Equal(pe) {
			return ele.Location()
		}
	}

	for _, block := range r.Body.Blocks {
		if block.Type != DocBlockTypeParagraph {
			continue
		}
		for _, ele := range block.Paragraph.Elements {
			if ele.Equal(pe) {
				return ele.Location()
			}
		}
	}
	return nil
}

func (r *DocContent) GetParagraph(p *DocParagraph) *DocParagraph {
	if p.Equal(r.Title) {
		return r.Title
	}
	for _, block := range r.Body.Blocks {
		if block.Type != DocBlockTypeParagraph {
			continue
		}
		if p.Equal(block.Paragraph) {
			return block.Paragraph
		}
	}
	return nil
}

func (r *DocContent) GetTableBySize(p *DocTable) *DocTable {
	for _, block := range r.Body.Blocks {
		if block.Type != DocBlockTypeTable {
			continue
		}
		if p.EqualBySize(block.Table) {
			return block.Table
		}
	}
	return nil
}

func (r *DocParagraphElement) Equal(r2 *DocParagraphElement) bool {
	if r == nil && r2 == nil {
		return true
	}
	if r == nil || r2 == nil {
		return false
	}
	if r.Type != r2.Type {
		return false
	}
	switch r.Type {
	case DocParagraphElementTypeTextRun:
		return r.TextRun.Text == r2.TextRun.Text
	case DocParagraphElementTypeDocsLink:
		return r.DocsLink.URL == r2.DocsLink.URL
	case DocParagraphElementTypePerson:
		return r.Person.OpenID == r2.Person.OpenID
	case DocParagraphElementTypeEquation:
		return r.Equation.Equation == r2.Equation.Equation
	case DocParagraphElementTypeReminder:
		return false // 不支持
	case DocParagraphElementTypeFile:
		return r.File.FileToken == r2.File.FileToken
	case DocParagraphElementTypeJira:
		return r.Jira.Token == r2.Jira.Token
	}

	return false
}

func (r *DocParagraph) Equal(r2 *DocParagraph) bool {
	if r == nil && r2 == nil {
		return true
	}
	if r == nil || r2 == nil {
		return false
	}
	if len(r.Elements) != len(r2.Elements) {
		return false
	}
	for i, ele := range r.Elements {
		if !ele.Equal(r2.Elements[i]) {
			return false
		}
	}
	return true
}

func (r *DocParagraphElement) Location() *DocLocation {
	switch r.Type {
	case DocParagraphElementTypeTextRun:
		return r.TextRun.Location
	case DocParagraphElementTypeDocsLink:
		return r.DocsLink.Location
	case DocParagraphElementTypePerson:
		return r.Person.Location
	case DocParagraphElementTypeEquation:
		return r.Equation.Location
	case DocParagraphElementTypeReminder:
		return r.Reminder.Location
	case DocParagraphElementTypeFile:
		return r.File.Location
	case DocParagraphElementTypeJira:
		return r.Jira.Location
	}
	return nil
}

func (r *DocTable) EqualBySize(r2 *DocTable) bool {
	if r == nil && r2 == nil {
		return true
	}
	if r == nil || r2 == nil {
		return false
	}
	if r.RowSize != r2.RowSize {
		return false
	}
	if r.ColumnSize != r2.ColumnSize {
		return false
	}
	return true
}
