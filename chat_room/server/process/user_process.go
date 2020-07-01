package process

import (
	"encoding/json"
	"fmt"
	"gostudy/chatroom/common/message"
	"gostudy/chatroom/common/utils"
	"gostudy/chatroom/server/model"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int // 标识该连接是那个用户的

}
// 登录逻辑请求处理
func (up *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	// 1.先从msg 中取出msg.data，并反序列化成LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data),&loginMsg)
	if err != nil {
		fmt.Println("login msg unmarshal error ",err)
		return
	}
	// 2.用户信息验证
	// 获取已经注册的用户并进行验证...
	var resMsg message.Message
	resMsg.Type = message.LoginMsgResType
	var loginResMsg message.LoginResMsg
	// 从数据库获取用户信息
	user,err := model.MyUserDao.LoginValidate(loginMsg.UserId,loginMsg.UserPwd)
	fmt.Println("获取用户信息",user)
	if err != nil {
		switch err {
		case model.ERROR_USER_WRONG_PWD:
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		case model.ERROR_USER_NOT_EXISTS:
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		default:
			loginResMsg.Code = 505
			loginResMsg.Error = "Unknow Error"
		}

		fmt.Println("获取用户信息失败",err)
		return
	}else{
		loginResMsg.Code = 200
		// 登录成功后，将该用户信息放入userManage中记录
		up.UserId = loginMsg.UserId
		userManager.AddOnlineUser(up)
		// 通知其他用户，当前用户已经上线
		up.NotifyOthersOnline(loginMsg.UserId)
		// 将当前在线用户放入userIds
		for id,_ := range userManager.onlineUsers {
			loginResMsg.UserIds = append(loginResMsg.UserIds,id)
		}
		fmt.Println("登录成功")
	}

	// 序列化
	data,err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("login result-1 marshal error",err)
		return
	}
	resMsg.Data = string(data)
	data,err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("login result-2 marshal error",err)
		return
	}

	// 发送操作
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 用户注册
func (this *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	var registerMs message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data),&registerMs)
	if err != nil {
		fmt.Println("register msg unmarshal error ",err)
		return
	}
	var regMsg message.Message
	regMsg.Type = message.RegisterResMsgType
	var registerResMsg = message.RegisterResponseMsg{}

	err = model.MyUserDao.UserRegister(&registerMs.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 505
			registerResMsg.Error = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMsg.Code = 506
			registerResMsg.Error = "注册发生未知错误"
		}

	}else{
		// 注册成功
		registerResMsg.Code = 200
	}
	data,err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("register msg marshal error ",err)
		return
	}
	msg.Data = string(data)
	data,err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json marshal err",err)
		return
	}
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	// 发送数据到服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册数据错误",err)
		return
	}
	// 获取服务器返回的数据
	regMsg,err = tf.ReadPkg()
	if err != nil {
		fmt.Println("获取注册数据错误",err)
		return
	}
	err = json.Unmarshal([]byte(msg.Data),&registerResMsg)
	if err != nil {
		fmt.Println("反序列化错误",err)
		return
	}
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功")
	}else{
		fmt.Println("注册失败")
	}
	return
}

// 通知所有在线用户方法
// userId int 当前登录用户id
func (this *UserProcess) NotifyOthersOnline(userId int)  {
	// 遍历onlineUsers然后一个个发送 NotifyUserStatusMsg
	for id,up := range userManager.onlineUsers {
		if id == userId {
			continue
		}
		// 通知其他用户
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int)  {
	data,err := this.GetNotifyMsg(userId)
	if err != nil {
		fmt.Println("获取推送消息结构体错误")
		return
	}
	// 发送数据
	tr := utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("推送在线用户错误",err)
	}
}

// 获取推送在线用户的消息结构体
func (this *UserProcess) GetNotifyMsg(userId int) (data []byte,err error) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType
	var notifyUserStatusMsg = message.NotifyUserStatusMsg{}
	notifyUserStatusMsg.UserId = userId
	notifyUserStatusMsg.Status = message.USER_ONLINE

	// 序列化
	data,err = json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("notify online error-1",err)
		return
	}
	msg.Data = string(data)
	data,err = json.Marshal(msg)
	if err != nil {
		fmt.Println("notify online error-2",err)
		return
	}
	return
}