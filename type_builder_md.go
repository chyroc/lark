package lark

import "strings"

// MdBuilder Markdown标签
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uADOwUjLwgDM14CM4ATN#top_anchor
var MdBuilder mdBuilder

type mdBuilder string

// Italic 斜体
func (r mdBuilder) Italic(s string) string {
	return "*" + s + "*"
}

// Bold 加粗
func (r mdBuilder) Bold(s string) string {
	return "**" + s + "**"
}

// Strikethrough 删除线
func (r mdBuilder) Strikethrough(s string) string {
	return "~~" + s + "~~"
}

// AtAll 艾特所有人
func (r mdBuilder) AtAll() string {
	return "<at id=all></at>"
}

// AtUser 艾特
//
// id 支持 open_id，user_id
func (r mdBuilder) AtUserID(id string) string {
	return "<at id=" + id + "></at>"
}

// AtUser 艾特
func (r mdBuilder) AtUserEmail(email string) string {
	return "<at email=" + email + "></at>"
}

// Link 超链接
func (r mdBuilder) Link(url, title string) string {
	return "[" + title + "](" + url + ")"
}

var reservedWordsMapping = map[string]string{
	"|": "&#124;",
	"`": "&#96;",
	"]": "&#93;",
	"[": "&#91;",
	">": "&gt;",
	"<": "&lt;",
	"@": "&#64;",
	"#": "&#35;",
	"-": "&#45;",
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
	return urlStr
}

// LinkOrigin 超链接(保持原样，防止类似于 &region变成®ion 这种情况)
func (r mdBuilder) LinkOrigin(url, title string) string {
	return "[" + escape(title) + "](" + preprocessURL(url) + ")"
}

// Image 图片
//
// hover_text 为PC端hover图片展示的tips文案
// image_key 需上传图片到飞书后获得，格式如：img_xxxxx-xxxxx-49ca-bcf6-224624457a99
// 不支持在text对象的lark_md类型中使用
func (r mdBuilder) Image(imageKey, hoverText string) string {
	return "![" + hoverText + "](" + imageKey + ")"
}
