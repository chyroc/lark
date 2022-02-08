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

import (
	"fmt"
	"strconv"
)

// docs: https://open.feishu.cn/document/ukTMukTMukTM/ukDM2YjL5AjN24SOwYjN

// DocContent 描述一篇文档的富文本内容，由标题和正文两部分组成
type DocContent struct {
	Title *DocParagraph `json:"title"`
	Body  *DocBody      `json:"body"`
}

// DocLocation 描述文档内各元素的位置
type DocLocation struct {
	ZoneID     string `json:"zoneId"`     // 编辑区域标识正文是 "0" 表格单元格是类似 "xr1m4jw7egd9nefz1s0mdsetenl5fbe3lygxc1azupv81i5t2rjmosw5ta0esuwtn8ksya" 的字符串
	StartIndex int    `json:"startIndex"` // 元素起始位置偏移量，从 0 开始，范围区间 [0, EOF)
	EndIndex   int    `json:"endIndex"`   // 元素结束位置偏移量，范围区间 [0, EOF)
}

// DocBody 文档的正文部分，由 Block 组成。
type DocBody struct {
	Blocks []*DocBlock `json:"blocks"` // 文档结构是按行排列的，每行内容是一个 Block
}

// DocBlock 文档的一行内容的基本元素。
type DocBlock struct {
	Type           string             `json:"type"`                     // 由下列类型组成，每个 Block 可以指定一种类型
	Paragraph      *DocParagraph      `json:"paragraph,omitempty"`      // 文本段落
	Gallery        *DocGallery        `json:"gallery,omitempty"`        // 图片
	File           *DocFile           `json:"file,omitempty"`           // 文件上传
	ChatGroup      *DocChatGroup      `json:"chatGroup,omitempty"`      // 群名片
	Table          *DocTable          `json:"table,omitempty"`          // 格式化表格
	HorizontalLine *DocHorizontalLine `json:"horizontalLine,omitempty"` // 水平分割线
	EmbeddedPage   *DocEmbeddedPage   `json:"embeddedPage,omitempty"`   // 内嵌网页，例如西瓜视频等
	Sheet          *DocSheet          `json:"sheet,omitempty"`          // 电子表格
	Bitable        *DocBitable        `json:"bitable,omitempty"`        // 数据表或看板
	Diagram        *DocDiagram        `json:"diagram,omitempty"`        // 绘图或uml图
	Jira           *DocJira           `json:"jira,omitempty"`           // jira filter或jira issue
	Poll           *DocPoll           `json:"poll,omitempty"`           // 投票
	Code           *DocCode           `json:"code,omitempty"`           // 新代码块
	DocsApp        *DocDocsApp        `json:"docsApp,omitempty"`        // 团队互动应用
	Callout        *DocCallout        `json:"callout,omitempty"`        // 高亮块
	UndefinedBlock *DocUndefinedBlock `json:"undefinedBlock,omitempty"` // 未支持的block，全部用undefineBlock表示
}

// DocParagraph 一行文本段落，可以由多个行内元素组成。
type DocParagraph struct {
	Style    *DocParagraphStyle     `json:"style"`    // 段落样式
	Elements []*DocParagraphElement `json:"elements"` // 段落元素组成一个段落
	Location *DocLocation           `json:"location"` // 元素位置。仅返回，不支持写入
	LineID   string                 `json:"lineId"`
}

// DocParagraphStyle 一行文本段落的样式，支持定义段落标题、列表、引用、行对齐方式。
type DocParagraphStyle struct {
	HeadingLevel int           `json:"headingLevel"` // 标题级别，支持 1-9 级标题，取值范围：[1,9]
	Collapse     bool          `json:"collapse"`     // 标题是否折叠，仅headingLevel为[1,9]时有效
	List         *DocStyleList `json:"list"`         // 有序列表/无序列表/任务列表/旧版代码块
	Quote        bool          `json:"quote"`        // 引用样式
	Link         struct {
		URL string `json:"url"`
	}
	Align string `json:"align"` //	行对齐方式 左对齐（默认）：left 右对齐：right 居中对齐：center
}

