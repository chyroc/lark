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

// Button 回传数据按钮
func Button(text string, value interface{}) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Text:       Text(text),
		Value:      value,
		Type:       "default",
		ActionType: lark.MessageContentCardElementButtonRequest,
	}
}

// SubmitButton 触发表单容器的提交事件
func SubmitButton(name, text string, value interface{}) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Name:       name,
		Text:       Text(text),
		Value:      value,
		Type:       "default",
		ActionType: lark.MessageContentCardElementButtonFormSubmit,
	}
}

// ResetButton 触发表单容器的取消提交事件
func ResetButton(name, text string, value interface{}) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Name:       name,
		Text:       Text(text),
		Value:      value,
		Type:       "default",
		ActionType: lark.MessageContentCardElementButtonFormReset,
	}
}

func LinkButton(text string, url string) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Text:       Text(text),
		URL:        url,
		ActionType: lark.MessageContentCardElementButtonMulti,
	}
}

func MultiLinkButton(text string, url *lark.MessageContentCardObjectURL) *lark.MessageContentCardElementButton {
	return &lark.MessageContentCardElementButton{
		Text:       Text(text),
		MultiURL:   url,
		ActionType: lark.MessageContentCardElementButtonMulti,
	}
}
