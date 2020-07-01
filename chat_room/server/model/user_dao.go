package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gostudy/chatroom/common/message"
	"strconv"
)

var MyUserDao *UserDao // 全局变量

// 使用工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		RedisPool: pool,
	}
	return userDao
}
// 定义一个user dao
// 对user结构体的各种操作
type UserDao struct {
	RedisPool *redis.Pool
}

// 根据userid返回User实例
func (this *UserDao) getUserById(conn redis.Conn,userId int) (user *User,err error) {
	// 根据指定id去查询用户
	//res,err := conn.Do("AUTH","dXaogkwwrHYoljok")
	//if err != nil {
	//	fmt.Println("auth failed redis",res)
	//	return
	//}
	fmt.Println("userId:",userId,"user_" + strconv.Itoa(userId))
	user = &User{}
	// string函数 int转str问题
	data,err := redis.String(conn.Do("HGet","users" ,userId))
	fmt.Println("redis - result",data,err)
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}
	// 反序列化
	err = json.Unmarshal([]byte(data),user)
	if err != nil {
		fmt.Println("用户信息反序列化失败",err)
		return
	}
	return
}
// 用户登录验证
func (this *UserDao) LoginValidate(userId int,userPwd string) (user *User,err error){
	// 先从连接池中获取连接
	conn := this.RedisPool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_WRONG_PWD
		return
	}
	return
}

// 用户信息存储格式：{"userId":100,"userPwd":123456,"userName":"asher"}

func (this *UserDao) UserRegister(user *message.User) (err error) {
	conn := this.RedisPool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	if err == nil {
		// 如果用户已经存在
		err = ERROR_USER_EXISTS
		return
	}
	// 注册用户
	data,err := json.Marshal(user)
	fmt.Println(string(data))
	if err != nil {
		return
	}
	_,err = redis.Int64(conn.Do("HSet","users",user.UserId,string(data)))
	if err != nil {
		fmt.Println("注册保存出错",err)
		return
	}
	return
}