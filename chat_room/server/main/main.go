package main

import (
	"fmt"
	"gostudy/chatroom/server/model"
	"net"
)

func main()  {
	fmt.Println("监听端口：8889")
	// 服务器启动时，就去初始化redis连接池
	initPool("127.0.0.1:6379",8,0,100)
	// 初始化userDao
	initUserDao()

	listen,err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err",err)
		return
	}
	for {
		fmt.Println("等待连接...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error :",err)
			continue
		}
		// 一旦连接成功，启动协程和客户端保持通讯
		go process(conn)
	}
}

func process(conn net.Conn) {
	// 读取客户端发送的信息
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}

	err := processor.ProcessDeal()
	if err != nil {
		fmt.Println("deal process error",err)
	}
}

// 初始化userDao pool是全局变量
func initUserDao()  {
	model.MyUserDao = model.NewUserDao(pool)
}