package srv

// websocket 服务端

import (
	"github.com/go-eyas/toolkit/websocket"
)

var WSServer *websocket.WS

func WSInit(conf *WSConfig) {
	WSServer = websocket.New(conf)
}
