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

// MessageContent ...
type MessageContent struct {
	Text               *MessageContentText               // 文本消息
	Image              *MessageContentImage              // 图片消息
	File               *MessageContentFile               // 文件消息
	Folder             *MessageContentFolder             // 文件夹消息
	Audio              *MessageContentAudio              // 音频消息
	Media              *MessageContentMedia              // 视频消息
	Sticker            *MessageContentSticker            // 表情包消息
	RedBag             *MessageContentRedBag             // 红包消息
	ShareCalendarEvent *MessageContentShareCalendarEvent // 分享日程消息
	ShareChat          *MessageContentShareChat          // 分享群名片消息
	ShareUser          *MessageContentShareUser          // 分享个人名片消息
	System             *MessageContentSystem             // 系统消息
	Location           *MessageContentLocation           // 位置消息
	VideoChat          *MessageContentVideoChat          // 视频通话消息
	Post               *MessageContentPost               // 富文本
	MsgType            MsgType
}

// MessageContentText 文本消息
type MessageContentText struct {
	Text string `json:"text"`
}

// MessageContentImage 图片消息
type MessageContentImage struct {
	ImageKey string `json:"image_key"`
}

// MessageContentFile 文件消息
type MessageContentFile struct {
	FileKey  string `json:"file_key"`
	FileName string `json:"file_name"`
}

// MessageContentFolder 文件夹消息
type MessageContentFolder struct {
	FileKey  string `json:"file_key"`
	FileName string `json:"file_name"` // 文件夹名称
}

// MessageContentAudio 音频消息
type MessageContentAudio struct {
	FileKey  string `json:"file_key"`
	Duration int    `json:"duration"`
}

// MessageContentMedia 视频消息
type MessageContentMedia struct {
	FileKey  string `json:"file_key"`
	ImageKey string `json:"image_key"`
	FileName string `json:"file_name"`
	Duration int    `json:"duration"`
}

// MessageContentSticker 表情包消息
type MessageContentSticker struct {
	FileKey string `json:"file_key"`
}

// MessageContentRedBag 红包消息
type MessageContentRedBag struct {
	Text string `json:"text"`
}

// MessageContentShareCalendarEvent 日程卡片消息
type MessageContentShareCalendarEvent struct {
	Summary   string `json:"summary"`
	StartTime string `json:"start_time"` // 毫秒级时间戳
	EndTime   string `json:"end_time"`   // 毫秒级时间戳
}

// MessageContentShareChat 群名片消息
type MessageContentShareChat struct {
	ChatID string `json:"chat_id"`
}

// MessageContentShareUser 个人名片消息
type MessageContentShareUser struct {
	UserID string `json:"user_id"` // open_id
}

// MessageContentSystem 系统消息消息
type MessageContentSystem struct {
	Template   string   `json:"template"`
	FromUser   []string `json:"from_user"`
	ToChatters []string `json:"to_chatters"`
}

// MessageContentLocation 位置消息
type MessageContentLocation struct {
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

// MessageContentVideoChat 视频通话消息
type MessageContentVideoChat struct {
	Topic     string `json:"topic"`
	StartTime string `json:"start_time"` // 时间戳毫秒级
}
