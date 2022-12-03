package main

import (
	"errors"
	"fmt"
	"goDemo/src/chatroom/common/message"
	"goDemo/src/chatroom/server/process"
	"goDemo/src/chatroom/server/utils"
	"io"
	"net"
)

// Processor 创建Processor结构体
type Processor struct {
	Conn net.Conn
}

// 根据客户端发送消息种类不同，决定调用哪个函数处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//看看从客户端是否能接收到群发消息
	//fmt.Println("mes =", mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//处理转发消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func (this *Processor) process2() (err error) {
	//读取客户端发送的信息
	for {
		//将读取数据包封装成函数readPkg(),return Message,Err
		//创建一个Transfer实例，完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也正常退出。。。")
				return err
			} else {
				err = errors.New("readPkg read error")
				return err
			}
		}
		fmt.Println("mes =", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
