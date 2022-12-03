package process

import (
	"fmt"
	"goDemo/src/chatroom/client/model"
	"goDemo/src/chatroom/common/message"
)

// 客户端维护的map
var onlineUsers = make(map[int]*message.User, 10)

// CurUser 在用户登录成功后完成对CurUser初始化
var CurUser model.CurUser

// 客户端显示当前在线用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	//原来没有
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
