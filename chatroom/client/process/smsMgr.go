package process

import (
	"encoding/json"
	"fmt"
	"goDemo/src/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//显示即可
	//反序列化mes.Data
	var smsMes message.SmsMes
	_ = json.Unmarshal([]byte(mes.Data), &smsMes)

	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
