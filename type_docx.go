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

// docs: https://bytedance.feishu.cn/docx/doxcnP2DBm4Rwd5SdnivQX1iSBf

// DocxDocument 表示一篇新版文档。
type DocxDocument struct {
	DocumentID string `json:"document_id,omitempty"`
	RevisionID int64  `json:"revision_id,omitempty"`
	Title      string `json:"title,omitempty"`
}

// DocxBlock ...
type DocxBlock struct {
	BlockID   string        `json:"block_id,omitempty"`   // Block 唯一标识
	ParentID  string        `json:"parent_id,omitempty"`  // block 的父亲 id
	Children  []string      `json:"children,omitempty"`   // block 的孩子 id 列表
	BlockType DocxBlockType `json:"block_type,omitempty"` // block 类型, 可选值有: `1`：页面 Block, `2`：文本 Block, `3`：标题 1 Block, `4`：标题 2 Block, `5`：标题 3 Block, `6`：标题 4 Block, `7`：标题 5 Block, `8`：标题 6 Block, `9`：标题 7 Block, `10`：标题 8 Block, `11`：标题 9 Block, `12`：无序列表 Block, `13`：有序列表 Block, `14`：代码块 Block, `15`：引用 Block, `16`：公式 Block 暂不支持, `17`：待办事项 Block, `18`：多维表格 Block, `19`：高亮块 Block, `20`：会话卡片 Block, `21`：流程图 & UML Block, `22`：分割线 Block, `23`：文件 Block, `24`：分栏 Block, `25`：分栏列 Block, `26`：内嵌 Block Block, `27`：图片 Block, `28`：开放平台小组件 Block, `29`：思维笔记 Block, `30`：电子表格 Block, `31`：表格 Block, `32`：表格单元格 Block, `33`：视图 Block, `999`：未支持 Block
	// BlockData 只能是以下其中一种，并且需与 BlockType 相对应:
	Page           *DocxBlockText          `json:"page,omitempty"`            // 文档 Block
	Text           *DocxBlockText          `json:"text,omitempty"`            // 文本 Block
	Heading1       *DocxBlockText          `json:"heading1,omitempty"`        // 一级标题 Block
	Heading2       *DocxBlockText          `json:"heading2,omitempty"`        // 二级标题 Block
	Heading3       *DocxBlockText          `json:"heading3,omitempty"`        // 三级标题 Block
	Heading4       *DocxBlockText          `json:"heading4,omitempty"`        // 四级标题 Block
	Heading5       *DocxBlockText          `json:"heading5,omitempty"`        // 五级标题 Block
	Heading6       *DocxBlockText          `json:"heading6,omitempty"`        // 六级标题 Block
	Heading7       *DocxBlockText          `json:"heading7,omitempty"`        // 七级标题 Block
	Heading8       *DocxBlockText          `json:"heading8,omitempty"`        // 八级标题 Block
	Heading9       *DocxBlockText          `json:"heading9,omitempty"`        // 九级标题 Block
	Bullet         *DocxBlockText          `json:"bullet,omitempty"`          // 无序列表 Block
	Ordered        *DocxBlockText          `json:"ordered,omitempty"`         // 有序列表 Block
	Code           *DocxBlockText          `json:"code,omitempty"`            // 代码块 Block
	Quote          *DocxBlockText          `json:"quote,omitempty"`           // 引用 Block
	Equation       *DocxBlockText          `json:"equation,omitempty"`        // 公式 Block
	Todo           *DocxBlockText          `json:"todo,omitempty"`            // 任务 Block
	Bitable        *DocxBlockBitable       `json:"bitable,omitempty"`         // 多维表格 Block
	Callout        *DocxBlockCallout       `json:"callout,omitempty"`         // 高亮块 Block
	ChatCard       *DocxBlockChatCard      `json:"chat_card,omitempty"`       // 群聊卡片 Block
	Diagram        *DocxBlockDiagram       `json:"diagram,omitempty"`         // 流程图/UML Block
	Divider        DocxBlockDivider        `json:"divider,omitempty"`         // 分割线 Block
	File           *DocxBlockFile          `json:"file,omitempty"`            // 文件 Block
	Grid           *DocxBlockGrid          `json:"grid,omitempty"`            // 分栏 Block
	GridColumn     *DocxBlockGridColumn    `json:"grid_column,omitempty"`     // 分栏列 Block
	Iframe         *DocxBlockIframe        `json:"iframe,omitempty"`          // 内嵌 Block
	Image          *DocxBlockImage         `json:"image,omitempty"`           // 图片 Block
	ISV            *DocxBlockISV           `json:"isv,omitempty"`             // 三方 Block
	Mindnote       *DocxBlockMindnote      `json:"mindnote,omitempty"`        // 思维笔记 Block
	Sheet          *DocxBlockSheet         `json:"sheet,omitempty"`           // 电子表格 Block
	Table          *DocxBlockTable         `json:"table,omitempty"`           // 表格 Block
	TableCell      *DocxBlockTableCell     `json:"table_cell,omitempty"`      // 单元格 Block
	View           *DocxBlockView          `json:"view,omitempty"`            // 视图 Block
	OKR            *DocxBlockOKR           `json:"okr,omitempty"`             // OKR Block
	OKRObjective   *DocxBlockOKRObjective  `json:"okr_objective,omitempty"`   // OKR Objective Block
	OKRKeyResult   *DocxBlockOKRKeyResult  `json:"okr_key_result,omitempty"`  // OKR KR Block
	Undefined      *DocxBlocUndefined      `json:"undefined,omitempty"`       // 未支持 Block
	QuoteContainer *DocxBlocQuoteContainer `json:"quote_container,omitempty"` // 引用容器 Block
}

