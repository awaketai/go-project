package process

import "fmt"

// 用户登录在线管理

// 在服务器端有且只有一个，声名为全局变量
var userManager *UserManager

type UserManager struct {
	onlineUsers map[int]*UserProcess
}

// userManager初始化
func init()  {
	fmt.Println("初始化")
	userManager = &UserManager{
		make(map[int]*UserProcess,1024),
	}
}

// onlineUser添加/修改
func (this *UserManager) AddOnlineUser(up *UserProcess)  {
	this.onlineUsers[up.UserId] = up
}

// onlineUser删除
func (this *UserManager) DelOnlineUser(userId int)  {
	delete(this.onlineUsers,userId)
}

// 返回当前所有在线用户
func (this *UserManager) GetAllOnlineUser() map[int]*UserProcess {

	return this.onlineUsers
}

// 根据userId返回UseProcess
func (this *UserManager) GetOnlineUserById(userId int) (up *UserProcess,err error) {
	up,ok := this.onlineUsers[userId];
	if !ok {
		err = fmt.Errorf("用户 %d 不存在",userId)
		return
	}
	return up,nil
}
