package process

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
)

// 消息进程
type SmsProcess struct {

}

// 发送群聊消息
func (this *SmsProcess) SendGroupSms(content string) (err error) {
	var msg = message.Message{}
	msg.Type = message.SmsMsgType

	var smsMsg = message.SmsMsg{}
	smsMsg.Content = content
	smsMsg.UserId = currentUser.UserId
	smsMsg.UserStatus = currentUser.UserStatus

	// 序列化
	data,err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("发送消息序列化失败-1")
		return
	}
	msg.Data = string(data)
	data,err = json.Marshal(msg)
	if err != nil {
		fmt.Println("发送消息序列化失败-1")
		return
	}
	// 发送数据
	tf := utils.Transfer{
		Conn: currentUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送消息失败",err)
		return
	}
	return
}
