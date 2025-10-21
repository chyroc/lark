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
//
// docs: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/data-structure/document
type DocxDocument struct {
	DocumentID string `json:"document_id,omitempty"` // 文档的唯一标识，也是该文档对应的页面 Block 的 ID。https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/docx-overview
	RevisionID int64  `json:"revision_id,omitempty"` // 文档版本的标识，指定要查询或更新的文档版本。文档被编辑或评论会导致 revision_id 更改。 取值范围为：-1 或 0 < revision_id <= 文档最新版本。其中，-1 表示文档最新版本。文档创建后，版本为 1。注意：revision_id 发生更改通常意味着文档更新，但不一定是文档内容的变化，也有可能是文档被他人评论所致。
	Title      string `json:"title,omitempty"`       // 文档标题，只支持返回纯文本。
}

type DocxReferenceSynced struct {
	SourceBlockID    string `json:"source_block_id"`
	SourceDocumentID string `json:"source_document_id"`
}

type DocxSourceSynced struct {
	Align    DocxAlign          `json:"align,omitempty"`    // 对齐方式, 可选值有: `1`：居左排版, `2`：居中排版, `3`：居右排版
	Elements []*DocxTextElement `json:"elements,omitempty"` // 文本元素
}

// DocxBlock ...
//
// 在一篇文档中，有多个不同类型的段落，这些段落被定义为块（Block）。块是文档中的最小构建单元，是内容的结构化组成元素，有着明确的含义。块有多种形态，可以是一段文字、一张电子表格、一张图片或一个多维表格等。每个块都有唯一的 block_id 作为标识。
//
// docs: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/data-structure/block
type DocxBlock struct {
	BlockID   string        `json:"block_id,omitempty"`   // 块的唯一标识。创建块时会自动生成
	ParentID  string        `json:"parent_id,omitempty"`  // 块的父块 ID。除了根节点 Page 块外，其余块均有父块
	Children  []string      `json:"children,omitempty"`   // 块的子块 ID 列表
	BlockType DocxBlockType `json:"block_type,omitempty"` // 块的枚举值，表示块的类型
	// BlockData 只能是以下其中一种，并且需与 BlockType 相对应:
	Page            *DocxBlockText           `json:"page,omitempty"`             // 文档 Block, 1 DocxBlockTypePage
	Text            *DocxBlockText           `json:"text,omitempty"`             // 文本 Block, 2 DocxBlockTypeText
	Heading1        *DocxBlockText           `json:"heading1,omitempty"`         // 一级标题 Block, 3 DocxBlockTypeHeading1
	Heading2        *DocxBlockText           `json:"heading2,omitempty"`         // 二级标题 Block, 4 DocxBlockTypeHeading2
	Heading3        *DocxBlockText           `json:"heading3,omitempty"`         // 三级标题 Block, 5 DocxBlockTypeHeading3
	Heading4        *DocxBlockText           `json:"heading4,omitempty"`         // 四级标题 Block, 6 DocxBlockTypeHeading4
	Heading5        *DocxBlockText           `json:"heading5,omitempty"`         // 五级标题 Block, 7 DocxBlockTypeHeading5
	Heading6        *DocxBlockText           `json:"heading6,omitempty"`         // 六级标题 Block, 8 DocxBlockTypeHeading6
	Heading7        *DocxBlockText           `json:"heading7,omitempty"`         // 七级标题 Block, 9 DocxBlockTypeHeading7
	Heading8        *DocxBlockText           `json:"heading8,omitempty"`         // 八级标题 Block, 10 DocxBlockTypeHeading8
	Heading9        *DocxBlockText           `json:"heading9,omitempty"`         // 九级标题 Block, 11 DocxBlockTypeHeading9
	Bullet          *DocxBlockText           `json:"bullet,omitempty"`           // 无序列表 Block, 12 DocxBlockTypeBullet
	Ordered         *DocxBlockText           `json:"ordered,omitempty"`          // 有序列表 Block, 13 DocxBlockTypeOrdered
	Code            *DocxBlockText           `json:"code,omitempty"`             // 代码块 Block, 14 DocxBlockTypeCode
	Quote           *DocxBlockText           `json:"quote,omitempty"`            // 引用 Block, 15 DocxBlockTypeQuote
	Equation        *DocxBlockText           `json:"equation,omitempty"`         // 公式 Block, 16 DocxBlockTypeEquation
	Todo            *DocxBlockText           `json:"todo,omitempty"`             // 任务 Block, 17 DocxBlockTypeTodo
	Bitable         *DocxBlockBitable        `json:"bitable,omitempty"`          // 多维表格 Block, 18 DocxBlockTypeBitable
	Callout         *DocxBlockCallout        `json:"callout,omitempty"`          // 高亮块 Block, 19 DocxBlockTypeCallout
	ChatCard        *DocxBlockChatCard       `json:"chat_card,omitempty"`        // 群聊卡片 Block, 20 DocxBlockTypeChatCard
	Diagram         *DocxBlockDiagram        `json:"diagram,omitempty"`          // 流程图/UML Block, 21 DocxBlockTypeDiagram
	Divider         DocxBlockDivider         `json:"divider,omitempty"`          // 分割线 Block, 22 DocxBlockTypeDivider
	File            *DocxBlockFile           `json:"file,omitempty"`             // 文件 Block, 23 DocxBlockTypeFile
	Grid            *DocxBlockGrid           `json:"grid,omitempty"`             // 分栏 Block, 24 DocxBlockTypeGrid
	GridColumn      *DocxBlockGridColumn     `json:"grid_column,omitempty"`      // 分栏列 Block, 25 DocxBlockTypeGridColumn
	Iframe          *DocxBlockIframe         `json:"iframe,omitempty"`           // 内嵌 Block, 26 DocxBlockTypeIframe
	Image           *DocxBlockImage          `json:"image,omitempty"`            // 图片 Block, 27 DocxBlockTypeImage
	ISV             *DocxBlockISV            `json:"isv,omitempty"`              // 三方 Block, 28 DocxBlockTypeISV
	Mindnote        *DocxBlockMindnote       `json:"mindnote,omitempty"`         // 思维笔记 Block, 29 DocxBlockTypeMindnote
	Sheet           *DocxBlockSheet          `json:"sheet,omitempty"`            // 电子表格 Block, 30 DocxBlockTypeSheet
	Table           *DocxBlockTable          `json:"table,omitempty"`            // 表格 Block, 31 DocxBlockTypeTable
	TableCell       *DocxBlockTableCell      `json:"table_cell,omitempty"`       // 单元格 Block, 32 DocxBlockTypeTableCell
	View            *DocxBlockView           `json:"view,omitempty"`             // 视图 Block, 33 DocxBlockTypeView
	QuoteContainer  *DocxBlockQuoteContainer `json:"quote_container,omitempty"`  // 引用容器 Block, 34 DocxBlockTypeQuoteContainer
	Task            *DocxBlockTask           `json:"task,omitempty"`             // 任务容器 Block, 35 DocxBlockTask
	OKR             *DocxBlockOKR            `json:"okr,omitempty"`              // OKR Block, 36 DocxBlockTypeOKR
	OKRObjective    *DocxBlockOKRObjective   `json:"okr_objective,omitempty"`    // OKR Objective, 37 DocxBlockTypeOKRObjective
	OKRKeyResult    *DocxBlockOKRKeyResult   `json:"okr_key_result,omitempty"`   // OKR KR Block, 38 DocxBlockTypeOKRKeyResult
	OKRProgress     *DocxBlockOKRProgress    `json:"okr_progress,omitempty"`     // OKR Progress, 39 DocxBlockTypeOKRProgress
	AddOns          *DocxBlockAddOns         `json:"add_ons"`                    // 文档小组件 Block, 40 DocxBlockTypeAddOns
	JiraIssue       *DocxBlockJiraIssue      `json:"jira_issue,omitempty"`       // Jira 问题 Block, 41 DocxBlockJiraIssue
	WikiCatalog     *DocxBlockWikiCatalog    `json:"wiki_catalog,omitempty"`     // Wiki 子目录 Block, 42 DocxBlockWikiCatalog
	Board           *DocxBlockBoard          `json:"board,omitempty"`            // 画板 Block, 43 DocxBlockBoard
	SourceSynced    *DocxSourceSynced        `json:"source_synced,omitempty"`    //源同步块,49 DocSourceSynced
	ReferenceSynced *DocxReferenceSynced     `json:"reference_synced,omitempty"` //引用同步块,50 DocReferenceSynced
	Undefined       *DocxBlocUndefined       `json:"undefined,omitempty"`        // 未支持 Block, 999 DocxBlockTypeUndefined
}

