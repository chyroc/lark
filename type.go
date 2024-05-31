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

// MsgType 消息类型
type MsgType string

// MsgTypeText ...
const (
	MsgTypeText               MsgType = "text"                 // 文本
	MsgTypePost               MsgType = "post"                 // 富文本
	MsgTypeImage              MsgType = "image"                // 图片
	MsgTypeFile               MsgType = "file"                 // 文件
	MsgTypeFolder             MsgType = "folder"               // 文件夹
	MsgTypeAudio              MsgType = "audio"                // 音频
	MsgTypeMedia              MsgType = "media"                // 视频
	MsgTypeSticker            MsgType = "sticker"              // 表情包
	MsgTypeInteractive        MsgType = "interactive"          // 卡片消息
	MsgTypeRedBag             MsgType = "red_bag"              // 红包
	MsgTypeShareCalendarEvent MsgType = "share_calendar_event" // 日程分享卡片
	MsgTypeShareChat          MsgType = "share_chat"           // 分享群卡片
	MsgTypeShareUser          MsgType = "share_user"           // 分享个人卡片
	MsgTypeMergeForward       MsgType = "merge_forward"        // 合并转发消息内容
	MsgTypeSystem             MsgType = "system"               // 系统消息
	MsgTypeLocation           MsgType = "location"             // 地理位置
	MsgTypeVideoChat          MsgType = "video_chat"           // 视频通话
)

// MsgTypePtr ...
func MsgTypePtr(v MsgType) *MsgType {
	return &v
}

// ContainerIDType 容器类型
type ContainerIDType string

// ContainerIDTypeChat ...
const (
	ContainerIDTypeChat   ContainerIDType = "chat"
	ContainerIDTypeThread ContainerIDType = "thread"
)

// ContainerIDTypePtr ...
func ContainerIDTypePtr(v ContainerIDType) *ContainerIDType {
	return &v
}

// IDType ID类型
type IDType string

// IDTypeUserID ...
const (
	IDTypeUserID  IDType = "user_id"  // 以 user_id 来识别成员
	IDTypeUnionID IDType = "union_id" // 以 union_id 来识别成员
	IDTypeOpenID  IDType = "open_id"  // 以 open_id 来识别成员
	IDTypeAppID   IDType = "app_id"   // 以 app_id 来识别成员
	IDTypeChatID  IDType = "chat_id"  // 以 chat_id 来识别成员
	IDTypeEmail   IDType = "email"    // 以 email 来识别成员
)

// IDTypePtr ...
func IDTypePtr(idType IDType) *IDType {
	return &idType
}

// DepartmentIDType ID类型
type DepartmentIDType string

// DepartmentIDTypeDepartmentID ...
const (
	DepartmentIDTypeDepartmentID     DepartmentIDType = "department_id"      // 以 department_id 来识别
	DepartmentIDTypeOpenDepartmentID DepartmentIDType = "open_department_id" // 以 open_department_id 来识别
)

// DepartmentIDTypePtr ...
func DepartmentIDTypePtr(v DepartmentIDType) *DepartmentIDType {
	return &v
}

// MailUserType ...
type MailUserType string

// MailUserTypeUser ...
const (
	MailUserTypeUser         MailUserType = "USER"          // 内部用户
	MailUserTypeDepartment   MailUserType = "DEPARTMENT"    // 部门
	MailUserTypeCompany      MailUserType = "COMPANY"       // 全员
	MailUserTypeExternalUser MailUserType = "EXTERNAL_USER" // 外部用户
	MailUserTypeMailGroup    MailUserType = "MAIL_GROUP"    // 邮件组
	MailUserTypeOtherMember  MailUserType = "OTHER_MEMBER"  // 内部成员
)

// MailUserTypePtr ...
func MailUserTypePtr(v MailUserType) *MailUserType {
	return &v
}

// EmployeeType 用户类型
type EmployeeType string

// EmployeeTypeID ...
const (
	EmployeeTypeID EmployeeType = "employee_id" // 员工id
	EmployeeTypeNo EmployeeType = "employee_no" // 员工工号
)

// EmployeeTypePtr ...
func EmployeeTypePtr(v EmployeeType) *EmployeeType {
	return &v
}

// ChatMode 群模式(group/topic/p2p)
type ChatMode string

