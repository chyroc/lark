package card

import "encoding/json"

// Form 表单容器
func Form(name string, elements ...Component) *ComponentForm {
	return &ComponentForm{
		Name:     name,
		Elements: elements,
	}
}

// ComponentForm 表单容器
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/form-container
//
// 在使用卡片收集内容时, 可能存在需要用户提交多个表单项的场景。表单容器允许用户在前端本地录入一批表单项后, 通过点击一次 提交 按钮, 将这一批本地缓存的表单内容一次回调至开发者的服务端, 实现异步提交多个表单项数据的效果。
//
// 注意事项
//
// 表单容器支持飞书 V6.6 及以上版本的客户端。在低于该版本的飞书客户端上, 表单容器的内容将展示为“请升级至最新版本客户端, 以查看内容”的占位图。
// 容器类组件最多支持嵌套五层组件。建议你避免在表单容器中嵌套多层组件。多层嵌套会压缩内容的展示空间, 影响卡片的展示效果。如你希望卡片承接更复杂的表单内容, 建议通过卡片链接跳转至 H5 或小程序实现表单能力。
//
// 嵌套规则
// 表单容器不支持内嵌表格、图表、和表单容器组件。
// 表单容器中不可直接内嵌标签为 div 的组件。你可先内嵌分栏组件, 再在分栏组件中内嵌标签为 div 类型的组件。
// 表单容器组件不可被内嵌在其它组件内, 只可放在卡片根节点下。
//
//go:generate generate_set_attrs -type=ComponentForm
//go:generate generate_to_map -type=ComponentForm
type ComponentForm struct {
	componentBase

	// 表单容器的唯一标识。用于识别用户提交的数据属于哪个表单容器。在同一张卡片内, 该字段的值全局唯一。
	Name string `json:"name,omitempty"`

	// 表单容器的子节点。可内嵌其它容器类组件和展示、交互组件, 不支持内嵌表格、图表、和表单容器组件。
	Elements []Component `json:"elements,omitempty"`
}

// MarshalJSON ...
func (r ComponentForm) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "form"
	return json.Marshal(m)
}

// SetName set ComponentForm.Name attribute
func (r *ComponentForm) SetName(val string) *ComponentForm {
	r.Name = val
	return r
}

// SetElements set ComponentForm.Elements attribute
func (r *ComponentForm) SetElements(val ...Component) *ComponentForm {
	r.Elements = val
	return r
}

// toMap conv ComponentForm to map
func (r *ComponentForm) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 2)
	if r.Name != "" {
		res["name"] = r.Name
	}
	if len(r.Elements) != 0 {
		res["elements"] = r.Elements
	}
	return res
}
