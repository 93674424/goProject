package process2

import (
	"encoding/json"
	"fmt"
	"goDemo/src/chatroom/common/message"
	"goDemo/src/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

// SendGroupMes 转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端onlineUsers map,将消息转发出去
	//取出mes内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	data, _ := json.Marshal(mes)
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	_ = tf.WritePkg(data)
}
