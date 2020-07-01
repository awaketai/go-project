package process

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
	"net"
)

// 消息发送管理
type SmsProcess struct {

}

// 对接收到的用户消息进行发送
func (this *SmsProcess) SendGroupMsg(msg *message.Message)  {
	// 遍历所有在线用户，发送消息
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data),&smsMsg)
	if err != nil {
		fmt.Println("发序列化用户消息失败")
		return
	}
	data,err := json.Marshal(msg)
	if err != nil {
		fmt.Println("发送消息序列化失败",err)
		return
	}
	for id,up := range userManager.onlineUsers {

		if id == smsMsg.UserId {
			continue
		}
		this.SendMsgToOnlineUser(data,up.Conn)
	}
}

// 转发消息到在线用户
func (this *SmsProcess) SendMsgToOnlineUser(data []byte,conn net.Conn)  {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送群消息失败",err)
	}
}