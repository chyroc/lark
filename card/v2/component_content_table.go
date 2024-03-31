package card

import (
	"encoding/json"
	"strconv"
)

func Table(columns []*TableColumn, rows ...map[string]interface{}) *ComponentTable {
	return &ComponentTable{
		Columns: columns,
		Rows:    rows,
	}
}

// ComponentTable 表格
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/table
//
// 飞书卡片支持表格组件, 并支持在表格中添加普通文本、选项标签、人员列表以及数字格式的内容。
//
// 注意事项
// 表格组件仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
// 表格组件支持飞书 V7.4 及以上版本的客户端。在低于该版本的飞书客户端上, 表格的内容将展示为一句“请升级客户端为最新版本后查看内容”的占位图。
//
// 嵌套规则
// 表格组件不可被内嵌在其它组件内, 只可放在卡片根节点下。
// 表格组件不支持内嵌其它组件。
//
//go:generate generate_set_attrs -type=ComponentTable
//go:generate generate_to_map -type=ComponentTable
type ComponentTable struct {
	componentBase

	// 每页最大展示的数据行数。支持 [1,10] 整数。
	PageSize int64 `json:"page_size,omitempty"`

	// 表格的行高。单元格高度如无法展示一整行内容, 则上下裁剪内容。可取值:
	//
	// low: 低
	// middle: 中
	// high: 高
	// [32,124]px: 自定义行高, 单位为像素, 如 40px。取值范围是 [32,124]
	RowHeight RowHeight `json:"row_height,omitempty"`

	// 表头样式风格。详见下方 header_style 字段说明。
	HeaderStyle *TableHeaderStyle `json:"header_style,omitempty"`

	// 列对象数组。详见下方 column 字段说明。
	Columns []*TableColumn `json:"columns,omitempty"`

	// 行对象数组。与列定义对应的数据。用 "name":VALUE 的形式, 定义每一行的数据内容。name即你自定义的列标记。
	Rows []map[string]any `json:"rows,omitempty"`
}

type RowHeight = string

const (
	RowHeightLow    RowHeight = "low"    // 低
	RowHeightMiddle RowHeight = "middle" // 中
	RowHeightHigh   RowHeight = "high"   // 高
)

type TableHeaderStyle struct {
	// 表头文本对齐方式。可取值：
	//
	// left：左对齐
	// center：居中对齐
	// right：右对齐
	TextAlign TextAlign `json:"text_align,omitempty"`

	// 表头文本大小。可取值：
	//
	// normal：正文（14px）
	// heading：标题（16px）
	TextSize TextSize `json:"text_size,omitempty"`

	// 表头背景色。可取值：
	//
	// grey：灰色
	// none：无背景色
	BackgroundStyle Color `json:"background_style,omitempty"`

	// 文本颜色。可取值：
	//
	// default：客户端浅色主题模式下为黑色; 客户端深色主题模式下为白色
	// grey：灰色
	TextColor Color `json:"text_color,omitempty"`

	// 表头文本是否加粗。可取值：
	//
	// true：加粗
	// false：不加粗
	Bold bool `json:"bold,omitempty"`

	// 表头文本的行数。支持大于等于 1 的整数或 full_displayed 字符串：
	//
	// full_displayed：不限制表头行数, 完整展示文本
	// 大于等于 1 的整数
	Lines TableHeaderStyleLine `json:"lines,omitempty"`
}

type TableHeaderStyleLine int64

func (r TableHeaderStyleLine) MarshalJSON() ([]byte, error) {
	if r <= 0 {
		return []byte("full_displayed"), nil
	}
	return []byte(strconv.FormatInt(int64(r), 10)), nil
}