// ChatModeP2P ...
const (
	ChatModeP2P   ChatMode = "p2p"
	ChatModeGroup ChatMode = "group"
	ChatModeTopic ChatMode = "topic"
)

// ChatModePtr ...
func ChatModePtr(v ChatMode) *ChatMode {
	return &v
}

// ChatType 群类型(private/public)
type ChatType string

// ChatTypePrivate ...
const (
	ChatTypePrivate ChatType = "private"
	ChatTypePublic  ChatType = "public"
)

// ChatTypePtr ...
func ChatTypePtr(v ChatType) *ChatType {
	return &v
}

// ImageType 图片类型
type ImageType string

// ImageTypeMessage ...
const (
	ImageTypeMessage ImageType = "message" // 用于发送消息
	ImageTypeAvatar  ImageType = "avatar"  // 用于设置头像
)

// ImageTypePtr ...
func ImageTypePtr(v ImageType) *ImageType {
	return &v
}

// FileType 文件类型
type FileType string

// FileTypeOpus ...
const (
	FileTypeOpus    FileType = "opus"    // opus音频文件；其他格式的音频文件，请转为opus格式后上传，转换方式可参考：ffmpeg -i SourceFile.mp3 -acodec libopus -ac 1 -ar 16000 TargetFile.opus
	FileTypeMp4     FileType = "mp4"     // mp4视频文件
	FileTypePdf     FileType = "pdf"     // pdf格式文件
	FileTypeDoc     FileType = "doc"     // doc格式文件
	FileTypeDocx    FileType = "docx"    // docx格式文件
	FileTypeWiKi    FileType = "wiki"    // wiki格式文件
	FileTypeBitable FileType = "bitable" // 多维表格
	FileTypeXls     FileType = "xls"     // xls格式文件
	FileTypePpt     FileType = "ppt"     // ppt格式文件
	FileTypeStream  FileType = "stream"  // stream格式文件
)

// FileTypePtr ...
func FileTypePtr(v FileType) *FileType {
	return &v
}

// CalendarRole 对日历的访问权限
type CalendarRole string

// CalendarRoleUnknown ...
const (
	CalendarRoleUnknown        CalendarRole = "unknown"          // 未知权限
	CalendarRoleFreeBusyReader CalendarRole = "free_busy_reader" // 游客，只能看到忙碌/空闲信息
	CalendarRoleReader         CalendarRole = "reader"           // 订阅者，查看所有日程详情
	CalendarRoleWriter         CalendarRole = "writer"           // 编辑者，创建及修改日程
	CalendarRoleOwner          CalendarRole = "owner"            // 管理员，管理日历及共享设置
)

// CalendarRolePtr ...
func CalendarRolePtr(v CalendarRole) *CalendarRole {
	return &v
}

// 参与人类型
type CalendarEventAttendeeType string

// CalendarEventAttendeeTypeUser ...
const (
	CalendarEventAttendeeTypeUser       CalendarEventAttendeeType = "user"        // 用户
	CalendarEventAttendeeTypeChat       CalendarEventAttendeeType = "chat"        // 群组
	CalendarEventAttendeeTypeResource   CalendarEventAttendeeType = "resource"    // 会议室
	CalendarEventAttendeeTypeThirdParty CalendarEventAttendeeType = "third_party" // 邮箱
)

// CalendarEventAttendeeTypePtr ...
func CalendarEventAttendeeTypePtr(v CalendarEventAttendeeType) *CalendarEventAttendeeType {
	return &v
}

// CalendarType ...
type CalendarType string

// CalendarTypeUnknown ...
const (
	CalendarTypeUnknown  CalendarType = "unknown"  // 未知类型
	CalendarTypePrimary  CalendarType = "primary"  // 用户或应用的主日历
	CalendarTypeShared   CalendarType = "shared"   // 由用户或应用创建的共享日历
	CalendarTypeGoogle   CalendarType = "google"   // 用户绑定的谷歌日历
	CalendarTypeResource CalendarType = "resource" // 会议室日历
	CalendarTypeExchange CalendarType = "exchange" // 用户绑定的Exchange日历
)

// CalendarTypePtr ...
func CalendarTypePtr(v CalendarType) *CalendarType {
	return &v
}

// CalendarPermission ...
type CalendarPermission string