// DocxBlockBitable ...
//
// 多维表格 Block。目前支持通过指定 view_type 创建空 Bitable。
type DocxBlockBitable struct {
	Token    string              `json:"token,omitempty"`     // 多维表格文档 Token
	ViewType DocxBitableViewType `json:"view_type,omitempty"` // 视图类型, 1:数据表, 2:看板
}

// DocxBlockCallout ...
//
// 高亮块 Block
type DocxBlockCallout struct {
	BackgroundColor DocxCalloutBackgroundColor `json:"background_color,omitempty"` // 高亮块背景色
	BorderColor     DocxCalloutBorderColor     `json:"border_color,omitempty"`     // 边框色
	TextColor       DocxFontColor              `json:"text_color,omitempty"`       // 文字颜色
	EmojiID         string                     `json:"emoji_id,omitempty"`         // 高亮块图标
}

// DocxBlockChatCard ...
//
// 群聊卡片 Block
type DocxBlockChatCard struct {
	ChatID string    `json:"chat_id,omitempty"` // 群聊天会话 OpenID，其值以 ‘oc_’ 开头，表示专为开放能力接口定义的群组 ID。对于写操作，如果用户不在该群则返回无权限错误。
	Align  DocxAlign `json:"align,omitempty"`   // 对齐方式1：居左排版, 2：居中排版, 3：居右排版
}

// DocxBlockDiagram ...
//
// 流程图/UML Block
type DocxBlockDiagram struct {
	DiagramType DocxDiagramType `json:"diagram_type,omitempty"` // 绘图类型1：流程图, 2：UML 图
}

// DocxBlockDivider ...
//
// 分割线 Block
type DocxBlockDivider struct{}

// DocxBlockFile 文件 Block
//
// 文件 Block。文件 Block 不能独立存在，须与视图 Block 一同出现，并且文件视图是通过控制视图 Block 的 view_type 实现的，比如卡片视图、预览视图等。在创建文件 Block 时，系统会自动生成默认的视图 Block。
type DocxBlockFile struct {
	Token string `json:"token,omitempty"` // 附件 Token
	Name  string `json:"name,omitempty"`  // 文件名
}

// DocxBlockGrid ...
//
// 分栏 Block
type DocxBlockGrid struct {
	ColumnSize int64 `json:"column_size,omitempty"` // 分栏列数量，取值范围 [2,5)
}

// DocxBlockGridColumn ...
//
// 分栏列 Block
type DocxBlockGridColumn struct {
	WidthRatio int64 `json:"width_ratio,omitempty"` // 当前分栏列占整个分栏的比例
}

// DocxBlockIframe ...
//
// 内嵌 Block
type DocxBlockIframe struct {
	Component *DocxBlockIframeComponent `json:"component,omitempty"` // iframe 的组成元素
}

// DocxBlockIframeComponent ...
//
// iframe 的组成元素
type DocxBlockIframeComponent struct {
	IframeType DocxIframeComponentType `json:"iframe_type,omitempty"` // iframe 类型, 可选值有：1：哔哩哔哩, 2：西瓜视频, 3：优酷, 4：Airtable, 5：百度地图, 6：高德地图, 7：TikTok, 8：Figma, 9：墨刀, 10：Canva, 11：CodePen, 12：飞书问卷, 13：金数据, 14：谷歌地图, 15：Youtube, 99：Other
	URL        string                  `json:"url,omitempty"`         // iframe 目标 url（需要进行 url_encode）
}

// DocxBlockImage ...
//
// 图片 Block
type DocxBlockImage struct {
	Width  int64  `json:"width,omitempty"`  // 宽度单位 px
	Height int64  `json:"height,omitempty"` // 高度
	Token  string `json:"token,omitempty"`  // 图片 Token
}

// DocxBlockISV 三方 Block
type DocxBlockISV struct {
	ComponentID     string `json:"component_id,omitempty"`      // 团队互动应用唯一ID
	ComponentTypeID string `json:"component_type_id,omitempty"` // 团队互动应用类型，比如信息收集"blk_5f992038c64240015d280958"
}

// DocxBlockMindnote 思维笔记 Block
//
// 目前只支持查询返回占位信息，不支持创建及编辑。
type DocxBlockMindnote struct {
	Token string `json:"token,omitempty"` // 思维导图 token
}

// DocxBlockSheet 电子表格 Block
//
// 目前只指定 row_size 和 column_size 创建空 Sheet。
type DocxBlockSheet struct {
	Token      string `json:"token,omitempty"`       // 电子表格 token。创建接口使用，如果非空则进行深拷贝
	RowSize    int64  `json:"row_size,omitempty"`    // 电子表格列数量。创建空电子表格时使用，最大值 9
	ColumnSize int64  `json:"column_size,omitempty"` // 电子表格列数量。创建空电子表格时使用，最大值 9
}

// DocxBlockTable 表格 Block
type DocxBlockTable struct {
	Cells    []string                `json:"cells,omitempty"`    // 单元格数组，数组元素为 Table Cell Block 的 ID
	Property *DocxBlockTableProperty `json:"property,omitempty"` // 表格属性
}