// DocStyleList 列表段落，包括有序、无序列表，任务列表，旧版代码块。
//
// 注意 如需插入代码块，请使用 Block.Code，Paragraph.List.Code 即将下线。
type DocStyleList struct {
	Type        string `json:"type"`        // 有序列表：number / 无序列表：bullet / 任务列表：checkBox / 已完成的任务列表：checkedBox / 旧版代码块：code
	IndentLevel int    `json:"indentLevel"` // 列表的缩进级别，支持指定一行的缩进  除代码块以外的列表都支持设置缩进，支持 1-16 级缩进，取值范围：[1,16]
	Number      int    `json:"number"`      // 用于指定列表的行号，仅对有序列表和代码块生效 如果为有序列表设置了缩进，行号可能会显示为字母或者罗马数字
}

// ListTag ...
func (r *DocStyleList) ListTag() string {
	switch r.Type {
	case "number":
		return strconv.FormatInt(int64(r.Number), 10) + "."
	case "bullet":
		return "-"
	case "checkBox":
		return "- [ ]"
	case "checkedBox":
		return "- [x]"
	default:
		panic(fmt.Sprintf("unknown style list " + r.Type))
	}
}

// DocParagraphElement 组成段落的行内元素，一行可以包含多个行内元素。
type DocParagraphElement struct {
	Type             string               `json:"type"`             // 由下列类型组成，每个元素可以指定一种类型
	TextRun          *DocTextRun          `json:"textRun"`          // 文本
	DocsLink         *DocDocsLink         `json:"docsLink"`         // 文档链接，可以根据链接自动识别为标题
	Person           *DocPerson           `json:"person"`           // @用户
	Equation         *DocEquation         `json:"equation"`         // LaTeX 公式
	Reminder         *DocReminder         `json:"reminder"`         // 任务列表截止时间，当且仅当这行是 checkBox 类型时才会生效
	File             *DocFile             `json:"file"`             // 文件的行内展示样式，会展示文件的名称
	Jira             *DocJira             `json:"jira"`             // 行内的jira issue
	UndefinedElement *DocUndefinedElement `json:"undefinedElement"` // 暂未支持的行内元素，不支持写入
}

// DocTextRun 段落中的文本内容。
type DocTextRun struct {
	Text     string        `json:"text"`     // 具体的文本内容
	Style    *DocTextStyle `json:"style"`    // 文本内容的样式，支持 BIUS、颜色等
	Location *DocLocation  `json:"location"` //	仅返回，不支持写入。描述元素在当前区域（由 zoneId 指定）的位置
	LineID   string        `json:"lineId"`   // 仅返回，不支持写入。对于段落样式为 Heading 的，可以通过 lineId 拼接出段落的跳转链接
}

// DocTextStyle 文本样式
type DocTextStyle struct {
	Bold          bool         `json:"bold"`          // 加粗
	Italic        bool         `json:"italic"`        // 斜体
	StrikeThrough bool         `json:"strikeThrough"` // 删除线
	Underline     bool         `json:"underline"`     // 下划
	CodeInline    bool         `json:"codeInline"`    // 行内代码样式
	BackColor     *DocRGBColor `json:"backColor"`     // 背景色
	TextColor     *DocRGBColor `json:"textColor"`     // 字体颜色
	Link          *DocLink     `json:"link"`          // 超链接
}

// DocRGBColor 支持为文本指定颜色，会根据线上文本支持的颜色进行校验，不支持色板以外的颜色。
//
// 支持的字体颜色：rgb(216, 57, 49), rgb(222, 120, 2), rgb(220, 155, 4), rgb(46, 161, 33), rgb(36, 91, 219), rgb(36, 91, 219), rgb(100, 37, 208), rgb(143, 149, 158)
//
// 支持的背景颜色：rgb(251, 191, 188), rgb(254, 212, 164), rgb(255, 246, 122), rgb(183, 237, 177), rgb(186, 206, 253), rgb(205, 178, 250), rgb(222, 224, 227), rgb(247, 105, 100), rgb(255, 165, 61), rgb(255, 233, 40), rgb(98, 210, 86), rgb(78, 131, 253), rgb(147, 90, 246), rgb(187, 191, 196)
type DocRGBColor struct {
	Red   int     `json:"red"`   //	取值范围：[0,255]
	Green int     `json:"green"` //	取值范围：[0,255]
	Blue  int     `json:"blue"`  //	取值范围：[0,255]
	Alpha float32 `json:"alpha"` //	取值范围：[0,1]
}

// DocLink 超链接
type DocLink struct {
	URL string `json:"url"` // 超链接的地址，读写都必须使用query escape编码。比如https%3A%2F%2Fwww.baidu.com%2Fmaps%2Fembed%2Fv1%2Fplace%3Fkey%3DAIzaSyAfJZc8JxNRe909WC_QBILdlM55NqGnI30%26q%3DCentral%2BPark%26center%3D40.7828647%2C-73.9675438%26zoom%3D17z%26language%3Den
}

