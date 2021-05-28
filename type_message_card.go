package lark

import (
	"encoding/json"
)

type MessageContentCard struct {
	Header  *MessageContentCardHeader  `json:"header,omitempty"`
	Config  *MessageContentCardConfig  `json:"config,omitempty"`
	Modules []MessageContentCardModule `json:"elements,omitempty"`
}

func (r MessageContentCard) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

type MessageContentCardHeader struct {
	Template MessageContentCardHeaderTemplate `json:"template,omitempty"`
	Title    *MessageContentCardObjectText    `json:"title,omitempty"`
}

type MessageContentCardConfig struct {
	EnableForward bool `json:"enable_forward,omitempty"` // 是否允许卡片被转发，默认 true，转发后，卡片上的“回传交互”组件将自动置为禁用态。用户不能在转发后的卡片操作提交数据
}

type MessageContentCardHeaderTemplate string

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

type MessageContentCardModule interface {
	IsMessageContentCardModule()
}

type MessageContentCardModuleDIV struct {
	Text   *MessageContentCardObjectText    `json:"text,omitempty"`   // 单个文本展示，和field至少要有一个
	Fields []*MessageContentCardObjectField `json:"fields,omitempty"` // 多个文本展示，和text至少要有一个
	Extra  MessageContentCardElement        `json:"extra,omitempty"`  // 展示附加元素，最多可展示一个元素
}

func (r MessageContentCardModuleDIV) IsMessageContentCardModule() {}

func (r MessageContentCardModuleDIV) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagDIV})
}

type MessageContentCardModuleHR struct{}

func (r MessageContentCardModuleHR) IsMessageContentCardModule() {}

func (r MessageContentCardModuleHR) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagHR})
}

type MessageContentCardModuleImage struct {
	ImgKey string                        `json:"img_key,omitempty"` // 图片资源，获取方法：上传图片
	Alt    *MessageContentCardObjectText `json:"alt,omitempty"`     // 图片hover时展示说明，为空则不展示
	Title  *MessageContentCardObjectText `json:"title,omitempty"`   // 图片标题
	Mode   string                        `json:"mode,omitempty"`    // 图片显示模式。
}

func (r MessageContentCardModuleImage) IsMessageContentCardModule() {}

func (r MessageContentCardModuleImage) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagImage})
}

type MessageContentCardModuleAction struct {
	Actions []MessageContentCardElement `json:"actions,omitempty"` // 放置交互元素
	Layout  string                      `json:"layout,omitempty"`  // 交互元素布局，窄版样式默认纵向排列，使用 bisected 为二等分布局，每行两列交互元素，使用 trisection 为三等分布局，每行三列交互元素，使用 flow 为流式布局元素会按自身大小横向排列并在空间不够的时候折行
}

func (r MessageContentCardModuleAction) IsMessageContentCardModule() {}

func (r MessageContentCardModuleAction) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagAction})
}

type MessageContentCardModuleNote struct {
	Elements []MessageContentCardElement `json:"elements,omitempty"` // text对象或image元素
}

func (r MessageContentCardModuleNote) IsMessageContentCardModule() {}

func (r MessageContentCardModuleNote) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardModuleTagNote})
}

type MessageContentCardModuleTag string

const (
	MessageContentCardModuleTagDIV    MessageContentCardModuleTag = "div"
	MessageContentCardModuleTagHR     MessageContentCardModuleTag = "hr"
	MessageContentCardModuleTagImage  MessageContentCardModuleTag = "img"
	MessageContentCardModuleTagAction MessageContentCardModuleTag = "action"
	MessageContentCardModuleTagNote   MessageContentCardModuleTag = "note"
)

type MessageContentCardElement interface {
	IsMessageContentCardElement()
}

type MessageContentCardElementImage struct {
	ImgKey  string                        `json:"img_key,omitempty"` // 图片资源
	Alt     *MessageContentCardObjectText `json:"alt,omitempty"`     // 图片hover说明
	Preview bool                          `json:"preview,omitempty"` // 点击后是否放大图片，缺省为true。在配置 card_link 后可设置为false，使用户点击卡片上的图片也能响应card_link链接跳转
}

func (r MessageContentCardElementImage) IsMessageContentCardElement() {}

func (r MessageContentCardElementImage) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagImage})
}

type MessageContentCardElementButton struct {
	Text     *MessageContentCardObjectText    `json:"text,omitempty"`      // 按钮中的文本
	URL      string                           `json:"url,omitempty"`       // 跳转链接，和multi_url互斥
	MultiURL *MessageContentCardObjectURL     `json:"multi_url,omitempty"` // 多端跳转链接
	Type     string                           `json:"type,omitempty"`      // 配置按钮样式，默认为"default"
	Value    interface{}                      `json:"value,omitempty"`     // 点击后返回业务方
	Confirm  *MessageContentCardObjectConfirm `json:"confirm,omitempty"`   // 二次确认的弹框
}

