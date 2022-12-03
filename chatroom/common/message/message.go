package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// UserOnline 定义用户状态常量
const (
	UserOnline = iota
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

// LoginMes 定义两个消息，后面需要再增加
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code    uint   `json:"code"` //返回状态码 500表示该用户未注册 200 表示登录成功
	UsersId []int  //增加字段：保存用户id的切片
	Error   string `json:"error"` //返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code  uint   `json:"code"`  //返回状态码 400 表示该用户已经占用 200表示注册成功
	Error string `json:"error"` //返回错误信息
}

// NotifyUserStatusMes 配合服务器端推送用户上线状态变化消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

// SmsMes 增加一个SmsMes,发送的
type SmsMes struct {
	Content string        `json:"content"` //内容
	User    `json:"user"` //匿名结构体
}
