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

// === MessageContentCard ===

// MessageContentCard ...
type MessageContentCard struct {
	Header  *MessageContentCardHeader  `json:"header,omitempty"`   // 用于配置卡片标题内容。
	Config  *MessageContentCardConfig  `json:"config,omitempty"`   // 配置卡片属性
	Modules []MessageContentCardModule `json:"elements,omitempty"` // 用于定义卡片正文内容
}

// String ...
func (r MessageContentCard) String() string {
	return jsonString(r)
}

// // UnmarshalJSON ...
// func (r *MessageContentCard) UnmarshalJSON(bytes []byte) error {
// 	var post struct {
// 		Header  *MessageContentCardHeader `json:"header,omitempty"`   // 用于配置卡片标题内容。
// 		Config  *MessageContentCardConfig `json:"config,omitempty"`   // 配置卡片属性
// 		Modules []json.RawMessage         `json:"elements,omitempty"` // 用于定义卡片正文内容
// 	}
// 	if err := json.Unmarshal(bytes, &post); err != nil {
// 		return err
// 	}
//
// 	r.Header = post.Header
// 	r.Config = post.Config
// 	r.Modules = make([]MessageContentCardModule, 0)
//
// 	for vIdx, v := range post.Modules {
// 		_ = vIdx
// 		var module MessageContentCardModule
// 		if err := json.Unmarshal(v, &module); err != nil {
// 			return err
// 		}
// 		r.Modules = append(r.Modules, module)
// 	}
//
// 	panic("implement me")
// }

func (r *MessageContentCard) SetHeader(val *MessageContentCardHeader) *MessageContentCard {
	r.Header = val
	return r
}

func (r *MessageContentCard) SetConfig(val *MessageContentCardConfig) *MessageContentCard {
	r.Config = val
	return r
}

func (r *MessageContentCard) SetModules(val ...MessageContentCardModule) *MessageContentCard {
	r.Modules = val
	return r
}

// === MessageContentCard ===

// === MessageContentCardHeader ===

// MessageContentCardHeader 卡片标题
//
// https://open.feishu.cn/document/ukTMukTMukTM/ukTNwUjL5UDM14SO1ATN
type MessageContentCardHeader struct {
	Template MessageContentCardHeaderTemplate `json:"template,omitempty"` // 控制标题背景颜色，取值参考注意事项
	Title    *MessageContentCardObjectText    `json:"title,omitempty"`    // 卡片标题
}

func (r *MessageContentCardHeader) SetBlue() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateBlue
	return r
}

func (r *MessageContentCardHeader) SetWathet() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateWathet
	return r
}

func (r *MessageContentCardHeader) SetTurquoise() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateTurquoise
	return r
}

func (r *MessageContentCardHeader) SetGreen() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateGreen
	return r
}

func (r *MessageContentCardHeader) SetYellow() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateYellow
	return r
}

func (r *MessageContentCardHeader) SetOrange() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateOrange
	return r
}

func (r *MessageContentCardHeader) SetRed() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateRed
	return r
}

func (r *MessageContentCardHeader) SetCarmine() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateCarmine
	return r
}

func (r *MessageContentCardHeader) SetViolet() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateViolet
	return r
}

func (r *MessageContentCardHeader) SetPurple() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplatePurple
	return r
}

func (r *MessageContentCardHeader) SetIndigo() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateIndigo
	return r
}

func (r *MessageContentCardHeader) SetGrey() *MessageContentCardHeader {
	r.Template = MessageContentCardHeaderTemplateGrey
	return r
}

func (r *MessageContentCardHeader) SetEnableForward(val *MessageContentCardObjectText) *MessageContentCardHeader {
	r.Title = val
	return r
}

// === MessageContentCardHeader ===

// === MessageContentCardConfig ===

// MessageContentCardConfig 配置卡片属性
//
// https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
type MessageContentCardConfig struct {
	EnableForward bool `json:"enable_forward,omitempty"` // 是否允许卡片被转发，默认 true，转发后，卡片上的“回传交互”组件将自动置为禁用态。用户不能在转发后的卡片操作提交数据
	UpdateMulti   bool `json:"update_multi,omitempty"`   // 更新卡片的内容是否对所有收到这张卡片的人员可见。 默认为false，即仅操作用户可见卡片的更新内容。
}

