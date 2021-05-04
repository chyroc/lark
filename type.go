package lark

// const type

// MsgType 消息类型
type MsgType string

const (
	MsgTypeText        MsgType = "text"        // 文本
	MsgTypePost        MsgType = "post"        // 富文本
	MsgTypeImage       MsgType = "image"       // 图片
	MsgTypeFile        MsgType = "file"        // 文件
	MsgTypeAudio       MsgType = "audio"       // 语音
	MsgTypeMedia       MsgType = "media"       // 视频
	MsgTypeSticker     MsgType = "sticker"     // 表情包
	MsgTypeInteractive MsgType = "interactive" // 卡片消息
	MsgTypeShareChat   MsgType = "share_chat"  // 分享群卡片
	MsgTypeShareUser   MsgType = "share_user"  // 分享个人卡片
)

// ContainerIDType 容器类型
type ContainerIDType string

const (
	ContainerIDTypeChat ContainerIDType = "chat"
)

// IDType ID类型
type IDType string

const (
	IDTypeUserID  IDType = "user_id"  // 以 user_id 来识别成员
	IDTypeUnionID IDType = "union_id" // 以 union_id 来识别成员
	IDTypeOpenID  IDType = "open_id"  // 以 open_id 来识别成员
	IDTypeAppID   IDType = "app_id"   // 以 app_id 来识别成员
	IDTypeChatID  IDType = "chat_id"  // 以 chat_id 来识别成员
	IDTypeEmail   IDType = "email"    // 以 email 来识别成员
)

func IDTypePtr(idType IDType) *IDType {
	return &idType
}

// ImageType 图片类型
type ImageType string

const (
	ImageTypeMessage ImageType = "message" // 用于发送消息
	ImageTypeAvatar  ImageType = "avatar"  // 用于设置头像
)

// FileType 文件类型
type FileType string

const (
	FileTypeOpus   FileType = "opus"   // 上传opus音频文件；其他格式的音频文件，请转为opus格式后上传，转换方式可参考：ffmpeg -i SourceFile.mp3 -acodec libopus -ac 1 -ar 16000 TargetFile.opus
	FileTypeMp4    FileType = "mp4"    // 上传mp4视频文件
	FileTypePdf    FileType = "pdf"    // 上传pdf格式文件
	FileTypeDoc    FileType = "doc"    // 上传doc格式文件
	FileTypeXls    FileType = "xls"    // 上传xls格式文件
	FileTypePpt    FileType = "ppt"    // 上传ppt格式文件
	FileTypeStream FileType = "stream" // 上传stream格式文件
)

// EmployeeType 用户类型
type EmployeeType string

const (
	EmployeeTypeID EmployeeType = "employee_id" // 员工id
	EmployeeTypeNo EmployeeType = "employee_no" // 员工工号
)

type I18nNames struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名,**示例值**："群聊"
	EnUs string `json:"en_us,omitempty"` // 英文名,**示例值**："group chat"
	JaJp string `json:"ja_jp,omitempty"` // 日文名,**示例值**："グループチャット"
}

type Sender struct {
	Id         string `json:"id,omitempty"`          // 该字段标识发送者的id
	IDType     IDType `json:"id_type,omitempty"`     // 该字段标识发送者的id类型
	SenderType string `json:"sender_type,omitempty"` // 该字段标识发送者的类型
}

type MessageBody struct {
	Content string `json:"content,omitempty"` // 消息jsonContent
}

type Mention struct {
	Key    string `json:"key,omitempty"`     // mention key
	Id     string `json:"id,omitempty"`      // 用户open id
	IDType IDType `json:"id_type,omitempty"` // id 可以是open_id，user_id或者union_id
	Name   string `json:"name,omitempty"`    // 被at用户的姓名
}

type HelpdeskCustomizedField struct {
	ID      string `json:"id"`       // id ,示例值："123"
	Value   string `json:"value"`    // value ,示例值："value"
	KeyName string `json:"key_name"` // key name ,示例值："key"
}

type MessageContent struct {
	Text     string `json:"text"`
	ImageKey string `json:"image_key"`
}