// DocxBlockType ...
type DocxBlockType int64

const (
	DocxBlockTypePage           DocxBlockType = 1   // 页面(根) Block: DocxBlockText
	DocxBlockTypeText           DocxBlockType = 2   // 文本 Block: DocxBlockText
	DocxBlockTypeHeading1       DocxBlockType = 3   // 一级标题 Block: DocxBlockText
	DocxBlockTypeHeading2       DocxBlockType = 4   // 二级标题 Block: DocxBlockText
	DocxBlockTypeHeading3       DocxBlockType = 5   // 三级标题 Block: DocxBlockText
	DocxBlockTypeHeading4       DocxBlockType = 6   // 四级标题 Block: DocxBlockText
	DocxBlockTypeHeading5       DocxBlockType = 7   // 五级标题 Block: DocxBlockText
	DocxBlockTypeHeading6       DocxBlockType = 8   // 六级标题 Block: DocxBlockText
	DocxBlockTypeHeading7       DocxBlockType = 9   // 七级标题 Block: DocxBlockText
	DocxBlockTypeHeading8       DocxBlockType = 10  // 八级标题 Block: DocxBlockText
	DocxBlockTypeHeading9       DocxBlockType = 11  // 九级标题 Block: DocxBlockText
	DocxBlockTypeBullet         DocxBlockType = 12  // 无序列表 Block: DocxBlockText
	DocxBlockTypeOrdered        DocxBlockType = 13  // 有序列表 Block: DocxBlockText
	DocxBlockTypeCode           DocxBlockType = 14  // 代码块 Block: DocxBlockText
	DocxBlockTypeQuote          DocxBlockType = 15  // 引用 Block: DocxBlockText
	DocxBlockTypeEquation       DocxBlockType = 16  // 公式 Block: 文档没有了, 估计下线了
	DocxBlockTypeTodo           DocxBlockType = 17  // 任务 Block: DocxBlockText
	DocxBlockTypeBitable        DocxBlockType = 18  // 多维表格 Block: DocxBlockBitable
	DocxBlockTypeCallout        DocxBlockType = 19  // 高亮块 Block: DocxBlockCallout
	DocxBlockTypeChatCard       DocxBlockType = 20  // 群聊卡片 Block: DocxBlockChatCard
	DocxBlockTypeDiagram        DocxBlockType = 21  // 流程图/UML Block: DocxBlockDiagram
	DocxBlockTypeDivider        DocxBlockType = 22  // 分割线 Block: DocxBlockDivider
	DocxBlockTypeFile           DocxBlockType = 23  // 文件 Block: DocxBlockFile
	DocxBlockTypeGrid           DocxBlockType = 24  // 分栏 Block: DocxBlockGrid
	DocxBlockTypeGridColumn     DocxBlockType = 25  // 分栏列 Block: DocxBlockGridColumn
	DocxBlockTypeIframe         DocxBlockType = 26  // 内嵌 Block: DocxBlockIframe
	DocxBlockTypeImage          DocxBlockType = 27  // 图片 Block: DocxBlockImage
	DocxBlockTypeISV            DocxBlockType = 28  // 开放平台小组件 Block: DocxBlockISV
	DocxBlockTypeMindnote       DocxBlockType = 29  // 思维笔记 Block: DocxBlockMindnote
	DocxBlockTypeSheet          DocxBlockType = 30  // 电子表格 Block: DocxBlockSheet
	DocxBlockTypeTable          DocxBlockType = 31  // 表格 Block: DocxBlockTable
	DocxBlockTypeTableCell      DocxBlockType = 32  // 单元格 Block: DocxBlockTableCell
	DocxBlockTypeView           DocxBlockType = 33  // 视图 Block: DocxBlockView
	DocxBlockTypeQuoteContainer DocxBlockType = 34  // 引用容器 Block: DocxBlockQuoteContainer
	DocxBlockTypeTask           DocxBlockType = 35  // 任务容器 Block: DocxBlockTask
	DocxBlockTypeOKR            DocxBlockType = 36  // OKR容器 Block: DocxBlockOKR
	DocxBlockTypeOKRObjective   DocxBlockType = 37  // OKR Objective Block: DocxBlockOKRObjective
	DocxBlockTypeOKRKeyResult   DocxBlockType = 38  // OKR KeyResult Block: DocxBlockOKRKeyResult
	DocxBlockTypeOKRProgress    DocxBlockType = 39  // OKR Progress Block. DocxBlockOKRProgress
	DocxBlockTypeAddOns         DocxBlockType = 40  // 新版文档小组件 Block: DocxBlockAddOns
	DocxBlockTypeJiraIssue      DocxBlockType = 41  // Jira 问题 Block: DocxBlockJiraIssue
	DocxBlockTypeWikiCatalog    DocxBlockType = 42  // Wiki 子目录 Block: DocxBlockWikiCatalog
	DocxBlockTypeBoard          DocxBlockType = 43  // 画板 Block: DocxBlockBoard
	DocSourceSynced             DocxBlockType = 49  // 源同步块,49 DocSourceSynced
	DocReferenceSynced          DocxBlockType = 50  // 引用同步块,50 DocReferenceSynced
	DocxBlockTypeUndefined      DocxBlockType = 999 // 未支持 Block: DocxBlockUndefined
)

