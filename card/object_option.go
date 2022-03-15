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

// option 有三种使用方式，分别是：selectMenu选项，selectMenu选人，overflow选项

// SelectOption 基础使用
func SelectOption(text, value string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Text:  Text(text),
		Value: value,
	}
}

// PersonOption 选人模式
// 指定待选人员，只需指定待选人的openId
func PersonOption(id string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Value: id,
	}
}

// LinkOption 跳转链接 option
// overflow 中可用
func LinkOption(text, url string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Text: Text(text),
		URL:  url,
	}
}