// DocxBlockTableProperty 表格属性
type DocxBlockTableProperty struct {
	RowSize     int64                              `json:"row_size,omitempty"`     // 行数
	ColumnSize  int64                              `json:"column_size,omitempty"`  // 列数
	ColumnWidth []int64                            `json:"column_width,omitempty"` // 列宽，单位px
	MergeInfo   []*DocxBlockTablePropertyMergeInfo `json:"merge_info,omitempty"`   // 单元格合并信息
}

// DocxBlockTablePropertyMergeInfo 单元格合并信息
type DocxBlockTablePropertyMergeInfo struct {
	RowSpan int64 `json:"row_span,omitempty"` // 从当前行索引起被合并的连续行数
	ColSpan int64 `json:"col_span,omitempty"` // 从当前列索引起被合并的连续列数
}

// DocxBlockTableCell 单元格 Block
type DocxBlockTableCell struct{}

// DocxBlockView 视图 Block
type DocxBlockView struct {
	ViewType DocxViewType `json:"view_type,omitempty"` // 视图类型, 可选值有：1：卡片视图, 2：预览视图, 3：内联试图
}

// DocxBlocUndefined 未支持 Block
type DocxBlocUndefined struct{}

// DocxBlocQuoteContainer 引用容器 Block
type DocxBlocQuoteContainer struct{}

// DocxBlockOKR OKR 信息
type DocxBlockOKR struct {
	OKRID               string                      `json:"okr_id,omitempty"`                // OKR ID
	PeriodDisplayStatus *DocxOKRPeriodDisplayStatus `json:"period_display_status,omitempty"` // 周期的状态
	PeriodNameEn        string                      `json:"period_name_en,omitempty"`        // 周期名 - 英文
	PeriodNameZh        string                      `json:"period_name_zh,omitempty"`        // 周期名 - 中文
	UserID              string                      `json:"user_id,omitempty"`               // OKR 所属的用户 ID
	VisibleSetting      *DocxOKRVisibleSetting      `json:"visible_setting,omitempty"`       // 可见性设置
}

// DocxBlockOKRObjective OKR Objective 信息
type DocxBlockOKRObjective struct {
	Confidential bool                 `json:"confidential,omitempty"`  // 是否在 OKR 平台设置了私密权限
	Content      *DocxBlockText       `json:"content,omitempty"`       // Objective 的文本内容
	ObjectiveID  string               `json:"objective_id,omitempty"`  // OKR Objective ID
	Position     int64                `json:"position,omitempty"`      // Objective 的位置编号，对应 Block 中 O1、O2 的 1、2
	ProgressRate *DocxOKRProgressRate `json:"progress_rate,omitempty"` // 进展信息
	Score        int64                `json:"score,omitempty"`         // 打分信息
	Visible      bool                 `json:"visible,omitempty"`       // OKR Block 中是否展示该 Objective
	Weight       int64                `json:"weight,omitempty"`        // Objective 的权重
}

// DocxBlockOKRKeyResult OKR KR 信息
type DocxBlockOKRKeyResult struct {
	Confidential bool                 `json:"confidential,omitempty"`  // 是否在 OKR 平台设置了私密权限
	Content      *DocxBlockText       `json:"content,omitempty"`       // Key Result 的文本内容
	KeyResultID  string               `json:"kr_id,omitempty"`         // OKR Key Result ID
	Position     int64                `json:"position,omitempty"`      // Key Result 的位置编号，对应 Block 中 KR1、KR2 的 1、2
	ProgressRate *DocxOKRProgressRate `json:"progress_rate,omitempty"` // 进展信息
	Score        int64                `json:"score,omitempty"`         // 打分信息
	Visible      bool                 `json:"visible,omitempty"`       // OKR Block 中是否展示该 Key Result
	Weight       float32              `json:"weight,omitempty"`        // Key Result 的权重
}

// DocxBlockText 文本 Block，其有多种 type。
type DocxBlockText struct {
	Style    *DocxTextStyle     `json:"style,omitempty"`    // 文本样式
	Elements []*DocxTextElement `json:"elements,omitempty"` // 文本元素
}

// DocxTextStyle ...
type DocxTextStyle struct {
	Align    DocxAlign        `json:"align,omitempty"`    // 对齐方式, 可选值有: `1`：居左排版, `2`：居中排版, `3`：居右排版
	Done     bool             `json:"done,omitempty"`     // 任务的完成状态
	Folded   bool             `json:"folded,omitempty"`   // 文本的折叠状态
	Language DocxCodeLanguage `json:"language,omitempty"` // 代码块语言, 可选值有: `1`：PlainText, `2`：ABAP, `3`：Ada, `4`：Apache, `5`：Apex, `6`：Assembly Language, `7`：Bash, `8`：CSharp, `9`：C++, `10`：C, `11`：COBOL, `12`：CSS, `13`：CoffeeScript, `14`：D, `15`：Dart, `16`：Delphi, `17`：Django, `18`：Dockerfile, `19`：Erlang, `20`：Fortran, `21`：FoxPro, `22`：Go, `23`：Groovy, `24`：HTML, `25`：HTMLBars, `26`：HTTP, `27`：Haskell, `28`：JSON, `29`：Java, `30`：JavaScript, `31`：Julia, `32`：Kotlin, `33`：LateX, `34`：Lisp, `35`：Logo, `36`：Lua, `37`：MATLAB, `38`：Makefile, `39`：Markdown, `40`：Nginx, `41`：Objective-C, `42`：OpenEdgeABL, `43`：PHP, `44`：Perl, `45`：PostScript, `46`：Power Shell, `47`：Prolog, `48`：ProtoBuf, `49`：Python, `50`：R, `51`：RPG, `52`：Ruby, `53`：Rust, `54`：SAS, `55`：SCSS, `56`：SQL, `57`：Scala, `58`：Scheme, `59`：Scratch, `60`：Shell, `61`：Swift, `62`：Thrift, `63`：TypeScript, `64`：VBScript, `65`：Visual Basic, `66`：XML, `67`：YAML
	Wrap     bool             `json:"wrap,omitempty"`     // 代码块是否自动换行
}