// -- text --

// DocxBlockText 文本 Block，其有多种 type。
//
// type:
// 1, DocxBlockTypePage
// 2, DocxBlockTypeText
// 3~11, DocxBlockTypeHeading1 ~ DocxBlockTypeHeading9
// 12, DocxBlockTypeBullet
// 13, DocxBlockTypeOrdered
// 14, DocxBlockTypeCode
// 15, DocxBlockTypeQuote
// 16, DocxBlockTypeEquation, 文档没有了, 估计下线了
// 17, DocxBlockTypeTodo
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

// -- text --

// -- Bitable --

// DocxBlockBitable ...
//
// 多维表格 Block。目前支持通过指定 view_type 创建空 Bitable。
//
// type: 18 DocxBlockTypeBitable
type DocxBlockBitable struct {
	Token    string              `json:"token,omitempty"`     // 多维表格在文档中的 Token，只读属性。 格式为 BitableToken_TableID，其中，BitableToken 是一篇多维表格唯一标识，TableID 是一张数据表的唯一标识，使用时请注意拆分。
	ViewType DocxBitableViewType `json:"view_type,omitempty"` // 视图类型, 1:数据表, 2:看板
}

// -- Bitable --

// -- Callout --

// DocxBlockCallout 高亮块 Block
//
// type: 19 DocxBlockTypeCallout
type DocxBlockCallout struct {
	BackgroundColor DocxCalloutBackgroundColor `json:"background_color,omitempty"` // 高亮块背景色
	BorderColor     DocxCalloutBorderColor     `json:"border_color,omitempty"`     // 边框色
	TextColor       DocxFontColor              `json:"text_color,omitempty"`       // 文字颜色
	EmojiID         string                     `json:"emoji_id,omitempty"`         // 高亮块图标
}

