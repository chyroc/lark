# lark

[![codecov](https://codecov.io/gh/chyroc/lark/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark "go report card")](https://goreportcard.com/report/github.com/chyroc/lark)
[![test status](https://github.com/chyroc/lark/actions/workflows/go.yml/badge.svg)](https://github.com/chyroc/lark/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark)

[English README](./README.md)

飞书/Lark 的开放接口 Go SDK，支持所有的开放接口，和事件回调。

使用代码生成创建。

## 安装

```shell
go get github.com/chyroc/lark
```

## 文档

https://godoc.org/github.com/chyroc/lark

## 使用

### 例子: 创建 Lark 客户端

- 使用简单的 App

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))
```

- 处理事件回调

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithEventCallbackVerify("<ENCRYPY_KEY>", "<VERIFICATION_TOKEN>"),
)
```

- 服务台

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithHelpdeskCredential("<HELPDESK_ID>", "HELPDESK_TOKEN"),
)
```

### 例子：处理事件回调

如果需要更多的例子，可以参考：[./examples/event_callback.go](./examples/event_callback.go) 。

处理消息的事件回调：

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithEventCallbackVerify("<ENCRYPY_KEY>", "<VERIFICATION_TOKEN>"),
)

// handle message callback
cli.EventCallback().HandlerEventIMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeader, event *lark.EventIMMessageReceiveV1) (string, error) {
    _, _, err := cli.Message().Reply(event.Message.MessageID).SendText(ctx, "hi, "+event.Message.Content)
    return "", err
})

http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    cli.EventCallback().ListenCallback(r.Context(), r.Body, w)
})

fmt.Println("start server ...")
log.Fatal(http.ListenAndServe(":9726", nil))
```


### 例子: 获取机器人信息

如果需要更多的例子，可以参考： [./examples/bot.go](./examples/bot.go) 。

获取机器人信息：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Bot().GetBotInfo(ctx, &lark.GetBotInfoReq{})
fmt.Println(resp, err)
```

### 例子: 发送消息

如果需要更多的例子，可以参考： [./examples/send_message.go](./examples/send_message.go) 。

发送文本消息：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendText(ctx, "<TEXT>")
fmt.Println(resp, err)
```

### 例子: 其他消息

如果需要更多的例子，可以参考： [./examples/other_message.go](./examples/other_message.go) 。

撤回消息：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message().DeleteMessage(ctx, &lark.DeleteMessageReq{
    MessageID: "<MESSAGE_ID>",
})
fmt.Println(resp, err)
```

### 例子: 群聊

如果需要更多的例子，可以参考： [./examples/chat.go](./examples/chat.go) 。

创建群聊：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{
    Name: ptr.String("<CHAT_NAME>"),
})
fmt.Println(resp, err)
```

### 例子: 文件

如果需要更多的例子，可以参考： [./examples/file.go](./examples/file.go) 。

上传图片

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

f, err := os.Open("<FILE_PATH>")
if err != nil {
    panic(err)
}
resp, _, err := cli.File().UploadImage(ctx, &lark.UploadImageReq{
    ImageType: lark.ImageTypeMessage,
    Image:     f,
})
fmt.Println(resp, err)
```

### 例子: 日历

如果需要更多的例子，可以参考： [./examples/calendar.go](./examples/calendar.go) 。

创建日历

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Calendar().CreateCalendar(ctx, &lark.CreateCalendarReq{
Summary: ptr.String("<SUMMARY>"),
})
fmt.Println(resp, err)
```

## Todo

- [ ] 单侧覆盖率 >= 80%
- [ ] 发送消息，提供helper pack
- [ ] 收到消息，提供helper unpack
- [ ] 生成所有 API
- [ ] 英文注释
- [ ] 结构化的数据

*Lark 是指针，是 mock 和 event handler 的实际注册 receiver

如何获取 app-ticket？

isv，刚开始应该是没有 app-ticket 的，假设通过某种方式，设置了一个 resend 请求，收到事件

https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM

https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN#1d112a8a

然后通过内置的 getter setter，设置到 store

这个时候可以获取到 app-token 了

如何获取 tenant-key ？

首次启用 app，购买，停用，卸载，可用性修改
https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN

schema=2 的消息的 header，都有 tenant_key

- Clone 的时候，这些值是否需要传递到复制后的上面？
- 
- user token
  - WithUser(user-token).Drive().CreateFolder() ?
  - Drive().CreateFolder(ctx, req, lark.WithUser(user-token)) ?
- 多租户
  - WithTenant(tenantKey).Drive().CreateFolder() ?
  - Drive().CreateFolder(ctx, req, lark.WithTenant(tenantKey)) ?
- 自定义机器人 https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN
- 消息卡片 builder
- 日历
- 文档
- 多维表格
- 会议室
- 会议
- 应用
- 邮箱
- 审批
- 服务台
- admin
- 实名
- ai
- 打卡
- okr
- 人事