type TableColumn struct {
	// 自定义列的标记。用于唯一指定行数据对象数组中, 需要将数据填充至这一行的具体哪个单元格中。
	Name string `json:"name,omitempty"`

	// 在表头展示的列名称。不填或为空则不展示列名称。
	DisplayName string `json:"display_name,omitempty"`

	// 列宽度。可取值：
	//
	// auto：自适应内容宽度
	// 自定义宽度：自定义表格的列宽度, 如 120px。取值范围是 [80px,600px] 的整数
	// 自定义宽度百分比：自定义列宽度占当前表格画布宽度的百分比（表格画布宽度 = 卡片宽度-卡片左右内边距）, 如 25%。取值范围是 [1%,100%]
	Width Width `json:"width,omitempty"`

	// 列内数据对齐方式。可选值：
	//
	// left：左对齐
	// center：居中对齐
	// right：右对齐
	TextAlign TextAlign `json:"text_align,omitempty"`

	// 列数据类型。可选值：
	//
	// text：不带格式的普通文本
	// options：选项标签
	// number：数字。默认在单元格中右对齐展示。若选择该数据类型, 你可继续在 column 中添加 format 字段, 设置数字的子格式属性
	DataType string `json:"data_type,omitempty"`

	// 该字段仅当 data_type 为 number 时生效, 你可以在该字段内选择设置小数点位数、货币单位以及千分位样式。
	Format *TableColumnFormat `json:"format,omitempty"`
}

// 该字段仅当 data_type 为 number 时生效, 你可以在该字段内选择设置小数点位数、货币单位以及千分位样式。
type TableColumnFormat struct {
	// 数字的小数点位数。默认不限制小数点位数, 原样透传展示开发者输入的数字。
	// 可填 0~10 的整数。小数点位数为 0 表示取整数。
	Precision int64 `json:"precision,omitempty"`

	// 数字前的货币单位。不填或为空不展示。可填 1 个字符的货币单位文本, 如 “¥”。
	Symbol string `json:"symbol,omitempty"`

	// 是否生效按千分位逗号分割的数字样式。
	Seperator bool `json:"seperator,omitempty"`
}

// MarshalJSON ...
func (r ComponentTable) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "table"
	return json.Marshal(m)
}

// SetPageSize set ComponentTable.PageSize attribute
func (r *ComponentTable) SetPageSize(val int64) *ComponentTable {
	r.PageSize = val
	return r
}

// SetRowHeight set ComponentTable.RowHeight attribute
func (r *ComponentTable) SetRowHeight(val RowHeight) *ComponentTable {
	r.RowHeight = val
	return r
}

// SetRowHeightLow set ComponentTable.RowHeight attribute to RowHeightLow
func (r *ComponentTable) SetRowHeightLow() *ComponentTable {
	r.RowHeight = RowHeightLow
	return r
}

// SetRowHeightMiddle set ComponentTable.RowHeight attribute to RowHeightMiddle
func (r *ComponentTable) SetRowHeightMiddle() *ComponentTable {
	r.RowHeight = RowHeightMiddle
	return r
}

// SetRowHeightHigh set ComponentTable.RowHeight attribute to RowHeightHigh
func (r *ComponentTable) SetRowHeightHigh() *ComponentTable {
	r.RowHeight = RowHeightHigh
	return r
}

// SetHeaderStyle set ComponentTable.HeaderStyle attribute
func (r *ComponentTable) SetHeaderStyle(val *TableHeaderStyle) *ComponentTable {
	r.HeaderStyle = val
	return r
}

// SetColumns set ComponentTable.Columns attribute
func (r *ComponentTable) SetColumns(val ...*TableColumn) *ComponentTable {
	r.Columns = val
	return r
}

// SetRows set ComponentTable.Rows attribute
func (r *ComponentTable) SetRows(val ...map[string]any) *ComponentTable {
	r.Rows = val
	return r
}

// toMap conv ComponentTable to map
func (r *ComponentTable) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 5)
	if r.PageSize != 0 {
		res["page_size"] = r.PageSize
	}
	if r.RowHeight != "" {
		res["row_height"] = r.RowHeight
	}
	if r.HeaderStyle != nil {
		res["header_style"] = r.HeaderStyle
	}
	if len(r.Columns) != 0 {
		res["columns"] = r.Columns
	}
	if len(r.Rows) != 0 {
		res["rows"] = r.Rows
	}
	return res
}
