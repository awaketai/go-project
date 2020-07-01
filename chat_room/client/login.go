package main

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
	"net"
)

type UserLogin struct {

}

/**
用户登录
 */
func login(userId int,userPwd string) (err error) {

	// 连接服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("连接错误：",err)
		return
	}
	defer conn.Close()
	data := msgDeal(userId,userPwd,"")
	// 发送消息到服务器
	// 先发送消息长度，然后发送消息内容
	// 先获取到消息长度，然后将其转换成为一个切片
	processor := utils.Transfer{
		Conn: conn,
	}
	err = processor.WritePkg(data)
	if err != nil {
		fmt.Println("send data error",err)
		return
	}
	// 处理服务端返回的消息
	msg,err := processor.ReadPkg()
	if err != nil {
		fmt.Println("read data error",err)
		return
	}
	// 对接收到的数据反序列化
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data),&loginResMsg)
	fmt.Println(loginResMsg)
	if err != nil {
		fmt.Println("resolve login result json unmarshal error",err)
		return
	}
	if loginResMsg.Code == 200 {
		fmt.Println("login success")
	}else{
		fmt.Println("login failed",err)

	}
	return
}

/**
	要发送给服务器的消息整理
 */
func msgDeal(userId int,userPwd,userName string) []byte {
	// 发送消息到服务器
	var msg message.Message
	msg.Type = message.LoginMsgType
	// 登录消息内容
	var loginMsg = message.LoginMsg{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}
	data,err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal error-1",err)
		return []byte{}
	}
	msg.Data = string(data)
	// 对msg进行序列化
	data,err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal error-2",err)
		return []byte{}
	}
	// 可以发送的消息
	return data
}