func (r MessageContentCardElementButton) IsMessageContentCardElement() {}

func (r MessageContentCardElementButton) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagButton})
}

type MessageContentCardElementSelectMenu struct {
	Tag           MessageContentCardElementTag      `json:"tag,omitempty"`            // select_static, select_person
	Placeholder   *MessageContentCardObjectText     `json:"placeholder,omitempty"`    // 占位符，无默认选项时必须有
	InitialOption string                            `json:"initial_option,omitempty"` // 默认选项的value字段值
	Options       []*MessageContentCardObjectOption `json:"options,omitempty"`        // 待选选项
	Value         map[string]interface{}            `json:"value,omitempty"`          // 用户选定后返回业务方的数据
	Confirm       *MessageContentCardObjectConfirm  `json:"confirm,omitempty"`        // 二次确认的弹框
}

func (r MessageContentCardElementSelectMenu) IsMessageContentCardElement() {}

type MessageContentCardElementOverflow struct {
	Options []*MessageContentCardObjectOption `json:"options,omitempty"` // 待选选项
	Value   interface{}                       `json:"value,omitempty"`   // 用户选定后返回业务方的数据
	Confirm *MessageContentCardObjectConfirm  `json:"confirm,omitempty"` // 二次确认的弹框
}

func (r MessageContentCardElementOverflow) IsMessageContentCardElement() {}

func (r MessageContentCardElementOverflow) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"tag": MessageContentCardElementTagOverflow})
}

type MessageContentCardElementDatePicker struct {
	Tag             MessageContentCardElementTag     `json:"tag,omitempty"`              // 如下三种取值 ：date_picker, picker_time, picker_datetime
	InitialDate     string                           `json:"initial_date,omitempty"`     // 格式"YYYY-MM-DD"	日期模式的初始值
	InitialTime     string                           `json:"initial_time,omitempty"`     // 格式"HH:mm"	时间模式的初始值
	InitialDatetime string                           `json:"initial_datetime,omitempty"` // 格式"YYYY-MM-DD HH:mm"	日期模式的初始值
	Placeholder     *MessageContentCardObjectText    `json:"placeholder,omitempty"`      // 占位符，无初始值时必填
	Value           interface{}                      `json:"value,omitempty"`            // 用户选定后返回业务方的数据
	Confirm         *MessageContentCardObjectConfirm `json:"confirm,omitempty"`          // 二次确认的弹框

}

func (r MessageContentCardElementDatePicker) IsMessageContentCardElement() {}

type MessageContentCardElementTag string

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

// 作为text对象被使用，支持"plain_text"和"lark_md"两种模式
type MessageContentCardObjectText struct {
	Tag     MessageContentCardObjectTextType `json:"tag,omitempty"`     // plain_text / lark_md
	Content string                           `json:"content,omitempty"` // 文本内容
	Lines   int                              `json:"lines,omitempty"`   // 内容显示行数
}

type MessageContentCardObjectTextType string

const (
	MessageContentCardTextTypePlainText   MessageContentCardObjectTextType = "plain_text"
	MessageContentCardTextTypePlainLarkMd MessageContentCardObjectTextType = "lark_md"
)

// field对象可用于内容模块的field字段，通过"is_short"字段控制是否并排布局
type MessageContentCardObjectField struct {
	IsShort bool                          `json:"is_short,omitempty"` // 是否并排布局
	Text    *MessageContentCardObjectText `json:"text,omitempty"`     // 	国际化文本内容
}

// url对象用作多端差异跳转链接
//
// 可用于button的multi_url字段，支持按键点击的多端跳转。
// 可用于lark_md类型text对象的href字段，支持超链接点击的多端跳转。
type MessageContentCardObjectURL struct {
	URL        string `json:"url,omitempty"` // 默认跳转链接，参考注意事项
	AndroidURL string `json:"android_url,omitempty"`
	IOSURL     string `json:"ios_url,omitempty"`
	PCURL      string `json:"pc_url,omitempty"`
}

// 作为selectMenu的选项对象
// 作为overflow的选项对象
type MessageContentCardObjectOption struct {
	Text     *MessageContentCardObjectText `json:"text,omitempty"`      // text对象	选项显示内容，非待选人员时必填
	Value    string                        `json:"value,omitempty"`     // 选项选中后返回业务方的数据
	URL      string                        `json:"url,omitempty"`       // *仅支持overflow，跳转指定链接，和multi_url字段互斥
	MultiURL *MessageContentCardObjectURL  `json:"multi_url,omitempty"` // *仅支持overflow，跳转对应链接，和url字段互斥
}

// 用于交互元素的二次确认
// 弹框默认提供确定和取消的按钮，无需开发者手动配置。
type MessageContentCardObjectConfirm struct {
	Title *MessageContentCardObjectText `json:"title,omitempty"` // 弹框标题
	Text  *MessageContentCardObjectText `json:"text,omitempty"`  // 弹框内容
}
