package message

// User 定义一个用户结构体
type User struct {
	//确定字段信息
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态...
}