func (r *MessageContentCardConfig) SetEnableForward(val bool) *MessageContentCardConfig {
	r.EnableForward = val
	return r
}

func (r *MessageContentCardConfig) SetUpdateMulti(val bool) *MessageContentCardConfig {
	r.UpdateMulti = val
	return r
}

// === MessageContentCardConfig ===

// MessageContentCardHeaderTemplate 控制标题背景颜色，取值参考注意事项
type MessageContentCardHeaderTemplate string

// MessageContentCardHeaderTemplateBlue ...
const (
	MessageContentCardHeaderTemplateBlue      MessageContentCardHeaderTemplate = "blue"
	MessageContentCardHeaderTemplateWathet    MessageContentCardHeaderTemplate = "wathet"
	MessageContentCardHeaderTemplateTurquoise MessageContentCardHeaderTemplate = "turquoise"
	MessageContentCardHeaderTemplateGreen     MessageContentCardHeaderTemplate = "green"
	MessageContentCardHeaderTemplateYellow    MessageContentCardHeaderTemplate = "yellow"
	MessageContentCardHeaderTemplateOrange    MessageContentCardHeaderTemplate = "orange"
	MessageContentCardHeaderTemplateRed       MessageContentCardHeaderTemplate = "red"
	MessageContentCardHeaderTemplateCarmine   MessageContentCardHeaderTemplate = "carmine"
	MessageContentCardHeaderTemplateViolet    MessageContentCardHeaderTemplate = "violet"
	MessageContentCardHeaderTemplatePurple    MessageContentCardHeaderTemplate = "purple"
	MessageContentCardHeaderTemplateIndigo    MessageContentCardHeaderTemplate = "indigo"
	MessageContentCardHeaderTemplateGrey      MessageContentCardHeaderTemplate = "grey"
)

// MessageContentCardElement ...
type MessageContentCardElement interface {
	IsMessageContentCardElement()
}

// MessageContentCardElementTag ...
type MessageContentCardElementTag string

// MessageContentCardElementTagImage ...
const (
	MessageContentCardElementTagImage          MessageContentCardElementTag = "img"
	MessageContentCardElementTagButton         MessageContentCardElementTag = "button"
	MessageContentCardElementTagOverflow       MessageContentCardElementTag = "overflow"
	MessageContentCardElementTagSelectStatic   MessageContentCardElementTag = "select_static"
	MessageContentCardElementTagSelectPerson   MessageContentCardElementTag = "select_person"
	MessageContentCardElementTagPickerDate     MessageContentCardElementTag = "date_picker"
	MessageContentCardElementTagPickerTime     MessageContentCardElementTag = "picker_time"
	MessageContentCardElementTagPickerDatetime MessageContentCardElementTag = "picker_datetime"
)

// === MessageContentCardElementButton ===

// MessageContentCardElementButton button
//
// button 属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
//
// https://open.feishu.cn/document/ukTMukTMukTM/uEzNwUjLxcDM14SM3ATN
type MessageContentCardElementButton struct {
	Text     *MessageContentCardObjectText    `json:"text,omitempty"`      // 按钮中的文本
	URL      string                           `json:"url,omitempty"`       // 跳转链接，和multi_url互斥
	MultiURL *MessageContentCardObjectURL     `json:"multi_url,omitempty"` // 多端跳转链接
	Type     string                           `json:"type,omitempty"`      // 配置按钮样式，默认为"default"，	"default"/"primary"/"danger"
	Value    interface{}                      `json:"value,omitempty"`     // 点击后返回业务方，	仅支持key-value形式的json结构，且key为String类型。
	Confirm  *MessageContentCardObjectConfirm `json:"confirm,omitempty"`   // 二次确认的弹框
}

// IsMessageContentCardElement ...
func (r MessageContentCardElementButton) IsMessageContentCardElement() {}

// MarshalJSON ...
func (r MessageContentCardElementButton) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagButton})
}

func (r *MessageContentCardElementButton) SetText(val *MessageContentCardObjectText) *MessageContentCardElementButton {
	r.Text = val
	return r
}

func (r *MessageContentCardElementButton) SetURL(val string) *MessageContentCardElementButton {
	r.URL = val
	return r
}

func (r *MessageContentCardElementButton) SetMultiURL(val *MessageContentCardObjectURL) *MessageContentCardElementButton {
	r.MultiURL = val
	return r
}

