package message

type User struct {
	// 因为存储的格式是userId,userName这种方式，为了反序列化成功，
	// 因此用户信息的json的字符串的key和结构体的字段对应的tag名字一致
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"` // 用户状态 eg。1.在线 2. 离线等
}