// DocxTextElement 文本元素，支持多种类型。
type DocxTextElement struct {
	TextRun     *DocxTextElementTextRun     `json:"text_run,omitempty"`     // 文字
	MentionUser *DocxTextElementMentionUser `json:"mention_user,omitempty"` // @用户
	MentionDoc  *DocxTextElementMentionDoc  `json:"mention_doc,omitempty"`  // @文档
	Reminder    *DocxTextElementReminder    `json:"reminder,omitempty"`     // 日期提醒
	File        *DocxTextElementInlineFile  `json:"file,omitempty"`         // 内联附件
	Undefined   *DocxTextElementUndefined   `json:"undefined,omitempty"`    // 未支持的 TextElement
	Equation    *DocxTextElementEquation    `json:"equation,omitempty"`     // 公式
}

// DocxTextElementStyle ...
type DocxTextElementStyle struct {
	Bold            bool                      `json:"bold,omitempty"`             // 加粗
	Italic          bool                      `json:"italic,omitempty"`           // 斜体
	Strikethrough   bool                      `json:"strikethrough,omitempty"`    // 删除线
	Underline       bool                      `json:"underline,omitempty"`        // 下划线
	InlineCode      bool                      `json:"inline_code,omitempty"`      // inline 代码
	BackgroundColor DocxFontBackgroundColor   `json:"background_color,omitempty"` // 背景色, 可选值有: `1`：浅粉红色, `2`：浅橙色, `3`：浅黄色, `4`：浅绿色, `5`：浅蓝色, `6`：浅紫色, `7`：浅灰色, `8`：暗粉红色, `9`：暗橙色, `10`：暗黄色, `11`：暗绿色, `12`：暗蓝色, `13`：暗紫色, `14`：暗灰色, `15`：暗银灰色
	TextColor       DocxFontColor             `json:"text_color,omitempty"`       // 字体颜色, 可选值有: `1`：粉红色, `2`：橙色, `3`：黄色, `4`：绿色, `5`：蓝色, `6`：紫色, `7`：灰色
	Link            *DocxTextElementStyleLink `json:"link,omitempty"`             // 链接
}

// DocxTextElementStyleLink ...
type DocxTextElementStyleLink struct {
	URL string `json:"url,omitempty"` // 超链接指向的 url (需要 url_encode)
}

// DocxTextElementTextRun ...
type DocxTextElementTextRun struct {
	Content          string                `json:"content,omitempty"`            // 文本内容
	TextElementStyle *DocxTextElementStyle `json:"text_element_style,omitempty"` // 文本局部样式
}

// DocxTextElementMentionUser ...
type DocxTextElementMentionUser struct {
	UserID string `json:"user_id,omitempty"` // 用户 OpenID
}

// DocxTextElementMentionDoc ...
type DocxTextElementMentionDoc struct {
	Token   string             `json:"token,omitempty"`    // 云文档 token
	ObjType DocxMentionObjType `json:"obj_type,omitempty"` // 云文档类型, 可选值有: `1`：Doc, `3`：Sheet, `8`：Bitable, `11`：MindNote, `12`：File, `15`：Slide, `16`：Wiki, `22`：Docx
	URL     string             `json:"url,omitempty"`      // 云文档链接（需要 url_encode)
	Title   string             `json:"title,omitempty"`    // 云文档标题，只读属性
}

// DocxTextElementReminder ...
type DocxTextElementReminder struct {
	CreateUserID string `json:"create_user_id,omitempty"` // 创建者用户 ID
	IsNotify     bool   `json:"is_notify,omitempty"`      // 是否通知
	IsWholeDay   bool   `json:"is_whole_day,omitempty"`   // 是日期还是整点小时
	ExpireTime   string `json:"expire_time,omitempty"`    // 事件发生的时间（毫秒级事件戳）
	NotifyTime   string `json:"notify_time,omitempty"`    // 触发通知的时间（毫秒级时间戳）
}

// DocxTextElementInlineFile ...
//
// 仅输出
type DocxTextElementInlineFile struct {
	FileToken     string `json:"file_token,omitempty"`      // 附件 token
	SourceBlockID string `json:"source_block_id,omitempty"` // 当前文档中该附件所处的 block 的 id
}

// DocxTextElementUndefined ...
//
// 仅输出
type DocxTextElementUndefined struct{}

// DocxTextElementEquation 公式
type DocxTextElementEquation struct {
	Content string `json:"content,omitempty"` // 符合 KaTeX 语法的公式内容。语法规则请参考：https://katex.org/docs/supported.html
}