func (r *MessageContentCardElementButton) SetDefault() *MessageContentCardElementButton {
	r.Type = "default"
	return r
}

func (r *MessageContentCardElementButton) SetPrimary() *MessageContentCardElementButton {
	r.Type = "primary"
	return r
}

func (r *MessageContentCardElementButton) SetDanger() *MessageContentCardElementButton {
	r.Type = "danger"
	return r
}

func (r *MessageContentCardElementButton) SetValue(val interface{}) *MessageContentCardElementButton {
	r.Value = val
	return r
}

func (r *MessageContentCardElementButton) SetConfirm(val *MessageContentCardObjectConfirm) *MessageContentCardElementButton {
	r.Confirm = val
	return r
}

// === MessageContentCardElementButton ===

// === MessageContentCardElementDatePicker ===

// MessageContentCardElementDatePicker 作为datePicker元素被使用，提供时间选择的功能。支持三种模式的时间选择：（1）日期（2）时间（3）日期时间
// datePicker属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQzNwUjL0cDM14CN3ATN
type MessageContentCardElementDatePicker struct {
	Tag             MessageContentCardElementTag     `json:"tag,omitempty"`              // 如下三种取值 ：date_picker, picker_time, picker_datetime
	InitialDate     string                           `json:"initial_date,omitempty"`     // 格式"YYYY-MM-DD"	日期模式的初始值
	InitialTime     string                           `json:"initial_time,omitempty"`     // 格式"HH:mm"	时间模式的初始值
	InitialDatetime string                           `json:"initial_datetime,omitempty"` // 格式"YYYY-MM-DD HH:mm"	日期模式的初始值
	Placeholder     *MessageContentCardObjectText    `json:"placeholder,omitempty"`      // 占位符，无初始值时必填
	Value           interface{}                      `json:"value,omitempty"`            // 用户选定后返回业务方的数据
	Confirm         *MessageContentCardObjectConfirm `json:"confirm,omitempty"`          // 二次确认的弹框
}

// IsMessageContentCardElement ...
func (r MessageContentCardElementDatePicker) IsMessageContentCardElement() {}

func (r *MessageContentCardElementDatePicker) SetInitialDate(val string) *MessageContentCardElementDatePicker {
	r.InitialDate = val
	return r
}

func (r *MessageContentCardElementDatePicker) SetInitialTime(val string) *MessageContentCardElementDatePicker {
	r.InitialTime = val
	return r
}

func (r *MessageContentCardElementDatePicker) SetInitialDatetime(val string) *MessageContentCardElementDatePicker {
	r.InitialDatetime = val
	return r
}

func (r *MessageContentCardElementDatePicker) SetPlaceholder(val *MessageContentCardObjectText) *MessageContentCardElementDatePicker {
	r.Placeholder = val
	return r
}

func (r *MessageContentCardElementDatePicker) SetValue(val interface{}) *MessageContentCardElementDatePicker {
	r.Value = val
	return r
}

func (r *MessageContentCardElementDatePicker) SetConfirm(val *MessageContentCardObjectConfirm) *MessageContentCardElementDatePicker {
	r.Confirm = val
	return r
}

// === MessageContentCardElementDatePicker ===

// === MessageContentCardElementImage ===

// MessageContentCardElementImage 图片模块
//
// 图片模块用于展示整张图片。建议需要着重展示的图片使用此模块，用户点击图片后可以查看大图。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUjNwUjL1YDM14SN2ATN
type MessageContentCardElementImage struct {
	ImgKey  string                        `json:"img_key,omitempty"` // 图片资源
	Alt     *MessageContentCardObjectText `json:"alt,omitempty"`     // 图片hover说明
	Preview bool                          `json:"preview,omitempty"` // 点击后是否放大图片，缺省为true。在配置 card_link 后可设置为false，使用户点击卡片上的图片也能响应card_link链接跳转
}

// IsMessageContentCardElement ...
func (r MessageContentCardElementImage) IsMessageContentCardElement() {}

// MarshalJSON ...
func (r MessageContentCardElementImage) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagImage})
}

func (r *MessageContentCardElementImage) SetAlt(val *MessageContentCardObjectText) *MessageContentCardElementImage {
	r.Alt = val
	return r
}

