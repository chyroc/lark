# lark

[![codecov](https://codecov.io/gh/chyroc/lark/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark "go report card")](https://goreportcard.com/report/github.com/chyroc/lark)
[![test status](https://github.com/chyroc/lark/actions/workflows/go.yml/badge.svg)](https://github.com/chyroc/lark/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark)

Feishu/Lark Open API Go Sdk Written Casually During Non-Working Hours.

Created By Code Generation.

## Install

```shell
go get github.com/chyroc/lark
```

## Usage

### Quick Start

create lark client and create chat:

```go
package main

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"

	"github.com/chyroc/lark"
)

func main() {
	ctx := context.Background()
	r := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

	resp, _, err := r.Chat().CreateChat(ctx, &lark.CreateChatReq{
		Name: ptr.String("chat-name"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("chat created: ", resp.ChatID)
}
```

### Docs

https://godoc.org/github.com/chyroc/lark

## Todo

- [ ] generate all api
- [x] test coverage >= 80%
- [ ] add english comment, from larksuite webpage
- [ ] struct data
