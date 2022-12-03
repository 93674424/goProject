package process

import (
	"encoding/json"
	"fmt"
	"goDemo/src/chatroom/client/utils"
	"goDemo/src/chatroom/common/message"
)

type SmsProcess struct {
}

// SendGroupMes 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err =", err)
		return
	}

	mes.Data = string(data)

	//对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err =", err)
		return
	}

	//将序列化后的mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err =", err)
		return
	}
	return
}
