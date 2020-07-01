package process

import (
	"fmt"
	"gostudy/chatroom/client/model"
	"gostudy/chatroom/common/message"
)

// 客户端要维护的map
var onlineUsers map[int] *message.User = make(map[int] *message.User,10)
// 全局变量当前用户信息，登录完成后完成对其初始化
var currentUser model.CurrentUser
// 处理返回的用户状态信息
func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg)  {
	user,ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		// 如果不存在才创建
		user = &message.User{
			UserId: notifyUserStatusMsg.UserId,
			UserStatus: notifyUserStatusMsg.Status,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId] = user

	outputOnlineUser()
}

// 客户端显示当前在线的用户
func outputOnlineUser()  {
	for id,_ := range onlineUsers {
		fmt.Println("用户id:\t",id)
	}
}

