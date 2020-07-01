package main

import (
	"fmt"
	"gostudy/chatroom/client/process"
	"os"
)

var (
	userId int
	userPwd string
	userName string
)

func main()  {
	// 接收用户选择，
	// 判断是否继续显示菜单
	var key int
	var loop = true
	for loop {
		fmt.Println("欢迎登陆多人聊天系统")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)：")

		// 获取输入
		fmt.Scanf("%d \n",&key)
		switch key {
		case 1:
			// 用户登录
			fmt.Println("请输入用户ID：")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n",&userPwd)

			loginDeal := &process.UserProcess{}
			err := loginDeal.Login(userId,userPwd)
			if err != nil {
				fmt.Println("登录失败",err)
				continue
			}
			fmt.Println("登陆聊天室")
			loop = false // 如果登陆，则退出循环
		case 2:
			fmt.Println("注册聊天室...")
			fmt.Println("请输入用户ID：")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Println("请输入用户昵称：")
			fmt.Scanf("%s\n",&userName)
			regDeal := &process.UserProcess{}

			regDeal.Register(userId,userPwd,userName)
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			os.Exit(1)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
