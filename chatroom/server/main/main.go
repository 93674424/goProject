package main

import (
	"fmt"
	"goDemo/src/chatroom/server/model"
	"net"
	"time"
)

//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	buf := make([]byte, 8096)
//	fmt.Println("等待客户端发送的数据...")
//
//	//conn.read在conn没有close情况下才会阻塞
//	//如果客户端关闭conn，则不会阻塞
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//		return
//	}
//	//根据buf[:4]转成uint32类型
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4])
//
//	//根据pkgLen读取消息内容
//	n, err := conn.Read(buf[:pkgLen])
//	if n != int(pkgLen) || err != nil {
//		return
//	}
//
//	//把pkgLen反序列化成message.Message &mes!!!
//	err = json.Unmarshal(buf[:pkgLen], &mes)
//	if err != nil {
//		err = errors.New("json.Unmarshal error")
//		return
//	}
//	return
//}
//
//func writePkg(conn net.Conn, data []byte) (err error) {
//
//	//先发送一个长度给对方
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
//	//发送长度
//	n, err := conn.Write(buf[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(bytes) fail = ", err)
//		return
//	}
//
//	//发送data本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write(bytes) fail = ", err)
//		return
//	}
//	return
//}

//// 专门处理登录请求
//func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//
//	//核心代码...
//	//先从mes取出mes.Date反序列化成LoginMes
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal err = ", err)
//		return
//	}
//
//	//申明一个resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	//申明一个loginResMes,完成赋值
//	var loginResMes message.LoginResMes
//	//如果用户id = 100，pwd=123456,认为合法，否则不合法
//
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		//合法
//		loginResMes.Code = 200
//	} else {
//		//不合法
//		loginResMes.Code = 500
//		loginResMes.Error = "该用户不存在，请注册再使用..."
//	}
//
//	//将 loginResMes序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail", err)
//	}
//
//	//将data赋值给resMes
//	resMes.Data = string(data)
//
//	//resMes序列化，准备发送
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail", err)
//		return
//	}
//
//	//发送data,封装到writePkg函数中
//	err = writePkg(conn, data)
//	return
//
//}

//// 根据客户端发送消息种类不同，决定调用哪个函数处理
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录
//		err = serverProcessLogin(conn, mes)
//		if err != nil {
//			fmt.Println("serverProcessLogin fail =", err)
//			return
//		}
//	case message.RegisterMesType:
//		//处理注册
//	default:
//		fmt.Println("消息类型不存在，无法处理。。。")
//	}
//	return
//}

// 处理和客户端通讯
func process(conn net.Conn) {

	//延时关闭conn
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("net.Conn err", err)
		}
	}(conn)

	//调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误=", err)
		return
	}

	////读取客户端发送的信息
	//for {
	//	//将读取数据包封装成函数readPkg(),return Message,Err
	//	mes, err := readPkg(conn)
	//	if err != nil {
	//		if err == io.EOF {
	//			fmt.Println("客户端退出，服务器端也正常退出。。。")
	//			return
	//		} else {
	//			err = errors.New("readPkg read error")
	//			return
	//		}
	//	}
	//	fmt.Println("mes =", mes)
	//
	//	err = serverProcessMes(conn, &mes)
	//	if err != nil {
	//		return
	//	}
	//}
}

func init() {
	//当服务器启动则初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

// 完成对UserDao初始化任务
func initUserDao() {
	//pool是全局变量
	//初始化顺序：initPool -> initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	//提示信息
	fmt.Println("服务器[新架构]在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("net.Listener err =", err)
		}
	}(listen)
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦连接成功启动一个协程和客户端保持通讯...
		go process(conn)
	}
}