// -- Callout --

// -- ChatCard --

// DocxBlockChatCard 群聊卡片 Block
//
// type: 20 DocxBlockTypeChatCard
type DocxBlockChatCard struct {
	ChatID string    `json:"chat_id,omitempty"` // 群聊天会话 OpenID，其值以 ‘oc_’ 开头，表示专为开放能力接口定义的群组 ID。对于写操作，如果用户不在该群则返回无权限错误。
	Align  DocxAlign `json:"align,omitempty"`   // 对齐方式1：居左排版, 2：居中排版, 3：居右排版
}

// -- ChatCard --

// DocxBlockDiagram 流程图/UML Block
//
// type: 21 DocxBlockTypeDiagram
type DocxBlockDiagram struct {
	DiagramType DocxDiagramType `json:"diagram_type,omitempty"` // 绘图类型1：流程图, 2：UML 图
}

// DocxBlockDivider 分割线 Block
//
// type: 22 DocxBlockTypeDivider
type DocxBlockDivider struct{}

// DocxBlockFile 文件 Block
//
// 文件 Block。文件 Block 不能独立存在，须与视图 Block 一同出现，并且文件视图是通过控制视图 Block 的 view_type 实现的，比如卡片视图、预览视图等。在创建文件 Block 时，系统会自动生成默认的视图 Block。
//
// type: 23 DocxBlockTypeFile
type DocxBlockFile struct {
	Token    string       `json:"token,omitempty"` // 附件 Token
	Name     string       `json:"name,omitempty"`  // 文件名
	ViewType FileViewType `json:"view_type,omitempty"`
}

