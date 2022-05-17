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
package doc

import (
	"encoding/json"

	"github.com/chyroc/lark"
)

// UpdateParagraphStyleRequest ...
func UpdateParagraphStyleRequest(location *lark.DocLocation, styles ...*UpdateParagraphStyle) *lark.UpdateDocRequest {
	style := new(lark.DocParagraphStyle)
	var masks []lark.UpdateDocMaskEnum
	for _, styleSet := range styles {
		masks = append(masks, styleSet.mask)
		switch styleSet.mask {
		case lark.UpdateDocMaskEnumQuote:
			style.Quote = styleSet.Quote
		case lark.UpdateDocMaskEnumList:
			style.List = styleSet.List
		case lark.UpdateDocMaskEnumHeadingLevel:
			style.HeadingLevel = styleSet.HeadingLevel
		case lark.UpdateDocMaskEnumAlign:
			style.Align = styleSet.Align
		}
	}
	bs, _ := json.Marshal(style)
	return &lark.UpdateDocRequest{
		RequestType: lark.UpdateDocRequestTypeUpdateParagraphStyle,
		UpdateParagraphStyleRequest: &lark.UpdateDocUpdateParagraphStyleRequest{
			Payload: string(bs),
			Range:   location,
			Fields: &lark.UpdateDocFieldMask{
				Masks: masks,
			},
		},
	}
}

func SetParagraphStyleHeadingLevel(val int64) *UpdateParagraphStyle {
	return &UpdateParagraphStyle{
		mask: lark.UpdateDocMaskEnumHeadingLevel,
		DocParagraphStyle: lark.DocParagraphStyle{
			HeadingLevel: val,
		},
	}
}

func SetParagraphStyleList(val *lark.DocStyleList) *UpdateParagraphStyle {
	return &UpdateParagraphStyle{
		mask: lark.UpdateDocMaskEnumList,
		DocParagraphStyle: lark.DocParagraphStyle{
			List: val,
		},
	}
}

func SetParagraphStyleQuote(val bool) *UpdateParagraphStyle {
	return &UpdateParagraphStyle{
		mask: lark.UpdateDocMaskEnumQuote,
		DocParagraphStyle: lark.DocParagraphStyle{
			Quote: val,
		},
	}
}

func SetParagraphStyleAlign(val string) *UpdateParagraphStyle {
	return &UpdateParagraphStyle{
		mask: lark.UpdateDocMaskEnumAlign,
		DocParagraphStyle: lark.DocParagraphStyle{
			Align: val,
		},
	}
}

type UpdateParagraphStyle struct {
	lark.DocParagraphStyle
	mask lark.UpdateDocMaskEnum
}
