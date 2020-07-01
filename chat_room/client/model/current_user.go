package model

import (
	"gostudy/chatroom/common/message"
	"net"
)

/**
	当前用户信息
 */

type CurrentUser struct {
	Conn net.Conn
	message.User
}
