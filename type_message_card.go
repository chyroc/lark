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
	Header      *MessageContentCardHeader     `json:"header,omitempty"`        // 用于配置卡片标题内容。
	Config      *MessageContentCardConfig     `json:"config,omitempty"`        // 配置卡片属性
	Modules     []MessageContentCardModule    `json:"elements,omitempty"`      // 用于定义卡片正文内容
	I18NModules *MessageContentCardI18NModule `json:"i18n_elements,omitempty"` // 用于定义卡片正文内容，支持多语言
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
	r.I18NModules = nil
	return r
}

func (r *MessageContentCard) SetI18NModules(val *MessageContentCardI18NModule) *MessageContentCard {
	r.I18NModules = val
	r.Modules = nil
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
	Icon     *MessageContentCardElementImage  `json:"icon,omitempty"`     // 标题的前缀图标。 不填则不展示。卡片标题最多只可配1个前缀图标
	Subtitle *MessageContentCardObjectText    `json:"subtitle,omitempty"` // 配置副标题的内容。支持配置国际化文案内容; 不允许只配置副标题内容，如只配置副标题，则内容展示为主标题样式; 副标题内容最多1行，超长文案末尾使用“…”省略
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

func (r *MessageContentCardHeader) SetTitle(val *MessageContentCardObjectText) *MessageContentCardHeader {
	r.Title = val
	return r
}

func (r *MessageContentCardHeader) SetIcon(imgKey string) *MessageContentCardHeader {
	r.Icon = &MessageContentCardElementImage{
		ImgKey: imgKey,
	}
	return r
}

func (r *MessageContentCardHeader) SetSubtitle(subtitle string) *MessageContentCardHeader {
	r.Subtitle = &MessageContentCardObjectText{
		Tag:     MessageContentCardTextTypePlainText,
		Content: subtitle,
	}
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
	MessageContentCardElementTagImage            MessageContentCardElementTag = "img"
	MessageContentCardElementTagButton           MessageContentCardElementTag = "button"
	MessageContentCardElementTagOverflow         MessageContentCardElementTag = "overflow"
	MessageContentCardElementTagSelectStatic     MessageContentCardElementTag = "select_static"
	MessageContentCardElementTagSelectPerson     MessageContentCardElementTag = "select_person"
	MessageContentCardElementTagPickerDate       MessageContentCardElementTag = "date_picker"
	MessageContentCardElementTagPickerTime       MessageContentCardElementTag = "picker_time"
	MessageContentCardElementTagPickerDatetime   MessageContentCardElementTag = "picker_datetime"
	MessageContentCardElementTagImageCombination MessageContentCardElementTag = "img_combination"
)

// === MessageContentCardElementButton ===

// MessageContentCardElementButton button
//
// button 属于交互元素的一种，可用于内容块的extra字段和交互块的actions字段。
//
// https://open.feishu.cn/document/ukTMukTMukTM/uEzNwUjLxcDM14SM3ATN
type MessageContentCardElementButton struct {
	Name       string                                    `json:"name,omitempty"`        // 组件的唯一标识，用于识别用户在交互后，点击的是哪个组件
	Text       *MessageContentCardObjectText             `json:"text,omitempty"`        // 按钮中的文本
	URL        string                                    `json:"url,omitempty"`         // 跳转链接，和multi_url互斥
	MultiURL   *MessageContentCardObjectURL              `json:"multi_url,omitempty"`   // 多端跳转链接
	Type       string                                    `json:"type,omitempty"`        // 配置按钮样式，默认为"default", 可选值: "default"/"primary"/"danger"
	Value      interface{}                               `json:"value,omitempty"`       // 点击后返回业务方，	仅支持key-value形式的json结构，且key为String类型。
	Confirm    *MessageContentCardObjectConfirm          `json:"confirm,omitempty"`     // 二次确认的弹框
	ActionType MessageContentCardElementButtonActionType `json:"action_type,omitempty"` // - 指定按钮的交互类型，只在按钮上扩展此属性。 - 枚举型。枚举值包括：- link：仅链接跳转- request：仅回传交互- multi：链接跳转+回传交互同时生效- form_submit：触发表单容器的提交事件，异步提交表单容器中所有用户填写的表单内容- form_reset：触发表单容器的取消提交事件，重置所有表单组件的输入值为初始值
}

type MessageContentCardElementButtonActionType string

const (
	MessageContentCardElementButtonLink       MessageContentCardElementButtonActionType = "link"        // 仅链接跳转
	MessageContentCardElementButtonRequest    MessageContentCardElementButtonActionType = "request"     // 仅回传交互
	MessageContentCardElementButtonMulti      MessageContentCardElementButtonActionType = "multi"       // 链接跳转+回传交互同时生效
	MessageContentCardElementButtonFormSubmit MessageContentCardElementButtonActionType = "form_submit" // 触发表单容器的提交事件，异步提交表单容器中所有用户填写的表单内容
	MessageContentCardElementButtonFormReset  MessageContentCardElementButtonActionType = "form_reset"  // 触发表单容器的取消提交事件，重置所有表单组件的输入值为初始值

)

// IsMessageContentCardElement ...
func (r MessageContentCardElementButton) IsMessageContentCardElement() {}

// MarshalJSON ...
func (r MessageContentCardElementButton) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagButton})
}

