package clls

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goDemo/src/chatroom/common/message"
	"net"
)

// 完成登录校验
func login(userId int, userPwd string) (err error) {

	//开始定协议...
	//fmt.Printf("userId = %d, userPwd = %s\n", userId, userPwd)
	//return nil

	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}

	//延时关闭
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("net.Conn err=", err)
		}
	}(conn)

	//准备它通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//把data赋给mes.Data字段
	mes.Data = string(data)

	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//这时data是要发送的消息
	//先把data长度发送给服务器
	//先获取到data长度->转成表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail = ", err)
		return
	}

	fmt.Printf("客户端发送消息长度为%d 内容=%s\n", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail = ", err)
		return
	}

	//time.Sleep(5 * time.Second)
	//fmt.Println("休眠了5秒...")

	//这里需要处理服务器端返回消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err", err)
		return
	}

	//将mes.Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return

}