// DocxOKRVisibleSetting OKR 可见设置
type DocxOKRVisibleSetting struct {
	ProgressFillAreaVisible bool `json:"progress_fill_area_visible,omitempty"`
	ProgressStatusVisible   bool `json:"progress_status_visible,omitempty"`
	ScoreVisible            bool `json:"score_visible,omitempty"`
}

// DocxOKRProgressRate OKR 进度
type DocxOKRProgressRate struct {
	Current        int64                     `json:"current,omitempty"`
	Mode           DocxOKRProgressRateMode   `json:"mode,omitempty"`
	Percent        int64                     `json:"percent,omitempty"`
	ProgressStatus DocxOKRProgressStatus     `json:"progress_status,omitempty"`
	Start          int64                     `json:"start,omitempty"`
	StatusType     DocxOKRProgressStatusType `json:"status_type,omitempty"`
	Target         int64                     `json:"target,omitempty"`
}

// 枚举类型

// DocxBlockType ...
type DocxBlockType int64

const (
	DocxBlockTypePage           DocxBlockType = 1   // 文档 Block
	DocxBlockTypeText           DocxBlockType = 2   // 文本 Block
	DocxBlockTypeHeading1       DocxBlockType = 3   // 一级标题 Block
	DocxBlockTypeHeading2       DocxBlockType = 4   // 二级标题 Block
	DocxBlockTypeHeading3       DocxBlockType = 5   // 三级标题 Block
	DocxBlockTypeHeading4       DocxBlockType = 6   // 四级标题 Block
	DocxBlockTypeHeading5       DocxBlockType = 7   // 五级标题 Block
	DocxBlockTypeHeading6       DocxBlockType = 8   // 六级标题 Block
	DocxBlockTypeHeading7       DocxBlockType = 9   // 七级标题 Block
	DocxBlockTypeHeading8       DocxBlockType = 10  // 八级标题 Block
	DocxBlockTypeHeading9       DocxBlockType = 11  // 九级标题 Block
	DocxBlockTypeBullet         DocxBlockType = 12  // 无序列表 Block
	DocxBlockTypeOrdered        DocxBlockType = 13  // 有序列表 Block
	DocxBlockTypeCode           DocxBlockType = 14  // 代码块 Block
	DocxBlockTypeQuote          DocxBlockType = 15  // 引用 Block
	DocxBlockTypeEquation       DocxBlockType = 16  // 公式 Block
	DocxBlockTypeTodo           DocxBlockType = 17  // 任务 Block
	DocxBlockTypeBitable        DocxBlockType = 18  // 多维表格 Block
	DocxBlockTypeCallout        DocxBlockType = 19  // 高亮块 Block
	DocxBlockTypeChatCard       DocxBlockType = 20  // 群聊卡片 Block
	DocxBlockTypeDiagram        DocxBlockType = 21  // 流程图/UML Block
	DocxBlockTypeDivider        DocxBlockType = 22  // 分割线 Block
	DocxBlockTypeFile           DocxBlockType = 23  // 文件 Block
	DocxBlockTypeGrid           DocxBlockType = 24  // 分栏 Block
	DocxBlockTypeGridColumn     DocxBlockType = 25  // 分栏列 Block
	DocxBlockTypeIframe         DocxBlockType = 26  // 内嵌 Block
	DocxBlockTypeImage          DocxBlockType = 27  // 图片 Block
	DocxBlockTypeISV            DocxBlockType = 28  // 三方 Block
	DocxBlockTypeMindnote       DocxBlockType = 29  // 思维笔记 Block
	DocxBlockTypeSheet          DocxBlockType = 30  // 电子表格 Block
	DocxBlockTypeTable          DocxBlockType = 31  // 表格 Block
	DocxBlockTypeTableCell      DocxBlockType = 32  // 单元格 Block
	DocxBlockTypeView           DocxBlockType = 33  // 视图 Block
	DocxBlockTypeQuoteContainer DocxBlockType = 34  // 引用容器 Block
	DocxBlockTypeTask           DocxBlockType = 35  // 任务容器 Block
	DocxBlockTypeOKR            DocxBlockType = 36  // OKR容器 Block
	DocxBlockTypeOKRObjective   DocxBlockType = 37  // OKR Objective容器 Block
	DocxBlockTypeOKRKeyResult   DocxBlockType = 38  // OKR KeyResult容器 Block
	DocxBlockTypeProgress       DocxBlockType = 39  // Progress容器 Block
	DocxBlockTypeUndefined      DocxBlockType = 999 // 未支持 Block
)

// DocxTextElementType ...
type DocxTextElementType string

const (
	DocxTextElementTypeTextRun     DocxTextElementType = "text_run"     // 文字
	DocxTextElementTypeMentionUser DocxTextElementType = "mention_user" // @用户
	DocxTextElementTypeMentionDoc  DocxTextElementType = "mention_doc"  // @文档
	DocxTextElementTypeFile        DocxTextElementType = "file"         // @文件
	DocxTextElementTypeReminder    DocxTextElementType = "reminder"     // 日期提醒
	DocxTextElementTypeUndefined   DocxTextElementType = "undefined"    // 未支持元素
	DocxTextElementTypeEquation    DocxTextElementType = "equation"     // 公式
)

// DocxAlign Block 的排版方式，比如居左等
type DocxAlign int64