func (r *MessageContentCardElementButton) SetName(val string) *MessageContentCardElementButton {
	r.Name = val
	return r
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

func (r *MessageContentCardElementButton) SetActionType(val MessageContentCardElementButtonActionType) *MessageContentCardElementButton {
	r.ActionType = val
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

type MessageContentCardI18NModule struct {
	EnUs []MessageContentCardModule `json:"en_us,omitempty"`
	ZhCn []MessageContentCardModule `json:"zh_cn,omitempty"`
	JaJp []MessageContentCardModule `json:"ja_jp,omitempty"`
}

// MessageContentCardModuleTag ...
type MessageContentCardModuleTag string

// MessageContentCardModuleTagDIV ...
const (
	MessageContentCardModuleTagDIV       MessageContentCardModuleTag = "div"
	MessageContentCardModuleTagMarkdown  MessageContentCardModuleTag = "markdown"
	MessageContentCardModuleTagHR        MessageContentCardModuleTag = "hr"
	MessageContentCardModuleTagImage     MessageContentCardModuleTag = "img"
	MessageContentCardModuleTagAction    MessageContentCardModuleTag = "action"
	MessageContentCardModuleTagNote      MessageContentCardModuleTag = "note"
	MessageContentCardModuleTagColumnSet MessageContentCardModuleTag = "column_set"
	MessageContentCardModuleTagForm      MessageContentCardModuleTag = "form"
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
	Content   string                                  `json:"content"`              // 使用已支持的markdown语法构造markdown内容
	TextAlign string                                  `json:"text_align,omitempty"` // 设置文本内容的对齐方式。枚举值包括： · left：左对齐（默认）· center：居中对齐· right：右对齐
	Href      map[string]*MessageContentCardObjectURL `json:"href"`                 // key是 content 中的 url值，差异化跳转：仅在需要PC、移动端跳转不同链接使用
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleMarkdown) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleMarkdown) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagMarkdown})
}

func (r *MessageContentCardModuleMarkdown) SetTextAlign(val string) *MessageContentCardModuleMarkdown {
	r.TextAlign = val
	return r
}

func (r *MessageContentCardModuleMarkdown) SetLeftTextAlign() *MessageContentCardModuleMarkdown {
	r.TextAlign = "left"
	return r
}

func (r *MessageContentCardModuleMarkdown) SetCenterTextAlign() *MessageContentCardModuleMarkdown {
	r.TextAlign = "center"
	return r
}