// DocxBlockGrid 分栏 Block
//
// type: 24 DocxBlockTypeGrid
type DocxBlockGrid struct {
	ColumnSize int64 `json:"column_size,omitempty"` // 分栏列数量，取值范围 [2,5)
}

// DocxBlockGridColumn 分栏列 Block
//
// type: 25 DocxBlockTypeGridColumn
type DocxBlockGridColumn struct {
	WidthRatio int64 `json:"width_ratio,omitempty"` // 当前分栏列占整个分栏的比例
}

// DocxBlockIframe 内嵌 Block
//
// type: 26 DocxBlockTypeIframe
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

// DocxBlockImage 图片 Block
//
// type: 27 DocxBlockTypeImage
type DocxBlockImage struct {
	Width  int64     `json:"width,omitempty"`  // 宽度单位 px
	Height int64     `json:"height,omitempty"` // 高度
	Token  string    `json:"token,omitempty"`  // 图片 Token
	Align  DocxAlign `json:"align,omitempty"`
}

// DocxBlockISV 三方 Block
//
// type: 28 DocxBlockTypeISV
type DocxBlockISV struct {
	ComponentID     string `json:"component_id,omitempty"`      // 团队互动应用唯一ID
	ComponentTypeID string `json:"component_type_id,omitempty"` // 团队互动应用类型，比如信息收集"blk_5f992038c64240015d280958"
}

// DocxBlockMindnote 思维笔记 Block
//
// 目前只支持查询返回占位信息，不支持创建及编辑。
//
// type: 29 DocxBlockTypeMindnote
type DocxBlockMindnote struct {
	Token string `json:"token,omitempty"` // 思维导图 token
}

// DocxBlockSheet 电子表格 Block
//
// 目前只指定 row_size 和 column_size 创建空 Sheet。
//
// type: 30 DocxBlockTypeSheet
type DocxBlockSheet struct {
	Token      string `json:"token,omitempty"`       // 电子表格 token。创建接口使用，如果非空则进行深拷贝
	RowSize    int64  `json:"row_size,omitempty"`    // 电子表格列数量。创建空电子表格时使用，最大值 9
	ColumnSize int64  `json:"column_size,omitempty"` // 电子表格列数量。创建空电子表格时使用，最大值 9
}

