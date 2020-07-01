package process

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
	"net"
	"os"
)

// 显示登录成功后的界面...
func ShowMenu()  {
	fmt.Println("----恭喜%s登录成功----")
	fmt.Println("----1.显示在线用户列表----")
	fmt.Println("----2.发送消息-----")
	fmt.Println("----3.信息列表----")
	fmt.Println("----4.退出系统----")

	fmt.Println("请选择1-4")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息,想要说什么？")
		fmt.Scanf("%s\n",&content)
		err := smsProcess.SendGroupSms(content)
		if err != nil {
			fmt.Println("发送消息失败",err)
		}
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		// 退出前给服务器发送消息，进行一些其他处理
		os.Exit(0)
	default:
		fmt.Println("无效的输入")
	}
}

// 和服务器端保持通信，服务端信息处理
func serverProcessMsg(conn net.Conn)  {
	// 创建transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{Conn: conn}
	for{
		msg,err := tf.ReadPkg()
		fmt.Println("read server message...")
		if err != nil{
			fmt.Println("read message error",err)
			continue
		}
		// 读取到消息处理
		//fmt.Printf("msg = %v",msg)

		switch msg.Type {
		case message.NotifyUserStatusMsgType:
			// 有人上线
			// 1.取出 NotifyUserStatusMsg
			// 2.把这个用户的信息，状态保存到客户端map中 map[int]User
			var notifyUserStatusMsg message.NotifyUserStatusMsg
			err := json.Unmarshal([]byte(msg.Data),&notifyUserStatusMsg)
			if err != nil {
				fmt.Println("反序列化失败")
			}
			updateUserStatus(&notifyUserStatusMsg)
		case message.SmsMsgType:
			// 群发消息
			outputGroupMsg(&msg)
		default:
			fmt.Println("类型无法识别")
		}
	}
}