func (r *MessageContentCardModuleMarkdown) SetRightTextAlign() *MessageContentCardModuleMarkdown {
	r.TextAlign = "right"
	return r
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

// MessageContentCardModuleColumnSet 列集
type MessageContentCardModuleColumnSet struct {
	Columns           []*MessageContentCardModuleColumn `json:"columns,omitempty"`            // 存放布局子容器的数组
	FlexMode          *string                           `json:"flex_mode,omitempty"`          // 移动端窄屏幕下，各列的自适应方式。 none：默认值。移动端与PC端表现一致，不进行布局自适应，在窄屏幕下按比例压缩列宽度。stretch：在移动端和PC的窄窗口模式下，列布局变为 行布局，且每列（行）宽度强制拉伸为100%，所有列自适应为上下堆叠排布。flow: 列流式排布（自动换行），当一行展示不下下一列时，自动换至下一行展示。bisect：在移动端和PC的窄窗口模式下，两列等分布局。trisect：在移动端和PC的窄窗口模式下，三列等分布局。
	BackgroundStyle   *string                           `json:"background_style,omitempty"`   // 列集的背景色样式。 default：默认的白底样式，dark mode下为黑底。grey：灰底样式
	HorizontalSpacing *string                           `json:"horizontal_spacing,omitempty"` // 列集内，每列与上一列的左右间距，默认为default。 列集内的所有水平边距保持统一间距。default：默认间距。small：窄间距
	Action            *MessageContentCardElement        `json:"action,omitempty"`
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleColumnSet) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleColumnSet) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagColumnSet})
}

func (r *MessageContentCardModuleColumnSet) SetFlexMode(val string) *MessageContentCardModuleColumnSet {
	r.FlexMode = &val
	return r
}

func (r *MessageContentCardModuleColumnSet) SetBackgroundStyle(val string) *MessageContentCardModuleColumnSet {
	r.BackgroundStyle = &val
	return r
}

func (r *MessageContentCardModuleColumnSet) SetHorizontalSpacing(val string) *MessageContentCardModuleColumnSet {
	r.HorizontalSpacing = &val
	return r
}

func (r *MessageContentCardModuleColumnSet) SetAction(val MessageContentCardElement) *MessageContentCardModuleColumnSet {
	r.Action = &val
	return r
}

// MessageContentCardModuleColumn 列
type MessageContentCardModuleColumn struct {
	Modules       []MessageContentCardModule `json:"elements,omitempty"`       // 需要在列内展示的卡片元素
	Width         *string                    `json:"width"`                    // 列宽度属性。枚举值包括： auto：列宽度与列内元素宽度一致,weighted:列宽度按weight定义的权重分布
	Weight        *int64                     `json:"weight,omitempty"`         // 当width为weighted时生效，为当前列的宽度占比
	VerticalAlign *string                    `json:"vertical_align,omitempty"` // 列内成员垂直对齐方式。枚举值包括： top：顶对齐,center：居中对齐,bottom：底部对齐
}

// MarshalJSON ...
func (r MessageContentCardModuleColumn) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": "column"})
}

func (r *MessageContentCardModuleColumn) SetWidth(val string) *MessageContentCardModuleColumn {
	r.Width = &val
	return r
}

// SetAutoWidth 列宽度与列内元素宽度一致
func (r *MessageContentCardModuleColumn) SetAutoWidth() *MessageContentCardModuleColumn {
	val := "auto"
	r.Width = &val
	return r
}

// 列宽度按weight定义的权重分布
func (r *MessageContentCardModuleColumn) SetWeightedWidth() *MessageContentCardModuleColumn {
	val := "weighted"
	r.Width = &val
	return r
}

func (r *MessageContentCardModuleColumn) SetWeight(val int64) *MessageContentCardModuleColumn {
	width := "weighted"
	r.Width = &width
	r.Weight = &val
	return r
}

func (r *MessageContentCardModuleColumn) SetVerticalAlign(val string) *MessageContentCardModuleColumn {
	r.VerticalAlign = &val
	return r
}

// 顶对齐
func (r *MessageContentCardModuleColumn) SetTopVerticalAlign() *MessageContentCardModuleColumn {
	val := "top"
	r.VerticalAlign = &val
	return r
}

func (r *MessageContentCardModuleColumn) SetCenterVerticalAlign() *MessageContentCardModuleColumn {
	val := "center"
	r.VerticalAlign = &val
	return r
}

func (r *MessageContentCardModuleColumn) SetBottomVerticalAlign() *MessageContentCardModuleColumn {
	val := "bottom"
	r.VerticalAlign = &val
	return r
}

// === MessageContentCardModuleNote ===

// === MessageContentCardModuleForm ===

// MessageContentCardModuleForm 表单容器
type MessageContentCardModuleForm struct {
	Name     string                      `json:"name"`
	Elements []MessageContentCardElement `json:"elements"` // Form 表单的叶子节点 可内嵌布局容器、呈现组件、交互组件
}