func (r *MessageContentCardElementImage) SetPreview(val bool) *MessageContentCardElementImage {
	r.Preview = val
	return r
}

// === MessageContentCardElementImage ===

// === MessageContentCardElementOverflow ===

// MessageContentCardElementOverflow overflow
//
// 作为overflow元素被使用，提供折叠的按钮型菜单
// overflow属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
//
// 通过options字段配置选项，可用于多个按扭的折叠隐藏功能。
type MessageContentCardElementOverflow struct {
	Options []*MessageContentCardObjectOption `json:"options,omitempty"` // 待选选项
	Value   interface{}                       `json:"value,omitempty"`   // 用户选定后返回业务方的数据
	Confirm *MessageContentCardObjectConfirm  `json:"confirm,omitempty"` // 二次确认的弹框
}

// IsMessageContentCardElement ...
func (r MessageContentCardElementOverflow) IsMessageContentCardElement() {}

// MarshalJSON ...
func (r MessageContentCardElementOverflow) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagOverflow})
}

func (r *MessageContentCardElementOverflow) SetValue(val interface{}) *MessageContentCardElementOverflow {
	r.Value = val
	return r
}

func (r *MessageContentCardElementOverflow) SetConfirm(val *MessageContentCardObjectConfirm) *MessageContentCardElementOverflow {
	r.Confirm = val
	return r
}

// === MessageContentCardElementOverflow ===

// === MessageContentCardElementSelectMenu ===

// MessageContentCardElementSelectMenu selectMenu
//
// 作为selectMenu元素被使用，提供选项菜单的功能
// selectMenu属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
//
// 选项模式（"tag":"select_static")：通过options字段配置选项，支持对多个选项进行展示供用户选择。
// 选人模式（"tag":"select_person")：通过options字段配置待选人员，无则使用当前群组作为待选人员。
//
// https://open.feishu.cn/document/ukTMukTMukTM/uIzNwUjLycDM14iM3ATN
type MessageContentCardElementSelectMenu struct {
	Tag           MessageContentCardElementTag      `json:"tag,omitempty"`            // select_static, select_person
	Placeholder   *MessageContentCardObjectText     `json:"placeholder,omitempty"`    // 占位符，无默认选项时必须有
	InitialOption string                            `json:"initial_option,omitempty"` // 默认选项的value字段值
	Options       []*MessageContentCardObjectOption `json:"options,omitempty"`        // 待选选项
	Value         map[string]interface{}            `json:"value,omitempty"`          // 用户选定后返回业务方的数据
	Confirm       *MessageContentCardObjectConfirm  `json:"confirm,omitempty"`        // 二次确认的弹框
}

// IsMessageContentCardElement ...
func (r MessageContentCardElementSelectMenu) IsMessageContentCardElement() {}

func (r *MessageContentCardElementSelectMenu) SetPlaceholder(val *MessageContentCardObjectText) *MessageContentCardElementSelectMenu {
	r.Placeholder = val
	return r
}

func (r *MessageContentCardElementSelectMenu) SetInitialOption(val string) *MessageContentCardElementSelectMenu {
	r.InitialOption = val
	return r
}

func (r *MessageContentCardElementSelectMenu) SetValue(val map[string]interface{}) *MessageContentCardElementSelectMenu {
	r.Value = val
	return r
}

func (r *MessageContentCardElementSelectMenu) SetConfirm(val *MessageContentCardObjectConfirm) *MessageContentCardElementSelectMenu {
	r.Confirm = val
	return r
}

// === MessageContentCardElementSelectMenu ===

// MessageContentCardModule 卡片正文内容
//
// 消息卡片的正文内容由模块组成，你可以“像搭积木一样自由堆砌卡片”，即堆砌模块的方式构造卡片内容。
// 卡片提供 5 类模块，你可以在 elements（或i18n_elements） 结构中自由堆砌生成需要的卡片内容，最多可堆叠50个模块。
type MessageContentCardModule interface {
	IsMessageContentCardModule()
}

// MessageContentCardModuleTag ...
type MessageContentCardModuleTag string

