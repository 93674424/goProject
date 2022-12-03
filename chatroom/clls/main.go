package clls

import (
	"fmt"
)

// 定义两个变量，用户id&password
var (
	userId  int
	userPwd string
)

func main() {

	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("--------------欢迎登录多人聊天系统--------------")
		fmt.Println("\t\t\t  1 登录聊天室")
		fmt.Println("\t\t\t  2 注册用户")
		fmt.Println("\t\t\t  3 退出菜单")
		fmt.Println("\t\t\t 请选择(1-3)")

		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			fmt.Println(err)
		}
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

	}

	//根据用户输入显示新的提示信息
	if key == 1 {
		//说明用户登录
		fmt.Println("请输入用户id:")
		_, err := fmt.Scanln(&userId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("请输入用户pwd:")
		_, err = fmt.Scanf("%s\n", &userPwd)
		if err != nil {
			fmt.Println(err)
		}
		//先将登录函数写到另一个文件login.go
		_ = login(userId, userPwd)
		//if err != nil {
		//	fmt.Println("登陆失败")
		//} else {
		//	fmt.Println("登录成功")
		//}
	} else if key == 2 {
		fmt.Println("进行用户注册逻辑。。。")
	}
}
