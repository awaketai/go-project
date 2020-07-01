package process

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
	"net"
)

type UserProcess struct {
	// 暂时不需要任何字段
	userId int
	userPwd string
	userName string
}
/**
用户登录
*/
func (this *UserProcess) Login(userId int,userPwd string) (err error) {

	// 连接服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("连接错误：",err)
		return
	}
	defer conn.Close()
	data := msgDeal(userId,userPwd,"asher")
	// 发送消息到服务器
	// 先发送消息长度，然后发送消息内容
	// 先获取到消息长度，然后将其转换成为一个切片
	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send data error",err)
		return
	}
	// 处理服务端返回的消息
	msg,err := tf.ReadPkg()
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
		// 初始化currentUser
		currentUser.Conn = conn
		currentUser.UserId = userId
		currentUser.UserStatus = message.USER_ONLINE

		// 如果登录成功，显示登录成功后的菜单

		// 启动协程：保持和服务端的通信，如果服务器有数据推送给客户端
		// 则接收并显示在客户端的终端

		// 登录成功后显示在线用户
		fmt.Println("当前在线用户：")
		for _,v := range loginResMsg.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户id：\t",v)
			// 客户端的onlineUser初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.USER_ONLINE,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		go serverProcessMsg(conn)

		for{
			ShowMenu()
		}
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

// 用户注册
func (this *UserProcess) Register(userId int,userPwd,userName string) (err error) {
	// 连接服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("连接错误：",err)
		return
	}
	defer conn.Close()

	var msg = message.Message{}
	msg.Type = message.RegisterType
	var registerMsg = message.RegisterMsg{}
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName

	// 序列化数据
	data,err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("注册序列化失败-1",err)
		return
	}
	msg.Data = string(data)
	data,err = json.Marshal(msg)
	if err != nil {
		fmt.Println("注册序列化失败-2",err)
		return
	}
	transfer := utils.Transfer{
		Conn: conn,
	}
	// 发送数据到服务器
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送数据失败",err)
		return
	}
	// 获取服务器返回的数据
	res,err := transfer.ReadPkg()
	if err != nil {
		fmt.Println("获取数据错误：",err)
		return
	}
	var regResMsg = message.RegisterResponseMsg{}
	err = json.Unmarshal([]byte(res.Data),&regResMsg)
	fmt.Println(regResMsg)
	if err != nil {
		fmt.Println("初测反序列化失败-3")
		return
	}
	if regResMsg.Code == 200 {
		fmt.Println("注册成功，重新登录")

	}else{
		fmt.Println("注册失败",regResMsg.Error)
	}
	return
}
