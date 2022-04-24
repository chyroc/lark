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
package card

import (
	"github.com/chyroc/lark"
)

// Text plain_text
func Text(text string) *lark.MessageContentCardObjectText {
	return &lark.MessageContentCardObjectText{
		Tag:     lark.MessageContentCardTextTypePlainText,
		Content: text,
	}
}

func I18nText(text *lark.I18NText) *lark.MessageContentCardObjectText {
	return &lark.MessageContentCardObjectText{
		Tag:  lark.MessageContentCardTextTypePlainText,
		I18N: text,
	}
}

// MarkdownText lark_md
func MarkdownText(md string) *lark.MessageContentCardObjectText {
	return &lark.MessageContentCardObjectText{
		Tag:     lark.MessageContentCardTextTypeLarkMd,
		Content: md,
	}
}