// CalendarPermissionPrivate ...
const (
	CalendarPermissionPrivate          = "private"             // 私密
	CalendarPermissionShowOnlyFreeBusy = "show_only_free_busy" // 仅展示忙闲信息
	CalendarPermissionPublic           = "public"              // 他人可查看日程详情
)

// CalendarPermissionPtr ...
func CalendarPermissionPtr(v CalendarPermission) *CalendarPermission {
	return &v
}

// I18nNames i18n name
type I18nNames struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名, 示例值: "群聊"
	EnUs string `json:"en_us,omitempty"` // 英文名, 示例值: "group chat"
	JaJp string `json:"ja_jp,omitempty"` // 日文名, 示例值: "グループチャット"
}

// Sender sender user
type Sender struct {
	ID         string `json:"id,omitempty"`          // 该字段标识发送者的id
	IDType     IDType `json:"id_type,omitempty"`     // 该字段标识发送者的id类型
	SenderType string `json:"sender_type,omitempty"` // 该字段标识发送者的类型
}

// MessageBody message body
type MessageBody struct {
	Content string `json:"content,omitempty"` // 消息jsonContent
}

// Mention mention user
type Mention struct {
	Key    string `json:"key,omitempty"`     // mention key
	ID     string `json:"id,omitempty"`      // 用户open id
	IDType IDType `json:"id_type,omitempty"` // id 可以是open_id，user_id或者union_id
	Name   string `json:"name,omitempty"`    // 被at用户的姓名
}

// AddMemberPermission 加 user/bot 入群权限
type AddMemberPermission string

// AddMemberPermissionAllMembers ...
const (
	AddMemberPermissionAllMembers AddMemberPermission = "all_members"
	AddMemberPermissionOnlyOwner  AddMemberPermission = "only_owner"
)

// AddMemberPermissionPtr ...
func AddMemberPermissionPtr(v AddMemberPermission) *AddMemberPermission {
	return &v
}

// MessageVisibility 入群消息可见性
type MessageVisibility string

// MessageVisibilityOnlyOwner ...
const (
	MessageVisibilityOnlyOwner  MessageVisibility = "only_owner"
	MessageVisibilityAllMembers MessageVisibility = "all_members"
	MessageVisibilityNotAnyone  MessageVisibility = "not_anyone"
)

// MessageVisibilityPtr ...
func MessageVisibilityPtr(v MessageVisibility) *MessageVisibility {
	return &v
}

// MembershipApproval 加群审批
type MembershipApproval string

// MembershipApprovalNoApprovalRequired ...
const (
	MembershipApprovalNoApprovalRequired MembershipApproval = "no_approval_required"
	MembershipApprovalApprovalRequired   MembershipApproval = "approval_required"
)

// MembershipApprovalPtr ...
func MembershipApprovalPtr(v MembershipApproval) *MembershipApproval {
	return &v
}

// ModerationPermission 发言权限
type ModerationPermission string

// ModerationPermissionAllMembers ...
const (
	ModerationPermissionAllMembers    ModerationPermission = "all_members"
	ModerationPermissionOnlyOwner     ModerationPermission = "only_owner"
	ModerationPermissionModeratorList ModerationPermission = "moderator_list"
)

// ModerationPermissionPtr ...
func ModerationPermissionPtr(v ModerationPermission) *ModerationPermission {
	return &v
}

// ShareCardPermission 群分享权限
type ShareCardPermission string

// ShareCardPermissionAllowed ...
const (
	ShareCardPermissionAllowed    ShareCardPermission = "allowed"
	ShareCardPermissionNotAllowed ShareCardPermission = "not_allowed"
)

// ShareCardPermissionPtr ...
func ShareCardPermissionPtr(v ShareCardPermission) *ShareCardPermission {
	return &v
}

// AtAllPermission at 所有人权限
type AtAllPermission string

// AtAllPermissionAllMembers ...
const (
	AtAllPermissionAllMembers AtAllPermission = "all_members"
	AtAllPermissionOnlyOwner  AtAllPermission = "only_owner"
)

// AtAllPermissionPtr ...
func AtAllPermissionPtr(v AtAllPermission) *AtAllPermission {
	return &v
}

// EditPermission 群编辑权限
type EditPermission string

// EditPermissionAllMembers ...
const (
	EditPermissionAllMembers EditPermission = "all_members"
	EditPermissionOnlyOwner  EditPermission = "only_owner"
)

