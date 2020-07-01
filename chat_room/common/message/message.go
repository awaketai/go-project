package message

// 确定消息类型
const (
	LoginMsgType  = "LoginMsg" // 登录消息类型
	LoginMsgResType = "LoginResMsg" // 登录返回消息类型
	RegisterType = "RegisterMsg" // 注册消息类型
	RegisterResMsgType = "RegisterResMsg" // 注册响应类型
	NotifyUserStatusMsgType = "NotifyUserStatusMsg" // 用户状态类型
	SmsMsgType = "SmsMsgType" // 发送群消息类型
)

// 用户状态常量
const (
	USER_ONLINE = iota
	USER_OFFLINE
	USER_BUSY
)


type Message struct {
	Type string `json:"type"`// 消息类型
	Data string `json:"data"`// 消息内容
}

// 登录属性结构
type LoginMsg struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResMsg struct {
	Code int `json:"code"` // 返回状态码 500表示未注册 200 登录成功
	UserIds []int // 保存用户id切片
	Error string `json:"error"`// 返回错误信息
}

// 注册类型
type RegisterMsg struct {
	User User `json:"user"`
}

// 注册响应类型
type RegisterResponseMsg struct {
	Code int `json:"code"` // 返回状态码 400：改用户已经占有 200 注册成功
	Error string `json:"error"`
}

// 配合服务器端推送用户状态变化的消息
type NotifyUserStatusMsg struct {
	UserId int `json:"userId"`
	Status int `json:"status"` // 用户状态
}

// 发送消息
type SmsMsg struct {
	Content string `json:"content"`// 消息内容
	User // 匿名结构体 当前包下的User
}

