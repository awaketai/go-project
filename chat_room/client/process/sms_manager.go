package process

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
)

// 将群发消息显示到各个客户端
func outputGroupMsg(msg *message.Message)  {
	// 1.反序列化
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data),&smsMsg)
	if err != nil {
		fmt.Println("群消息反序列化失败555",err.Error())
		return
	}
	// 显示信息
	str := fmt.Sprintf("用户ID：\t%d 对大家说：%s",smsMsg.UserId,smsMsg.Content)
	fmt.Println(str)
}
