package process2

import (
	"encoding/json"
	"fmt"
	"goDemo/src/chatroom/common/message"
	"goDemo/src/chatroom/server/model"
	"goDemo/src/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段有...
	Conn net.Conn
	//表示该conn是哪个用户
	UserId int
}

// ServerProcessRegister 专门处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//先从mes取出mes.Date反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	//申明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//使用model.MyUserDao到redis完成验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ErrorUserExists {
			registerResMes.Code = 505
			registerResMes.Error = model.ErrorUserExists.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}
	//将data赋值给resMes
	resMes.Data = string(data)

	//resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//发送data,封装到writePkg函数中
	//因为使用分层模式(mvc),先创建Transfer实例再读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

// ServerProcessLogin 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	//核心代码...
	//先从mes取出mes.Date反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	//申明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//申明一个loginResMes,完成赋值
	var loginResMes message.LoginResMes

	//使用model.MyUserDao到redis完成验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNoExists {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200
		//将登陆成功的用户放入userMgr中
		//将登陆成功的用户userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户id放入到loginResMes.UsersId
		//遍历userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登陆成功")
	}

	////如果用户id = 100，pwd=123456,认为合法，否则不合法
	//
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//合法
	//	loginResMes.Code = 200
	//} else {
	//	//不合法
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在，请注册再使用..."
	//}

	//将 loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}

	//将data赋值给resMes
	resMes.Data = string(data)

	//resMes序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//发送data,封装到writePkg函数中
	//因为使用分层模式(mvc),先创建Transfer实例再读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

// NotifyOthersOnlineUser 专门处理userId通知其他在线用户
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍历onlineUsers,然后逐个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知,单独写方法
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的NotifyUserStatusMes
	var mes = message.Message{
		Type: message.NotifyUserStatusMesType,
	}
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes json.Marshal err =", err)
		return
	}
	mes.Data = string(data)
	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}
	//发送,创建transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}