// EditPermissionPtr ...
func EditPermissionPtr(v EditPermission) *EditPermission {
	return &v
}

// HelpdeskCustomizedField helpdesk custom field
type HelpdeskCustomizedField struct {
	ID      string `json:"id"`       // id ,示例值："123"
	Value   string `json:"value"`    // value ,示例值："value"
	KeyName string `json:"key_name"` // key name ,示例值："key"
}

// HelpdeskDropdownOption 下拉列表选项
type HelpdeskDropdownOption struct {
	Tag         string                    `json:"tag,omitempty"`          // 选项ID
	DisplayName string                    `json:"display_name,omitempty"` // 展示名称
	Children    []*HelpdeskDropdownOption `json:"children,omitempty"`     // 同上：选项列表，只适用于多层下拉列表（最多可以设置三级下拉列表）
}

// HelpdeskCategory 知识库分类
type HelpdeskCategory struct {
	CategoryID string              `json:"category_id,omitempty"` // 知识库分类ID
	ID         string              `json:"id,omitempty"`          // 知识库分类ID，（旧版，请使用category_id）
	Name       string              `json:"name,omitempty"`        // 名称
	ParentID   string              `json:"parent_id,omitempty"`   // 父知识库分类ID
	HelpdeskID string              `json:"helpdesk_id,omitempty"` // 服务台ID
	Language   string              `json:"language,omitempty"`    // 语言
	Children   []*HelpdeskCategory `json:"children,omitempty"`    // 子分类详情
}

// EHRAttachment EHR 附件
type EHRAttachment struct {
	ID       string `json:"id,omitempty"`        // 下载文件所需要的 Token
	MimeType string `json:"mime_type,omitempty"` // 文件类型
	Name     string `json:"name,omitempty"`      // 名称
	Size     int    `json:"size,omitempty"`      // 大小
}

// EHREducation EHR 教育经历
type EHREducation struct {
	Level  int    `json:"level,omitempty"`  // 学历, 可选值有: `1`：小学, `2`：初中, `3`：高中, `4`：职业高级中学, `5`：中等专业学校, `6`：大专, `7`：本科, `8`：硕士, `9`：博士
	School string `json:"school,omitempty"` // 毕业学校
	Major  string `json:"major,omitempty"`  // 专业
	Degree int    `json:"degree,omitempty"` // 学位, 可选值有: `1`：学士, `2`：硕士, `3`：博士
	Start  string `json:"start,omitempty"`  // 开始日期
	End    string `json:"end,omitempty"`    // 结束日期
}

// EHREmergencyContact EHR 紧急联系人
type EHREmergencyContact struct {
	Name         string `json:"name,omitempty"`         // 紧急联系人姓名
	Relationship int    `json:"relationship,omitempty"` // 与紧急联系人的关系, 可选值有: `1`：父母, `2`：配偶, `3`：子女, `4`：兄弟姐妹, `5`：朋友, `6`：其他
	Mobile       string `json:"mobile,omitempty"`       // 手机号
}

// EHRWorkExperience EHR 工作经历
type EHRWorkExperience struct {
	Company     string `json:"company,omitempty"`     // 公司
	Department  string `json:"department,omitempty"`  // 部门
	Job         string `json:"job,omitempty"`         // 职位
	Start       string `json:"start,omitempty"`       // 开始日期
	End         string `json:"end,omitempty"`         // 截止日期
	Description string `json:"description,omitempty"` // 工作描述
}

// EventHeaderV2 ...
type EventHeaderV2 struct {
	EventID    string    `json:"event_id,omitempty"`    // 事件 ID
	EventType  EventType `json:"event_type,omitempty"`  // 事件类型
	CreateTime string    `json:"create_time,omitempty"` // 事件创建时间戳（单位：毫秒）
	Token      string    `json:"token,omitempty"`       // 事件 Token
	AppID      string    `json:"app_id,omitempty"`      // 应用 ID
	TenantKey  string    `json:"tenant_key,omitempty"`  // 租户 Key
}

// EventHeaderV1 ...
type EventHeaderV1 struct {
	UUID      string    `json:"event_id,omitempty"`    // 事件 ID
	EventType EventType `json:"event_type,omitempty"`  // 事件类型
	TS        string    `json:"create_time,omitempty"` // 事件创建时间戳（单位：毫秒）
	Token     string    `json:"token,omitempty"`       // 事件 Token
}
