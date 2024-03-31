package card

import "encoding/json"

// Divider 分割线
func Divider() *ComponentDivider {
	return &ComponentDivider{}
}

// ComponentDivider 分割线
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/divider
//
// 你可以在卡片中添加分割线组件, 使卡片内容更清晰。
type ComponentDivider struct {
	componentBase
}

// MarshalJSON ...
func (r ComponentDivider) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{"tag": "hr"})
}