// DocxBlockTable 表格 Block
//
// type: 31 DocxBlockTypeTable
type DocxBlockTable struct {
	Cells    []string                `json:"cells,omitempty"`    // 单元格数组，数组元素为 Table Cell Block 的 ID
	Property *DocxBlockTableProperty `json:"property,omitempty"` // 表格属性
}

// DocxBlockTableProperty 表格属性
type DocxBlockTableProperty struct {
	RowSize      int64                              `json:"row_size,omitempty"`      // 行数
	ColumnSize   int64                              `json:"column_size,omitempty"`   // 列数
	ColumnWidth  []int64                            `json:"column_width,omitempty"`  // 列宽，单位px
	HeaderRow    bool                               `json:"header_row,omitempty"`    // 是否设置首行为标题行。
	HeaderColumn bool                               `json:"header_column,omitempty"` // 是否设置首列为标题列。
	MergeInfo    []*DocxBlockTablePropertyMergeInfo `json:"merge_info,omitempty"`    // 单元格合并信息
}

// DocxBlockTablePropertyMergeInfo 单元格合并信息
type DocxBlockTablePropertyMergeInfo struct {
	RowSpan int64 `json:"row_span,omitempty"` // 从当前行索引起被合并的连续行数
	ColSpan int64 `json:"col_span,omitempty"` // 从当前列索引起被合并的连续列数
}

// DocxBlockTableCell 单元格 Block
//
// type: 32 DocxBlockTypeTableCell
type DocxBlockTableCell struct{}

// DocxBlockView 视图 Block
//
// type: 33 DocxBlockTypeView
type DocxBlockView struct {
	ViewType DocxViewType `json:"view_type,omitempty"` // 视图类型, 可选值有：1：卡片视图, 2：预览视图, 3：内联试图
}

// DocxBlockQuoteContainer 引用容器 Block
//
// type: 34 DocxBlockTypeQuoteContainer
type DocxBlockQuoteContainer struct{}

// DocxBlockTask ...
//
// 任务块的内容实体。注意你只能获取任务块的任务 ID，无法创建或编辑任务块。如需获取任务详情，调用获取任务详情接口。
//
// type: 35 DocxBlockTypeTask
type DocxBlockTask struct {
	TaskID string `json:"task_id,omitempty"` // 任务 ID
}

// DocxBlockOKR OKR 信息
//
// type: 36 DocxBlockTypeOKR
type DocxBlockOKR struct {
	OKRID               string                     `json:"okr_id,omitempty"`                // OKR ID
	PeriodDisplayStatus DocxOKRPeriodDisplayStatus `json:"period_display_status,omitempty"` // 周期的状态
	PeriodNameEn        string                     `json:"period_name_en,omitempty"`        // 周期名 - 英文
	PeriodNameZh        string                     `json:"period_name_zh,omitempty"`        // 周期名 - 中文
	UserID              string                     `json:"user_id,omitempty"`               // OKR 所属的用户 ID
	VisibleSetting      *DocxOKRVisibleSetting     `json:"visible_setting,omitempty"`       // 可见性设置
}

// DocxOKRVisibleSetting OKR 可见设置
type DocxOKRVisibleSetting struct {
	ProgressFillAreaVisible bool `json:"progress_fill_area_visible,omitempty"`
	ProgressStatusVisible   bool `json:"progress_status_visible,omitempty"`
	ScoreVisible            bool `json:"score_visible,omitempty"`
}

// DocxBlockOKRObjective OKR Objective 信息
//
// type: 37 DocxBlockTypeOKRObjective
type DocxBlockOKRObjective struct {
	ObjectiveID  string               `json:"objective_id,omitempty"`  // OKR Objective ID
	Confidential bool                 `json:"confidential,omitempty"`  // 是否在 OKR 平台设置了私密权限
	Position     int64                `json:"position,omitempty"`      // Objective 的位置编号，对应 Block 中 O1、O2 的 1、2
	Score        int64                `json:"score,omitempty"`         // 打分信息
	Visible      bool                 `json:"visible,omitempty"`       // OKR Block 中是否展示该 Objective
	Weight       int64                `json:"weight,omitempty"`        // Objective 的权重
	ProgressRate *DocxOKRProgressRate `json:"progress_rate,omitempty"` // 进展信息
	Content      *DocxBlockText       `json:"content,omitempty"`       // Objective 的文本内容
}