// DocDocsLink 文档链接，需要传入云文档的链接，支持所有云文档类型。有效链接示例：https://sample.feishu.cn/docs/doccnByZP6puODElAYySJkPIfUb
type DocDocsLink struct {
	URL      string       `json:"url"`      // 云文档链接
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocPerson 一个人
type DocPerson struct {
	OpenID   string       `json:"openId"`   //	用户的 openId，比如ou_3bbe8a09c20e89cce9bff989ed840674。
	Location *DocLocation `json:"location"` //	元素位置。仅返回，不支持写入
}

// DocReminder 日期提醒，当且仅当所在行样式是 checkBox 时有效。
type DocReminder struct {
	IsWholeDay   bool         `json:"isWholeDay"`   // 日期还是整点小时
	Timestamp    int          `json:"timestamp"`    // 如果设置时间，要求30分或整点；如果不设置IsWholeDay true，时间戳为H:59:59，H和时区相关
	ShouldNotify bool         `json:"shouldNotify"` // 是否通知
	NotifyType   int          `json:"notifyType"`   // 通知触发时机: isWholeDay 为false时： 0 现在1 提前5分钟2 提前15分钟3 提前30分钟4 提前1小时5 提前2小时6 提前1天7 提前2天 / isWholeDay 为true时：50 当天9点51 提前一天9点52 提前两天9点53 提前一周9点
	Location     *DocLocation `json:"location"`     // 元素位置。仅返回，不支持写入
}

// DocEquation LaTeX 公式
type DocEquation struct {
	Equation string       `json:"equation"` // LaTeX 公式，需要遵循 LaTeX 的语法，比如"E=mc^2"
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocUndefinedElement 未支持的行内元素
type DocUndefinedElement struct {
	Location *DocLocation `json:"location"`
}

// DocGallery 图片
type DocGallery struct {
	GalleryStyle *DocGalleryStyle `json:"galleryStyle"` // 图片样式，目前仅支持指定对齐方式
	ImageList    []*DocImageItem  `json:"imageList"`    // 图片对象，如果同时传入多张，会显示为同行多图。一行内最多显示 6 张图片
	Location     *DocLocation     `json:"location"`     // 元素位置。仅返回，不支持写入
}

// DocGalleryStyle 图片样式，目前仅支持指定对齐方式。
type DocGalleryStyle struct {
	Align string `json:"align"` // 图片对齐方式，仅当一行内只有一张图片时生效。 居中（默认）：center 左对齐：left 右对齐：right
}

// DocImageItem 单次写操作图片和附件最多 50 个，单个图片最大 20M。本地图片需要先通过上传素材或分片上传素材进行上传，支持jpg、jpeg、bmp、png 和 gif 格式。
//
// 支持指定图片宽高，为了保证图片显示效果，会根据图片原始宽高比例进行校验计算，最终的显示效果可能与手动指定的宽高有一定区别。
type DocImageItem struct {
	FileToken string `json:"fileToken"` // 图片 token，比如boxcnOj88GDkmWGm2zsTyCBqoLb，不支持编辑
	Width     int    `json:"width"`     // 图片宽，单位px
	Height    int    `json:"height"`    // 图片高，单位px
}

// DocFile 单次写操作图片和附件最多 50 个。本地文件需要先通过上传素材或分片上传素材进行上传。
type DocFile struct {
	FileToken string       `json:"fileToken"` // 附件 token，比如boxcnOj88GDkmWGm2zsTyCBqoLb，不支持编辑
	ViewType  string       `json:"viewType"`  // 文件展示样式。 预览：preview 卡片：card 行内：inline
	FileName  string       `json:"fileName"`  // 文件名，不支持编辑
	Location  *DocLocation `json:"location"`  // 元素位置。仅返回，不支持写入
}

// DocHorizontalLine 水平分割线。
type DocHorizontalLine struct {
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocEmbeddedPage 内嵌的网页。
type DocEmbeddedPage struct {
	Type     string       `json:"type"`     // 支持以下网页：  bilibili："bilibili" 西瓜视频："xigua" 优酷："youku" Airtable："airtable" 百度地图："baidumap"
	Url      string       `json:"url"`      // 第三方网页链接，读写都必须使用query escape编码。
	Width    float32      `json:"width"`    // 窗口宽度，如未指定使用默认值，如果提供会使用默认宽高等比计算高度
	Height   float32      `json:"height"`   // 窗口高度，如未指定使用默认值
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocChatGroup 群
type DocChatGroup struct {
	OpenChatID string       `json:"openChatId"` // 群聊天会话openId，比如oc_4149593da6ef5fe4de16b10cb4769c94 文档拷贝副本后，对于没有权限的群名片会替换为"none" 对于写操作，如果使用"none"或者用户不在该群都会返回无权限错误
	Location   *DocLocation `json:"location"`   // 元素位置。仅返回，不支持写入
}

// DocTable 表格
//
// 文档内普通表格，如要了解文档内的数据表格或多维表格，请查看下方 Sheet、Bitable 部分。
// 创建支持两种方式，二选一：
// 提供rowSize、columnSize、tableRows、tableStyle、mergedCells 带内容创建
// 指定rowSize和columnSize创建空表格
type DocTable struct {
	TableID     string           `json:"tableId"`     // 表格ID，不可编辑
	RowSize     int              `json:"rowSize"`     // 表格行数量，创建空表格时最大值 9
	ColumnSize  int              `json:"columnSize"`  // 表格列数量，创建空表格时最大值 9
	TableRows   []*DocTableRow   `json:"tableRows"`   // 表格行内容，带内容创建时必须和 rowSize、columnSize 一致
	TableStyle  *DocTableStyle   `json:"tableStyle"`  // 表格样式
	MergedCells []*DocMergedCell `json:"mergedCells"` // 合并单元格信息
	Location    *DocLocation     `json:"location"`    // 元素位置。仅返回，不支持写入
}

// DocTableRow 普通表格的一行。
type DocTableRow struct {
	RowIndex   int             `json:"rowIndex"`   // 行索引，从 0 开始，第一行是 0，不可编辑
	TableCells []*DocTableCell `json:"tableCells"` // 表格单元格内容
}

// DocTableCell 普通表格的一个单元格。
type DocTableCell struct {
	ZoneID      string      `json:"zoneId"`      // 单元格ID
	ColumnIndex int         `json:"columnIndex"` //	列索引，从 0 开始，第一列是 0，不可编辑
	Body        interface{} `json:"body"`        // object	单元格内容，支持 Table、Sheet、Bitable 以外的 Block
}

// DocTableStyle 表格样式，目前仅支持调整表格的列宽。
type DocTableStyle struct {
	TableColumnProperties []*DocTableColumnProperty `json:"tableColumnProperties"` // 列属性
}

// DocTableColumnProperty 普通表格的列宽。
type DocTableColumnProperty struct {
	Width int `json:"width"` // 列宽，单位 px，列宽最小 50，最大 1300，默认 100
}

// DocMergedCell 合并的单元格。
type DocMergedCell struct {
	MergedCellID     string `json:"mergedCellId"`     // 合并单元格 id，不可编辑
	RowStartIndex    int    `json:"rowStartIndex"`    // 合并单元格行起始索引，从 0 开始
	RowEndIndex      int    `json:"rowEndIndex"`      // 合并单元格行截止索引
	ColumnStartIndex int    `json:"columnStartIndex"` // 合并单元格列起始索引，从 0 开始
	ColumnEndIndex   int    `json:"columnEndIndex"`   // 合并单元格列截止索引
}

// DocSheet 文档内的数据表格。数据表格内容读写API 创建支持两种方式，二选一：
//
// 通过token深拷贝，token只允许doc内电子表格，不允许独立电子表格
// 指定rowSize和columnSize创建空sheet
type DocSheet struct {
	Token      string       `json:"token"`      // 	sheet token，比如shtcnKjLimyn2fW0uGseasQYGgh_m7fMrN，其中，下划线前的是sheet token，下划线后的部分是tableId，使用时请注意拆分。在创建时，如果传入 token，则进行深拷贝。
	RowSize    int          `json:"rowSize"`    // 电子表格列数量。创建空电子表格时使用，最大值为 9
	ColumnSize int          `json:"columnSize"` // 电子表格列数量。创建空电子表格时使用，最大值为 9
	Location   *DocLocation `json:"location"`   // 元素位置。仅返回，不支持写入
}

// DocBitable 文档内的多维表格，创建支持两种方式，二选一：
//
// 通过 token+viewType 深拷贝，token只允许doc内电子表格，不允许独立电子表格
// 不带 token，通过 viewType 创建空 bitable
type DocBitable struct {
	Token    string       `json:"token"`    // bitable token，比如shtcnKjLimyn2fW0uGseasQYGgh_m7fMrN创建接口或插入操作使用，如果非空，进行深拷贝。不可编辑
	ViewType string       `json:"viewType"` // 数据表：grid  看板：kanban
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocDiagram 绘图，包含流程图和 UML 图。目前不支持写入。
type DocDiagram struct {
	Token       string       `json:"token"`       // diagram token, 比如diacnK1MYEHBopBbIdc6A5AOVCh，不支持编辑
	DiagramType string       `json:"diagramType"` // 绘图类型。  流程图：flowchart UML图：uml
	Location    *DocLocation `json:"location"`    // 元素位置。仅返回，不支持写入
}

// DocJira Jira，包括 Jira filter 和 Jira issue。目前不支持写入
type DocJira struct {
	Token    string       `json:"token"`    // jira token，比如 jftcnsA0fY8V3CzYvtRPy9XsXxf，不支持编辑
	JiraType string       `json:"jiraType"` // Jira 类型。  过滤器：filter 问题：issue
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocPoll 投票。目前不支持写入。
type DocPoll struct {
	Token    string       `json:"token"`    // poll token，比如 jsncnxuT7beirSpf33NfcKrSwAh，不支持编辑
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}

// DocCode 代码
type DocCode struct {
	Language    string       `json:"language"`    // 代码语言 支持以下语言:Plain Text, ABAP, Ada, Apache, Apex, Assembly language, Bash, C, C#, C++, COBOL, CSS, CoffeeScript, D, Dart, Delphi, Dockerfile, Django, Erlang, Fortran, FoxPro, Go, Groovy, HTML, HTMLBars, HTTP, Haskell, JSON, Java, JavaScript, Julia, Kotlin, LaTeX, Lisp, Logo, Lua, MATLAB, Makefile, Markdown, Nginx, Objective-C, OpenEdge ABL, PHP, Perl, PostScript, PowerShell, Prolog, ProtoBuf, Python, R, RPG, Ruby, Rust, SAS, SCSS, SQL, Scala, Scheme, Scratch, Shell, Swift, Thrift, TypeScript, VBScript, Visual Basic, XML, YAML
	WrapContent bool         `json:"wrapContent"` // 是否自动换行
	Body        *DocBody     `json:"body"`        // 代码块内容（只支持Paragraph block)
	ZoneID      string       `json:"zoneId"`      // Read only. 代码块的ID，对代码块内容进行增加修改删除时需要
	Location    *DocLocation `json:"location"`    // 元素位置。仅返回，不支持写入
}

// DocCallout 高亮
type DocCallout struct {
	CalloutEmojiID         string       `json:"calloutEmojiId"`         // 高亮块表情支持以下：https://bytedance.feishu.cn/sheets/shtcnkDBCFZJDyGliM7IOqPuSgd?sheet=7FyRq7
	CalloutBackgroundColor *DocRGBColor `json:"calloutBackgroundColor"` // 高亮块背景色（分为深色系和浅色系，与边框色深浅色系对应）
	CalloutBorderColor     *DocRGBColor `json:"calloutBorderColor"`     // 高亮块边框色（分为深色系和浅色系，与背景色深浅色系对应）
	CalloutTextColor       *DocRGBColor `json:"calloutTextColor"`       // 高亮块内文字颜色
	Body                   *DocBody     `json:"body"`                   // 高亮块内容（仅支持Paragraph Block和HorizontalLine Block，ParagraphStyle支持除Collapse以外的所有）
	ZoneID                 string       `json:"zoneId"`                 // Read only. 高亮块的ID，对代码块内容进行增加修改删除时需要
	Location               *DocLocation `json:"location"`               // Read only. 高亮块位置
}

// DocDocsApp 团队互动应用。目前不支持写入。
type DocDocsApp struct {
	TypeID     string       `json:"typeId"`     // 团队互动应用类型，比如信息收集"blk_5f992038c64240015d280958"
	InstanceId string       `json:"instanceId"` // 团队互动应用唯一ID
	Location   *DocLocation `json:"location"`   // 元素位置。仅返回，不支持写入
}

// DocUndefinedBlock 未支持的 Block 类型
type DocUndefinedBlock struct {
	Location *DocLocation `json:"location"` // 元素位置。仅返回，不支持写入
}