// IsMessageContentCardModule ...
func (r MessageContentCardModuleForm) IsMessageContentCardModule() {}

// MarshalJSON ...
func (r MessageContentCardModuleForm) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagForm})
}

func (r *MessageContentCardModuleForm) SetText(val string) *MessageContentCardModuleForm {
	r.Name = val
	return r
}

func (r *MessageContentCardModuleForm) SetElements(val ...MessageContentCardElement) *MessageContentCardModuleForm {
	r.Elements = val
	return r
}

// === MessageContentCardModuleForm ===

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
	IsShort bool                          `json:"is_short"`       // 是否并排布局
	Text    *MessageContentCardObjectText `json:"text,omitempty"` // 国际化文本内容
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
	Tag     MessageContentCardObjectTextType `json:"tag"`             // plain_text / lark_md
	Content string                           `json:"content"`         // 文本内容
	Lines   int                              `json:"lines,omitempty"` // 内容显示行数
	I18N    *I18NText                        `json:"i18n,omitempty"`  // 国际化文本内容
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

// === MessageContentCardObjectInput ===

// MessageContentCardObjectInput 文本输入框组件
//
// docs: https://bytedance.feishu.cn/docx/Iy8EdadF9ohXepxGInHcdgFLn5g
type MessageContentCardObjectInput struct {
	Name          string                           `json:"name"`                     // 组件唯一标识, 默认为空。 当 input 组件嵌套在 Form 容器中时必填。指当前交互组件的唯一标识。帮助开发者在交互的回传事件中识别用户用户提交的具体是哪个交互组件。
	Placeholder   *MessageContentCardObjectText    `json:"placeholder"`              // 占位文案, 默认文案为“请输入” 指在用户未激活文本输入框时，展示的一句灰字引导文案。文本输入框的展示空间有限，建议你精简占位文案的表述。占位文案最多可输入100个字符，超出则被截断。
	MaxLength     int64                            `json:"max_length"`               // 最大文本输入长度, 默认值为1000。可取1~1000范围内的整数。 声明此字段后，可限制用户在文本输入框中填写的最大文本长度。如果终端用户录入的内容超过最大长度，组件将在文本输入框上报错提示，并拒绝提交用户输入的内容。
	DefaultValue  string                           `json:"default_value,omitempty"`  // 预填内容
	Label         *MessageContentCardObjectText    `json:"label,omitempty"`          // 文本标签
	LabelPosition string                           `json:"label_position,omitempty"` // 默认为top。枚举值仅有 top | left，填充其他值则报错。 指定文本标签展示在输入框的哪个相对位置上。在移动端等窄屏幕场景下，文本标签将自适应固定展示在输入框上方。
	Value         map[string]interface{}           `json:"value,omitempty"`          // 开发者可在交互事件中自定义的回传数据。支持回传纯字符或一个对象。
	Confirm       *MessageContentCardObjectConfirm `json:"confirm,omitempty"`        // 指在提交前是否弹出二次确认弹窗提示。只有用户点击确认后，才提交输入的内容。属性配置同 confirm 元素。 注意：input  组件嵌入在 form 容器中时，不生效单一组件绑定的 confirm元素。仅在用户点击包含提交属性的按钮时触发二次确认弹窗。
}

// IsMessageContentCardElement ...
func (r MessageContentCardObjectInput) IsMessageContentCardElement() {}

// func (r MessageContentCardObjectInput) IsMessageContentCardModule()  {}

func (r MessageContentCardObjectInput) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": "input"})
}

func (r *MessageContentCardObjectInput) SetName(val string) *MessageContentCardObjectInput {
	r.Name = val
	return r
}

func (r *MessageContentCardObjectInput) SetPlaceholder(val *MessageContentCardObjectText) *MessageContentCardObjectInput {
	r.Placeholder = val
	return r
}

func (r *MessageContentCardObjectInput) SetMaxLength(val int64) *MessageContentCardObjectInput {
	r.MaxLength = val
	return r
}

func (r *MessageContentCardObjectInput) SetDefaultValue(val string) *MessageContentCardObjectInput {
	r.DefaultValue = val
	return r
}

