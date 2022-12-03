package process

import (
	"encoding/json"
	"fmt"
	"goDemo/src/chatroom/client/utils"
	"goDemo/src/chatroom/common/message"
	"net"
	"os"
)

// ShowMenu 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("------恭喜???登陆成功------")
	fmt.Println("------1.显示在线用户列表----")
	fmt.Println("------2.发送消息-----------")
	fmt.Println("------3.信息列表-----------")
	fmt.Println("------4.退出系统-----------")
	fmt.Println("------请选择(1-4)----------")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	_, _ = fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你想对大家说点什么：")
		_, _ = fmt.Scanf("%s\n", &content)
		_ = smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("查看信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("你的输入有误请重新输入")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个Transfer实例，不停读取服务器发送的消息
	tf := &utils.Transfer{Conn: conn}
	for {
		fmt.Println("客户端正在等待服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑。。。
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线处理
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			_ = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2.把上线用户状态保存到客户端map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: //群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器段返回一个未知消息类型...")
		}
		//fmt.Printf("mes = %v", mes)
	}
}
