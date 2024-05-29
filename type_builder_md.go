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

import "strings"

// MdBuilder Markdown标签
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uADOwUjLwgDM14CM4ATN
// doc: https://open.feishu.cn/document/common-capabilities/message-card/message-cards-content/using-markdown-tags
// 新卡片: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/rich-text
var MdBuilder mdBuilder

type mdBuilder string

// Italic 斜体
func (mdBuilder) Italic(s string) string {
	return "*" + s + "*"
}

// Bold 加粗
func (mdBuilder) Bold(s string) string {
	return "**" + s + "**"
}

// Strikethrough 删除线
func (mdBuilder) Strikethrough(s string) string {
	return "~~" + s + "~~"
}

// AtAll 艾特所有人
func (mdBuilder) AtAll() string {
	return "<at id=all></at>"
}

// AtUserID 艾特
//
// id 支持 open_id，user_id
func (mdBuilder) AtUserID(id string) string {
	return "<at id=" + id + "></at>"
}

// AtUserIDs 艾特多个人
func (mdBuilder) AtUserIDs(ids ...string) string {
	return "<at ids=" + strings.Join(ids, ",") + "></at>"
}

// AtUserEmail 艾特, 支持 email
func (mdBuilder) AtUserEmail(email string) string {
	return "<at email=" + email + "></at>"
}

// Link 超链接
func (mdBuilder) Link(url, title string) string {
	return "[" + title + "](" + url + ")"
}

// LinkOrigin 超链接(保持原样，防止类似于 &region变成®ion 这种情况)
func (mdBuilder) LinkOrigin(url, title string) string {
	return "[" + escape(title) + "](" + preprocessURL(url) + ")"
}

// GreenText 绿色文本
func (mdBuilder) GreenText(s string) string {
	return "<font color='green'>" + s + "</font>"
}

// RedText 红色文本
func (mdBuilder) RedText(s string) string {
	return "<font color='red'>" + s + "</font>"
}

// GreyText 灰色文本
func (mdBuilder) GreyText(s string) string {
	return "<font color='grey'>" + s + "</font>"
}

// Telephone 可点击的电话号码
func (mdBuilder) Telephone(title string, telephone string) string {
	//  [文本展示的电话号码或其他文案内容](tel://移动端弹窗唤起的电话号码)
	return "[" + title + "](tel://" + telephone + ")"
}

// Image 图片
//
// hover_text 为PC端hover图片展示的tips文案
// image_key 需上传图片到飞书后获得，格式如：img_xxxxx-xxxxx-49ca-bcf6-224624457a99
// 不支持在text对象的lark_md类型中使用
func (mdBuilder) Image(imageKey, hoverText string) string {
	return "![" + hoverText + "](" + imageKey + ")"
}

// HorizontalRule 分割线
//
// 仅支持Markdown模块
// 不支持在text对象的lark_md类型中使用
// --符号必须跟在换行符后使用，且与换行符间有1个空格
func (mdBuilder) HorizontalRule() string {
	return "\n ---\n"
}

// Emoji 飞书表情
func (mdBuilder) Emoji(key EmojiType) string {
	return ":" + string(key) + ":"
}

// TextTagNeutral 中性色标签
func (mdBuilder) TextTagNeutral(title string) string {
	return buildTextTag(title, "neutral")
}

// TextTagBlue 蓝色标签
func (mdBuilder) TextTagBlue(title string) string {
	return buildTextTag(title, "blue")
}

// TextTagTurquoise 青绿色标签
func (mdBuilder) TextTagTurquoise(title string) string {
	return buildTextTag(title, "turquoise")
}

// TextTagLime 酸橙色标签
func (mdBuilder) TextTagLime(title string) string {
	return buildTextTag(title, "lime")
}

// TextTagOrange 橙色标签
func (mdBuilder) TextTagOrange(title string) string {
	return buildTextTag(title, "orange")
}

// TextTagViolet 紫罗兰色标签
func (mdBuilder) TextTagViolet(title string) string {
	return buildTextTag(title, "violet")
}

// TextTagIndigo 靛青色标签
func (mdBuilder) TextTagIndigo(title string) string {
	return buildTextTag(title, "indigo")
}

// TextTagWathet 天蓝色标签
func (mdBuilder) TextTagWathet(title string) string {
	return buildTextTag(title, "wathet")
}

// TextTagGreen 绿色标签
func (mdBuilder) TextTagGreen(title string) string {
	return buildTextTag(title, "green")
}

// TextTagYellow 黄色标签
func (mdBuilder) TextTagYellow(title string) string {
	return buildTextTag(title, "yellow")
}

// TextTagRed 红色标签
func (mdBuilder) TextTagRed(title string) string {
	return buildTextTag(title, "red")
}

// TextTagPurple 紫色标签
func (mdBuilder) TextTagPurple(title string) string {
	return buildTextTag(title, "purple")
}

// TextTagCarmine 洋红色标签
func (mdBuilder) TextTagCarmine(title string) string {
	return buildTextTag(title, "carmine")
}

// CodeBlock 代码块
func (mdBuilder) CodeBlock(language, code string) string {
	return "```" + language + "\n" + code + "\n```"
}

func buildTextTag(title, color string) string {
	return "<text_tag color='" + color + "'>" + title + "</text_tag>"
}

var reservedWordsMapping = map[string]string{
	" ":  "&nbsp;", // 空格
	"~":  "&sim;",  // 波浪线
	"-":  "&#45;",  // 短横线
	"!":  "&#33;",  // 感叹号
	"*":  "&#42;",  // 星号
	"/":  "&#47;",  // 斜杠
	"\\": "&#92;",  // 反斜杠
	"[":  "&#91;",  // 方括号
	"]":  "&#93;",  // 方括号
	"(":  "&#40;",  // 小括号
	")":  "&#41;",  // 小括号
	"#":  "&#35;",  // 井号
	":":  "&#58;",  // 冒号
	"+":  "&#43;",  // 加号
	"'":  "&#39;",  // 单引号
	"`":  "&#96;",  // 反引号
	"$":  "&#36;",  // 美元符号
	"|":  "&#124;", // 竖线
	">":  "&gt;",   // 大于号
	"<":  "&lt;",   // 小于号
	"@":  "&#64;",  // @
}

// markdown保留字转义
func escape(txt string) string {
	for k, v := range reservedWordsMapping {
		txt = strings.ReplaceAll(txt, k, v)
	}
	return txt
}

// url保留字转义
func preprocessURL(urlStr string) string {
	urlStr = strings.ReplaceAll(urlStr, "&reg", "&&#114;eg")
	urlStr = strings.ReplaceAll(urlStr, "&para", "&&#112;ara")
	return urlStr
}