const (
	DocxAlignLeft   DocxAlign = 1 // 居左排版
	DocxAlignCenter DocxAlign = 2 // 居中排版
	DocxAlignRight  DocxAlign = 3 // 居右排版
)

// DocxCalloutBackgroundColor 高亮块背景色（分为深色系和浅色系）
type DocxCalloutBackgroundColor int64

const (
	DocxCalloutBackgroundColorLightRed    DocxCalloutBackgroundColor = 1  // 浅红色
	DocxCalloutBackgroundColorLightOrange DocxCalloutBackgroundColor = 2  // 浅橙色
	DocxCalloutBackgroundColorLightYellow DocxCalloutBackgroundColor = 3  // 浅黄色
	DocxCalloutBackgroundColorLightGreen  DocxCalloutBackgroundColor = 4  // 浅绿色
	DocxCalloutBackgroundColorLightBlue   DocxCalloutBackgroundColor = 5  // 浅蓝色
	DocxCalloutBackgroundColorLightPurple DocxCalloutBackgroundColor = 6  // 浅紫色
	DocxCalloutBackgroundColorLightGrey   DocxCalloutBackgroundColor = 7  // 浅灰色
	DocxCalloutBackgroundColorDarkRed     DocxCalloutBackgroundColor = 8  // 暗红色
	DocxCalloutBackgroundColorDarkOrange  DocxCalloutBackgroundColor = 9  // 暗橙色
	DocxCalloutBackgroundColorDarkYellow  DocxCalloutBackgroundColor = 10 // 暗黄色
	DocxCalloutBackgroundColorDarkGreen   DocxCalloutBackgroundColor = 11 // 暗绿色
	DocxCalloutBackgroundColorDarkBlue    DocxCalloutBackgroundColor = 12 // 暗蓝色
	DocxCalloutBackgroundColorDarkPurple  DocxCalloutBackgroundColor = 13 // 暗紫色
	DocxCalloutBackgroundColorDarkGrey    DocxCalloutBackgroundColor = 14 // 暗灰色
)

// DocxCalloutBorderColor 高亮块边框色（色系与高亮块背景色色系一致）
type DocxCalloutBorderColor int64

const (
	DocxCalloutBorderColorRed    DocxCalloutBorderColor = 1 // 红色
	DocxCalloutBorderColorOrange DocxCalloutBorderColor = 2 // 橙色
	DocxCalloutBorderColorYellow DocxCalloutBorderColor = 3 // 黄色
	DocxCalloutBorderColorGreen  DocxCalloutBorderColor = 4 // 绿色
	DocxCalloutBorderColorBlue   DocxCalloutBorderColor = 5 // 蓝色
	DocxCalloutBorderColorPurple DocxCalloutBorderColor = 6 // 紫色
	DocxCalloutBorderColorGrey   DocxCalloutBorderColor = 7 // 灰色
)

// DocxFontBackgroundColor 字体背景色（分为深色系和浅色系）
type DocxFontBackgroundColor int64

const (
	DocxFontBackgroundColorLightPink   DocxFontBackgroundColor = 1  // 浅粉红色
	DocxFontBackgroundColorLightOrange DocxFontBackgroundColor = 2  // 浅橙色
	DocxFontBackgroundColorLightYellow DocxFontBackgroundColor = 3  // 浅黄色
	DocxFontBackgroundColorLightGreen  DocxFontBackgroundColor = 4  // 浅绿色
	DocxFontBackgroundColorLightBlue   DocxFontBackgroundColor = 5  // 浅蓝色
	DocxFontBackgroundColorLightPurple DocxFontBackgroundColor = 6  // 浅紫色
	DocxFontBackgroundColorLightGrey   DocxFontBackgroundColor = 7  // 浅灰色
	DocxFontBackgroundColorDarkPink    DocxFontBackgroundColor = 8  // 暗粉红色
	DocxFontBackgroundColorDarkOrange  DocxFontBackgroundColor = 9  // 暗橙色
	DocxFontBackgroundColorDarkYellow  DocxFontBackgroundColor = 10 // 暗黄色
	DocxFontBackgroundColorDarkGreen   DocxFontBackgroundColor = 11 // 暗绿色
	DocxFontBackgroundColorDarkBlue    DocxFontBackgroundColor = 12 // 暗蓝色
	DocxFontBackgroundColorDarkPurple  DocxFontBackgroundColor = 13 // 暗紫色
	DocxFontBackgroundColorDarkGrey    DocxFontBackgroundColor = 14 // 暗灰色
)

// FontColor 字体色
type DocxFontColor int64

const (
	DocxFontColorLightPink   DocxFontColor = 1 // 浅粉红色
	DocxFontColorLightOrange DocxFontColor = 2 // 浅橙色
	DocxFontColorLightYellow DocxFontColor = 3 // 浅黄色
	DocxFontColorLightGreen  DocxFontColor = 4 // 浅绿色
	DocxFontColorLightBlue   DocxFontColor = 5 // 浅蓝色
	DocxFontColorLightPurple DocxFontColor = 6 // 浅紫色
	DocxFontColorLightGrey   DocxFontColor = 7 // 浅灰色
)

// ViewType View Block 的视图类型
type DocxViewType int64

const (
	DocxViewTypeCard    DocxViewType = 1 // 卡片视图，独占一行的一种视图，在 Card 上可有一些简单交互
	DocxViewTypePreview DocxViewType = 2 // 预览视图，在当前页面直接预览插入的 Block 内容，而不需要打开新的页面
	DocxViewTypeInline  DocxViewType = 3 // 内联试图
)

