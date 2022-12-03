package model

import (
	"goDemo/src/chatroom/common/message"
	"net"
)

// CurUser 在客户端很多地方会使用到curUser将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
