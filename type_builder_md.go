package lark

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
	return "<at email=all></at>"
}

// AtUser 艾特
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

// Image 图片
func (r mdBuilder) Image(imageKey, hoverText string) string {
	return "![" + hoverText + "](" + imageKey + ")"
}
