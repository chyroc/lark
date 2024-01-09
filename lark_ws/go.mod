module github.com/chyroc/lark/lark_ws

go 1.20

replace github.com/chyroc/lark => ../

require (
	github.com/chyroc/lark v0.0.112
	github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket v1.5.1
)

require golang.org/x/net v0.17.0 // indirect