// MessageContentCardModuleTagDIV ...
const (
	MessageContentCardModuleTagDIV      MessageContentCardModuleTag = "div"
	MessageContentCardModuleTagMarkdown MessageContentCardModuleTag = "markdown"
	MessageContentCardModuleTagHR       MessageContentCardModuleTag = "hr"
	MessageContentCardModuleTagImage    MessageContentCardModuleTag = "img"
	MessageContentCardModuleTagAction   MessageContentCardModuleTag = "action"
	MessageContentCardModuleTagNote     MessageContentCardModuleTag = "note"
)

// === MessageContentCardModuleAction ===

// MessageContentCardModuleAction 交互模块
//
// 卡片提供 4 种交互控件（button，selectMenu，overflow，datePicker），你可以通过 actions 字段添加交互元素，实现交互功能。
// 卡片交互 有效期为30天 ，超过有效期的卡片不支持交互。
//
// https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
type MessageContentCardModuleAction struct {
	Actions []MessageContentCardElement `json:"actions,omitempty"` // 放置交互元素
	Layout  string                      `json:"layout,omitempty"`  // 交互元素布局，窄版样式默认纵向排列，使用 bisected 为二等分布局，每行两列交互元素，使用 trisection 为三等分布局，每行三列交互元素，使用 flow 为流式布局元素会按自身大小横向排列并在空间不够的时候折行
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleAction) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleAction) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagAction})
}

func (r *MessageContentCardModuleAction) SetActions(val ...MessageContentCardElement) *MessageContentCardModuleAction {
	r.Actions = val
	return r
}

func (r *MessageContentCardModuleAction) SetLayout(val string) *MessageContentCardModuleAction {
	r.Layout = val
	return r
}

// === MessageContentCardModuleAction ===

// === MessageContentCardModuleDIV ===

// MessageContentCardModuleDIV 内容模块
//
// 内容模块以文本内容为主体，同时可以选择组合图片、按钮等交互组件，实现内容混排的效果。
// 模块标签为 div , 可以单独通过 text 或 field 来展示文本内容，也可以配合一个 image 元素或一个 button, overflow, selectMenu, datePicker 等互动元素增加内容的丰富性。
//
// https://open.feishu.cn/document/ukTMukTMukTM/uMjNwUjLzYDM14yM2ATN
type MessageContentCardModuleDIV struct {
	Text   *MessageContentCardObjectText    `json:"text,omitempty"`   // 单个文本展示，和field至少要有一个
	Fields []*MessageContentCardObjectField `json:"fields,omitempty"` // 多个文本展示，和text至少要有一个
	Extra  MessageContentCardElement        `json:"extra,omitempty"`  // 附加的元素展示在文本内容右侧。 可附加的元素包括image、button、selectMenu、overflow、datePicker
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleDIV) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleDIV) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagDIV})
}

func (r *MessageContentCardModuleDIV) SetText(val *MessageContentCardObjectText) *MessageContentCardModuleDIV {
	r.Text = val
	return r
}

func (r *MessageContentCardModuleDIV) SetFields(val ...*MessageContentCardObjectField) *MessageContentCardModuleDIV {
	r.Fields = val
	return r
}

func (r *MessageContentCardModuleDIV) SetExtra(val MessageContentCardElement) *MessageContentCardModuleDIV {
	r.Extra = val
	return r
}

// === MessageContentCardModuleDIV ===

// === MessageContentCardModuleHR ===

// MessageContentCardModuleHR 分割线模块
//
// 模块之间的分割线
// 建议在内容需要明显进行分割时，使用分割线模块
//
// https://open.feishu.cn/document/ukTMukTMukTM/uQjNwUjL0YDM14CN2ATN
type MessageContentCardModuleHR struct{}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleHR) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleHR) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagHR})
}

// === MessageContentCardModuleHR ===

// === MessageContentCardModuleImage ===

