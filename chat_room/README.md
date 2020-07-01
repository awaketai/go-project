基于golang的简单聊天室
----

### 使用

本地路径在 ~/go/src/gostudy/chatroom下

编译：

```
cd ~/go/src
go build -o client.exe ./gostudy/chatroom/client/main
go build -o server.exe ./gostudy/chatroom/server/main
```

使用redis存储用户信息，存储格式：
    
    hash_key: users 
    hash_field:userId eg.100
    hash_value: eg.{"userId":100,"userPwd":123456,"userName":"asher"}
    
### 客户端：

main.go:

+ 显示第一级菜单

+ 根据用户输入，调用相应的处理器处理

sms_process.go

+ 处理和短消息相关的逻辑

+ 私聊

+ 群发

user_process.go

+ 处理和用户相关的业务

+ 登录、注册等

utils.go

常用的工具函数、结构体类

server.go

+ 显示登录成功界面

+ 保持和服务器进行通讯(协程)

+ 当服务器端有消息的时候，显示在客户端

### 服务端

main.go

+ 监听

+ 等待客户端连接

+ 初始化工作

process.go

总的处理器，根据客户端的请求，调用相应的处理器完成任务

sms_process.go

+ 处理和短消息相关的业务

+ 群聊、点对点聊天

user_process.go

+ 处理和用户相关的请求

+ 登录、注册、注销、用户列表管理等

user_manager.go

+ 维护用户在线列表