// DocxBitableViewType Bitable Block 的视图类型
type DocxBitableViewType int64

const (
	DocxBitableViewTypeTable  DocxBitableViewType = 1 // 数据表
	DocxBitableViewTypeKanban DocxBitableViewType = 2 // 看板
)

// DocxDiagramType 绘图类型
type DocxDiagramType int64

const DocxDiagramTypeUML DocxDiagramType = 2 // uml

// 内嵌 Block 支持的类型
type DocxIframeComponentType int64

const (
	DocxIframeComponentTypeBilibili      DocxIframeComponentType = 1  // 哔哩哔哩
	DocxIframeComponentTypeXigua         DocxIframeComponentType = 2  // 西瓜视频
	DocxIframeComponentTypeYouku         DocxIframeComponentType = 3  // 优酷
	DocxIframeComponentTypeAirtable      DocxIframeComponentType = 4  // Airtable
	DocxIframeComponentTypeBaiduMap      DocxIframeComponentType = 5  // 百度地图
	DocxIframeComponentTypeGaodeMap      DocxIframeComponentType = 6  // 高德地图
	DocxIframeComponentTypeTikTok        DocxIframeComponentType = 7  // TikTok
	DocxIframeComponentTypeFigma         DocxIframeComponentType = 8  // Figma
	DocxIframeComponentTypeModao         DocxIframeComponentType = 9  // 墨刀
	DocxIframeComponentTypeCanva         DocxIframeComponentType = 10 // Canva
	DocxIframeComponentTypeCodePen       DocxIframeComponentType = 11 // CodePen
	DocxIframeComponentTypeFeishuWenjuan DocxIframeComponentType = 12 // 飞书问卷
	DocxIframeComponentTypeJinshuju      DocxIframeComponentType = 13 // 金数据
	DocxIframeComponentTypeGoogleMap     DocxIframeComponentType = 14 // 谷歌地图
	DocxIframeComponentTypeYoutube       DocxIframeComponentType = 15 // Youtube
)

// DocxMentionObjType Mention
type DocxMentionObjType int64

const (
	DocxMentionObjTypeDoc      DocxMentionObjType = 1  // Doc
	DocxMentionObjTypeSheet    DocxMentionObjType = 3  // Sheet
	DocxMentionObjTypeBitable  DocxMentionObjType = 8  // Bitable
	DocxMentionObjTypeMindNote DocxMentionObjType = 11 // MindNote
	DocxMentionObjTypeFile     DocxMentionObjType = 12 // File
	DocxMentionObjTypeSlide    DocxMentionObjType = 15 // Slide
	DocxMentionObjTypeWiki     DocxMentionObjType = 16 // Wiki
	DocxMentionObjTypeDocx     DocxMentionObjType = 22 // Docx
)

// DocxCodeLanguage 代码块语言
type DocxCodeLanguage int64