// MessageContentCardModuleImage 图片模块
//
// https://open.feishu.cn/document/ukTMukTMukTM/uUjNwUjL1YDM14SN2ATN
type MessageContentCardModuleImage struct {
	ImgKey       string                        `json:"img_key,omitempty"`       // 图片资源，获取方法： https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create
	Alt          *MessageContentCardObjectText `json:"alt,omitempty"`           // hover图片时弹出的Tips文案，content取值为空时则不展示
	Title        *MessageContentCardObjectText `json:"title,omitempty"`         // 图片标题
	CustomWidth  int64                         `json:"custom_width,omitempty"`  // 自定义图片的最大展示宽度。 默认展示宽度撑满卡片的通栏图片，可在278px~580px范围内指定最大展示宽度。在飞书4.0以上版本生效
	CompactWidth bool                          `json:"compact_width,omitempty"` // 是否展示为紧凑型的图片。 默认为false，若配置为true，则展示最大宽度为278px的紧凑型图片
	Mode         string                        `json:"mode,omitempty"`          // 图片显示模式。 crop_center：居中裁剪模式，对长图会限高，并居中裁剪后展示，fit_horizontal：平铺模式，完整展示上传的图片
	Preview      bool                          `json:"preview,omitempty"`       // 点击后是否放大图片，缺省为true。在配置 card_link 后可设置为false，使用户点击卡片上的图片也能响应card_link链接跳转
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleImage) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleImage) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagImage})
}

// === MessageContentCardModuleImage ===

// === MessageContentCardModuleMarkdown ===

// MessageContentCardModuleMarkdown Markdown模块
//
// 你可以使用标准markdown语法标签，快捷地构造消息卡片内容。可使用markdown标签的内容模块包括：
//
// 使用独立的 Markdown模块，用一段markdown标签构造包含格式化文本、图片、分割线的消息卡片内容；
// 或将text对象的tag字段声明为lark_md，使markdown的格式化文本能嵌入多列文本布局、图片描述等组合使用场景
//
// 注意事项：
// 在text对象中无法使用与文本格式无关的markdown标签（比如图片、分割线）
// 目前只支持markdown语法的子集，支持的元素如下表：
//
// https://open.feishu.cn/document/ukTMukTMukTM/uMjNwUjLzYDM14yM2ATN
type MessageContentCardModuleMarkdown struct {
	Content string                                  `json:"content"` // 使用已支持的markdown语法构造markdown内容
	Href    map[string]*MessageContentCardObjectURL `json:"href"`    // key是 content 中的 url值，差异化跳转：仅在需要PC、移动端跳转不同链接使用
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleMarkdown) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleMarkdown) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagMarkdown})
}

func (r *MessageContentCardModuleMarkdown) SetHref(val map[string]*MessageContentCardObjectURL) *MessageContentCardModuleMarkdown {
	r.Href = val
	return r
}

// === MessageContentCardModuleMarkdown ===

// === MessageContentCardModuleNote ===

// MessageContentCardModuleNote 备注模块
//
// 备注模块用于展示次要信息。
// 建议使用备注模块来展示用于辅助说明或备注的次要信息，支持小尺寸的图片和文本。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ucjNwUjL3YDM14yN2ATN
type MessageContentCardModuleNote struct {
	Elements []MessageContentCardElement `json:"elements,omitempty"` // text对象或image元素
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleNote) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleNote) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagNote})
}

func (r *MessageContentCardModuleNote) SetElements(val ...MessageContentCardElement) *MessageContentCardModuleNote {
	r.Elements = val
	return r
}

// === MessageContentCardModuleNote ===

// === MessageContentCardObjectConfirm ===

// MessageContentCardObjectConfirm 用于交互元素的二次确认
// 弹框默认提供确定和取消的按钮，无需开发者手动配置。
type MessageContentCardObjectConfirm struct {
	Title *MessageContentCardObjectText `json:"title,omitempty"` // 弹框标题
	Text  *MessageContentCardObjectText `json:"text,omitempty"`  // 弹框内容
}

func (r *MessageContentCardObjectConfirm) SetTitle(val *MessageContentCardObjectText) *MessageContentCardObjectConfirm {
	r.Title = val
	return r
}

func (r *MessageContentCardObjectConfirm) SetText(val *MessageContentCardObjectText) *MessageContentCardObjectConfirm {
	r.Text = val
	return r
}

// === MessageContentCardObjectConfirm ===

// === MessageContentCardObjectField ===

// MessageContentCardObjectField field
//
// field对象可用于内容模块的field字段，通过"is_short"字段控制是否并排布局
//
// https://open.feishu.cn/document/ukTMukTMukTM/uYzNwUjL2cDM14iN3ATN
type MessageContentCardObjectField struct {
	IsShort bool                          `json:"is_short,omitempty"` // 是否并排布局
	Text    *MessageContentCardObjectText `json:"text,omitempty"`     // 	国际化文本内容
}

