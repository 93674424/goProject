package main

import (
	"fmt"
	"goDemo/src/chatroom/client/process"
	"os"
)

// 定义两个变量，用户id&password
var (
	userId   int
	userPwd  string
	userName string
)

func main() {

	//接收用户的选择
	var key int

	for true {
		fmt.Println("--------------欢迎登录多人聊天系统--------------")
		fmt.Println("\t\t\t  1 登录聊天室")
		fmt.Println("\t\t\t  2 注册用户")
		fmt.Println("\t\t\t  3 退出菜单")
		fmt.Println("\t\t\t 请选择(1-3)")

		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id:")
			_, _ = fmt.Scanln(&userId)
			fmt.Println("请输入用户pwd:")
			_, _ = fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//创建一个UserProcess实例
			up := &process.UserProcess{}
			_ = up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			_, _ = fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户昵称:")
			_, _ = fmt.Scanf("%s\n", &userName)
			//调用一个UserProcess实例,完成注册请求
			up := &process.UserProcess{}
			_ = up.Register(userId, userPwd, userName)

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

	}

	//根据用户输入显示新的提示信息
	//if key == 1 {
	//	//说明用户登录
	//
	//	//先将登录函数写到另一个文件login.go
	//	//重新调用
	//	login(userId, userPwd)
	//	if err != nil {
	//		fmt.Println("登陆失败")
	//	} else {
	//		fmt.Println("登录成功")
	//	}
	//} else if key == 2 {
	//	fmt.Println("进行用户注册逻辑。。。")
	//}
}