func (r *MessageContentCardObjectInput) SetLabel(val *MessageContentCardObjectText) *MessageContentCardObjectInput {
	r.Label = val
	return r
}

func (r *MessageContentCardObjectInput) SetLabelPosition(val string) *MessageContentCardObjectInput {
	r.LabelPosition = val
	return r
}

func (r *MessageContentCardObjectInput) SetValue(val map[string]interface{}) *MessageContentCardObjectInput {
	r.Value = val
	return r
}

func (r *MessageContentCardObjectInput) SetConfirm(val *MessageContentCardObjectConfirm) *MessageContentCardObjectInput {
	r.Confirm = val
	return r
}

// === MessageContentCardObjectInput ===

// === MessageContentCardElementImageCombination ===

// MessageContentCardElementImageCombination 图片模块
//
// 图片模块用于展示整张图片。建议需要着重展示的图片使用此模块，用户点击图片后可以查看大图。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUjNwUjL1YDM14SN2ATN
type MessageContentCardElementImageCombination struct {
	CombinationMode MessageContentCardElementCombinationMode `json:"combination_mode"`
	ImgList         []*MessageContentCardElementImage        `json:"img_list"`
}

// 多图混排的方式。枚举值包括：
// - double：双图混排，最多排布2张图。
// - triple：三图混排，最多排布3张图。
// - bisect：等分双列图混排，每行2个等大的正方形图，最多可排布6张图。
// - trisect：等分三列图混排，每行3个等大的正方形图，最多可排布9张图。
//
// 注意：
// - 如果排布的图片超过最多排布个数上限，则按图片数据顺序展示前面的图，超出部分的图不展示
// - 如果排布的图无法排满一行，则未排布的部分留白。请注意选择合适的混排方式
type MessageContentCardElementCombinationMode string

const (
	MessageContentCardElementCombinationModeDouble  MessageContentCardElementCombinationMode = "double"  // 双图混排，最多排布2张图
	MessageContentCardElementCombinationModeTriple  MessageContentCardElementCombinationMode = "triple"  // 三图混排，最多排布3张图
	MessageContentCardElementCombinationModeBisect  MessageContentCardElementCombinationMode = "bisect"  // 等分双列图混排，每行2个等大的正方形图，最多可排布6张图
	MessageContentCardElementCombinationModeTrisect MessageContentCardElementCombinationMode = "trisect" // 等分三列图混排，每行3个等大的正方形图，最多可排布9张图
)

// IsMessageContentCardElement ...
func (r MessageContentCardElementImageCombination) IsMessageContentCardElement() {}

// MarshalJSON ...
func (r MessageContentCardElementImageCombination) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagImageCombination})
}

func (r *MessageContentCardElementImageCombination) SetCombinationMode(val MessageContentCardElementCombinationMode) *MessageContentCardElementImageCombination {
	r.CombinationMode = val
	return r
}

func (r *MessageContentCardElementImageCombination) SetDouble() *MessageContentCardElementImageCombination {
	r.CombinationMode = MessageContentCardElementCombinationModeDouble
	return r
}

func (r *MessageContentCardElementImageCombination) SetTriple() *MessageContentCardElementImageCombination {
	r.CombinationMode = MessageContentCardElementCombinationModeTriple
	return r
}

func (r *MessageContentCardElementImageCombination) SetBisect() *MessageContentCardElementImageCombination {
	r.CombinationMode = MessageContentCardElementCombinationModeBisect
	return r
}

func (r *MessageContentCardElementImageCombination) SetTrisect() *MessageContentCardElementImageCombination {
	r.CombinationMode = MessageContentCardElementCombinationModeTrisect
	return r
}

func (r *MessageContentCardElementImageCombination) SetImageList(list ...string) *MessageContentCardElementImageCombination {
	res := []*MessageContentCardElementImage{}
	for _, v := range list {
		res = append(res, &MessageContentCardElementImage{ImgKey: v})
	}
	r.ImgList = res
	return r
}

// === MessageContentCardElementImageCombination ===

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

type I18NText struct {
	ZhCn string `json:"zh_cn"`
	EnUs string `json:"en_us"`
	JaJp string `json:"ja_jp"`
}