func (r *MessageContentCardObjectField) SetIsShort(val bool) *MessageContentCardObjectField {
	r.IsShort = val
	return r
}

func (r *MessageContentCardObjectField) SetText(val *MessageContentCardObjectText) *MessageContentCardObjectField {
	r.Text = val
	return r
}

// === MessageContentCardObjectField ===

// === MessageContentCardObjectOption ===

// MessageContentCardObjectOption 作为selectMenu的选项对象
// 作为overflow的选项对象
type MessageContentCardObjectOption struct {
	Text     *MessageContentCardObjectText `json:"text,omitempty"`      // text对象	选项显示内容，非待选人员时必填
	Value    string                        `json:"value,omitempty"`     // 选项选中后返回业务方的数据
	URL      string                        `json:"url,omitempty"`       // *仅支持overflow，跳转指定链接，和multi_url字段互斥
	MultiURL *MessageContentCardObjectURL  `json:"multi_url,omitempty"` // *仅支持overflow，跳转对应链接，和url字段互斥
}

func (r *MessageContentCardObjectOption) SetText(val *MessageContentCardObjectText) *MessageContentCardObjectOption {
	r.Text = val
	return r
}

func (r *MessageContentCardObjectOption) SetValue(val string) *MessageContentCardObjectOption {
	r.Value = val
	return r
}

func (r *MessageContentCardObjectOption) SetURL(val string) *MessageContentCardObjectOption {
	r.URL = val
	return r
}

func (r *MessageContentCardObjectOption) SetMultiURL(val *MessageContentCardObjectURL) *MessageContentCardObjectOption {
	r.MultiURL = val
	return r
}

// === MessageContentCardObjectOption ===

// === MessageContentCardObjectText ===

// MessageContentCardObjectText 作为text对象被使用，支持"plain_text"和"lark_md"两种模式
type MessageContentCardObjectText struct {
	Tag     MessageContentCardObjectTextType `json:"tag,omitempty"`     // plain_text / lark_md
	Content string                           `json:"content,omitempty"` // 文本内容
	Lines   int                              `json:"lines,omitempty"`   // 内容显示行数
}

// IsMessageContentCardElement ...
func (r MessageContentCardObjectText) IsMessageContentCardElement() {}

// MessageContentCardObjectTextType ...
type MessageContentCardObjectTextType string

// MessageContentCardTextTypePlainText ...
const (
	MessageContentCardTextTypePlainText MessageContentCardObjectTextType = "plain_text"
	MessageContentCardTextTypeLarkMd    MessageContentCardObjectTextType = "lark_md"
)

func (r *MessageContentCardObjectText) SetText(val string) *MessageContentCardObjectText {
	r.Content = val
	r.Tag = MessageContentCardTextTypePlainText
	return r
}

func (r *MessageContentCardObjectText) SetMd(val string) *MessageContentCardObjectText {
	r.Content = val
	r.Tag = MessageContentCardTextTypeLarkMd
	return r
}

func (r *MessageContentCardObjectText) SetLines(val int) *MessageContentCardObjectText {
	r.Lines = val
	return r
}

// === MessageContentCardObjectText ===

// === MessageContentCardObjectURL ===

// MessageContentCardObjectURL url对象用作多端差异跳转链接
//
// 可用于button的multi_url字段，支持按键点击的多端跳转。
// 可用于lark_md类型text对象的href字段，支持超链接点击的多端跳转。
type MessageContentCardObjectURL struct {
	URL        string `json:"url,omitempty"` // 默认跳转链接，参考注意事项
	AndroidURL string `json:"android_url,omitempty"`
	IOSURL     string `json:"ios_url,omitempty"`
	PCURL      string `json:"pc_url,omitempty"`
}

func (r *MessageContentCardObjectURL) SetURL(val string) *MessageContentCardObjectURL {
	r.URL = val
	return r
}

func (r *MessageContentCardObjectURL) SetAndroidURL(val string) *MessageContentCardObjectURL {
	r.AndroidURL = val
	return r
}

func (r *MessageContentCardObjectURL) SetIOSURL(val string) *MessageContentCardObjectURL {
	r.IOSURL = val
	return r
}

func (r *MessageContentCardObjectURL) SetPCURL(val string) *MessageContentCardObjectURL {
	r.PCURL = val
	return r
}

// === MessageContentCardObjectURL ===