const (
	DocxCodeLanguagePlainText    DocxCodeLanguage = 1  // PlainText
	DocxCodeLanguageABAP         DocxCodeLanguage = 2  // ABAP
	DocxCodeLanguageAda          DocxCodeLanguage = 3  // Ada
	DocxCodeLanguageApache       DocxCodeLanguage = 4  // Apache
	DocxCodeLanguageApex         DocxCodeLanguage = 5  // Apex
	DocxCodeLanguageAssembly     DocxCodeLanguage = 6  // Assembly
	DocxCodeLanguageBash         DocxCodeLanguage = 7  // Bash
	DocxCodeLanguageCSharp       DocxCodeLanguage = 8  // CSharp
	DocxCodeLanguageCPlusPlus    DocxCodeLanguage = 9  // C++
	DocxCodeLanguageC            DocxCodeLanguage = 10 // C
	DocxCodeLanguageCOBOL        DocxCodeLanguage = 11 // COBOL
	DocxCodeLanguageCSS          DocxCodeLanguage = 12 // CSS
	DocxCodeLanguageCoffeeScript DocxCodeLanguage = 13 // CoffeeScript
	DocxCodeLanguageD            DocxCodeLanguage = 14 // D
	DocxCodeLanguageDart         DocxCodeLanguage = 15 // Dart
	DocxCodeLanguageDelphi       DocxCodeLanguage = 16 // Delphi
	DocxCodeLanguageDjango       DocxCodeLanguage = 17 // Django
	DocxCodeLanguageDockerfile   DocxCodeLanguage = 18 // Dockerfile
	DocxCodeLanguageErlang       DocxCodeLanguage = 19 // Erlang
	DocxCodeLanguageFortran      DocxCodeLanguage = 20 // Fortran
	DocxCodeLanguageFoxPro       DocxCodeLanguage = 21 // FoxPro
	DocxCodeLanguageGo           DocxCodeLanguage = 22 // Go
	DocxCodeLanguageGroovy       DocxCodeLanguage = 23 // Groovy
	DocxCodeLanguageHTML         DocxCodeLanguage = 24 // HTML
	DocxCodeLanguageHTMLBars     DocxCodeLanguage = 25 // HTMLBars
	DocxCodeLanguageHTTP         DocxCodeLanguage = 26 // HTTP
	DocxCodeLanguageHaskell      DocxCodeLanguage = 27 // Haskell
	DocxCodeLanguageJSON         DocxCodeLanguage = 28 // JSON
	DocxCodeLanguageJava         DocxCodeLanguage = 29 // Java
	DocxCodeLanguageJavaScript   DocxCodeLanguage = 30 // JavaScript
	DocxCodeLanguageJulia        DocxCodeLanguage = 31 // Julia
	DocxCodeLanguageKotlin       DocxCodeLanguage = 32 // Kotlin
	DocxCodeLanguageLateX        DocxCodeLanguage = 33 // LateX
	DocxCodeLanguageLisp         DocxCodeLanguage = 34 // Lisp
	DocxCodeLanguageLogo         DocxCodeLanguage = 35 // Logo
	DocxCodeLanguageLua          DocxCodeLanguage = 36 // Lua
	DocxCodeLanguageMATLAB       DocxCodeLanguage = 37 // MATLAB
	DocxCodeLanguageMakefile     DocxCodeLanguage = 38 // Makefile
	DocxCodeLanguageMarkdown     DocxCodeLanguage = 39 // Markdown
	DocxCodeLanguageNginx        DocxCodeLanguage = 40 // Nginx
	DocxCodeLanguageObjective    DocxCodeLanguage = 41 // Objective
	DocxCodeLanguageOpenEdgeABL  DocxCodeLanguage = 42 // OpenEdgeABL
	DocxCodeLanguagePHP          DocxCodeLanguage = 43 // PHP
	DocxCodeLanguagePerl         DocxCodeLanguage = 44 // Perl
	DocxCodeLanguagePostScript   DocxCodeLanguage = 45 // PostScript
	DocxCodeLanguagePower        DocxCodeLanguage = 46 // Power
	DocxCodeLanguageProlog       DocxCodeLanguage = 47 // Prolog
	DocxCodeLanguageProtoBuf     DocxCodeLanguage = 48 // ProtoBuf
	DocxCodeLanguagePython       DocxCodeLanguage = 49 // Python
	DocxCodeLanguageR            DocxCodeLanguage = 50 // R
	DocxCodeLanguageRPG          DocxCodeLanguage = 51 // RPG
	DocxCodeLanguageRuby         DocxCodeLanguage = 52 // Ruby
	DocxCodeLanguageRust         DocxCodeLanguage = 53 // Rust
	DocxCodeLanguageSAS          DocxCodeLanguage = 54 // SAS
	DocxCodeLanguageSCSS         DocxCodeLanguage = 55 // SCSS
	DocxCodeLanguageSQL          DocxCodeLanguage = 56 // SQL
	DocxCodeLanguageScala        DocxCodeLanguage = 57 // Scala
	DocxCodeLanguageScheme       DocxCodeLanguage = 58 // Scheme
	DocxCodeLanguageScratch      DocxCodeLanguage = 59 // Scratch
	DocxCodeLanguageShell        DocxCodeLanguage = 60 // Shell
	DocxCodeLanguageSwift        DocxCodeLanguage = 61 // Swift
	DocxCodeLanguageThrift       DocxCodeLanguage = 62 // Thrift
	DocxCodeLanguageTypeScript   DocxCodeLanguage = 63 // TypeScript
	DocxCodeLanguageVBScript     DocxCodeLanguage = 64 // VBScript
	DocxCodeLanguageVisual       DocxCodeLanguage = 65 // Visual
	DocxCodeLanguageXML          DocxCodeLanguage = 66 // XML
	DocxCodeLanguageYAML         DocxCodeLanguage = 67 // YAML
)

// DocxOKRPeriodDisplayStatus OKR 周期的状态
type DocxOKRPeriodDisplayStatus string

const (
	DocxOKRPeriodDisplayStatusDefault DocxOKRPeriodDisplayStatus = "default" // 默认
	DocxOKRPeriodDisplayStatusNormal  DocxOKRPeriodDisplayStatus = "normal"  // 正常
	DocxOKRPeriodDisplayStatusInvalid DocxOKRPeriodDisplayStatus = "invalid" // 失效
	DocxOKRPeriodDisplayStatusHidden  DocxOKRPeriodDisplayStatus = "hidden"  // 隐藏
)

// DocxOKRProgressRateMode OKR 进展状态模式
type DocxOKRProgressRateMode string

const (
	DocxOKRProgressRateModeSimple   DocxOKRProgressRateMode = "simple"   // 简单模式
	DocxOKRProgressRateModeAdvanced DocxOKRProgressRateMode = "advanced" // 高级模式
)

// DocxOKRProgressStatus OKR 进展状态
type DocxOKRProgressStatus string

const (
	DocxOKRProgressStatusUnset    DocxOKRProgressStatus = "unset"    // 未设置
	DocxOKRProgressStatusNormal   DocxOKRProgressStatus = "normal"   // 正常
	DocxOKRProgressStatusRisk     DocxOKRProgressStatus = "risk"     // 有风险
	DocxOKRProgressStatusExtended DocxOKRProgressStatus = "extended" // 已延期
)

// DocxOKRProgressStatusType OKR 进展所展示的状态计算类型
type DocxOKRProgressStatusType string

const (
	DocxOKRProgressStatusTypeDefault DocxOKRProgressStatusType = "default" // 以风险最高的 Key Result 状态展示
	DocxOKRProgressStatusTypeCustom  DocxOKRProgressStatusType = "custom"  // 自定义
)
