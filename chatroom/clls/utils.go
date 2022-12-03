package clls

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"goDemo/src/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("等待客户端发送的数据...")

	//conn.read在conn没有close情况下才会阻塞
	//如果客户端关闭conn，则不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	//根据buf[:4]转成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	//把pkgLen反序列化成message.Message &mes!!!
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json.Unmarshal error")
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {

	//先发送一个长度给对方
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

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail = ", err)
		return
	}
	return
}