// DocxBlockOKRKeyResult OKR KR 信息
//
// type: 38 DocxBlockTypeOKRKeyResult
type DocxBlockOKRKeyResult struct {
	KeyResultID  string               `json:"kr_id,omitempty"`         // OKR Key Result ID
	Confidential bool                 `json:"confidential,omitempty"`  // 是否在 OKR 平台设置了私密权限
	Position     int64                `json:"position,omitempty"`      // Key Result 的位置编号，对应 Block 中 KR1、KR2 的 1、2
	Score        int64                `json:"score,omitempty"`         // 打分信息
	Visible      bool                 `json:"visible,omitempty"`       // OKR Block 中是否展示该 Key Result
	Weight       float32              `json:"weight,omitempty"`        // Key Result 的权重
	ProgressRate *DocxOKRProgressRate `json:"progress_rate,omitempty"` // 进展信息
	Content      *DocxBlockText       `json:"content,omitempty"`       // Key Result 的文本内容
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

// DocxBlockOKRProgress OKR 进展块的内容实体，为空结构体。
//
// type: 39 DocxBlockTypeOKRProgress
type DocxBlockOKRProgress struct{}

// DocxBlockAddOns 文档小组件块的内容实体
//
// type: 40 DocxBlockTypeAddOns
type DocxBlockAddOns struct {
	ComponentID     string `json:"component_id,omitempty"`      // AddOns 文档小组件 ID。该字段为空。
	ComponentTypeID string `json:"component_type_id,omitempty"` // 文档小组件类型 ID，用于区分不同类型的小组件，比如问答互动类。
	Record          string `json:"record"`                      // 文档小组件内容数据，JSON 字符串
}

// DocxBlockJiraIssue Jira 问题块的内容实体
//
// type: 41 DocxBlockTypeJiraIssue
type DocxBlockJiraIssue struct {
	ID  string `json:"id,omitempty"`  // Jira 问题的 ID。
	Key string `json:"key,omitempty"` //	Jira 问题的 Key。
}

// DocxBlockWikiCatalog Wiki 子目录块的内容实体
//
// type: 42 DocxBlockTypeWikiCatalog
type DocxBlockWikiCatalog struct {
	WikiToken string `json:"wiki_token,omitempty"` // 	知识库 token ，默认为当前文档所属知识库节点 token。
}

// DocxBlockBoard Wiki 子目录块的内容实体
//
// type: 43 DocxBlockTypeBoard
type DocxBlockBoard struct {
	Token  string    `json:"token,omitempty"`  // 画板的 token，只读属性，插入画板块后自动生成
	Align  DocxAlign `json:"align,omitempty"`  // 对齐方式。
	Width  int64     `json:"width,omitempty"`  // 宽度，单位像素；不填时自动适应文档宽度；值超出文档最大宽度时，页面渲染为文档最大宽度
	Height int64     `json:"height,omitempty"` // 高度，单位像素；不填时自动根据画板内容计算；值超出屏幕两倍高度时，页面渲染为屏幕两倍高度
}

// DocxBlocUndefined 未支持 Block
//
// 其它暂未支持的内容实体。只支持获取其 ID。这些内容实体不支持创建和更新，但支持删除。
//
// type: 999 DocxBlockTypeUndefined
type DocxBlocUndefined struct{}

// 枚举类型

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

type FileViewType int64

const (
	FileViewTypeCard    FileViewType = 1 // 卡片视图（默认）
	FileViewTypePreview FileViewType = 2 // 预览视图
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
