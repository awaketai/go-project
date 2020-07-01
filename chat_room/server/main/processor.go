package main

import (
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
	process2 "gostudy/chatroom/server/process"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}
// 根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessDeal(msg *message.Message) (err error)  {
	fmt.Println("客户端发送的消息：",msg)
	switch msg.Type {
	case message.LoginMsgType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(msg)
		// 处理登录逻辑
	case message.RegisterType:
		// 登录注册逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(msg)
	case message.SmsMsgType:
		// 转发群聊消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMsg(msg)

	default:
		fmt.Println("消息类型不存在，无法处理...")

	}
	return
}

// 处理主函数接收到的请求
func (this *Processor) ProcessDeal() (err error) {
	for{
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		msg,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				// 如果用户退出，将其从在线用户中删除

				fmt.Println("客户端退出...")
				return err
			}else{
				fmt.Println("read pkg error",err)
				return err
			}
		}
		err = this.serverProcessDeal(&msg)
		if err != nil {
			fmt.Println("err",err)
			return err
		}
	}
}
