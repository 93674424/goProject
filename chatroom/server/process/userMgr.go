package process2

import "fmt"

// UserMgr UserMgr实例有且只有一个将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对UserMgr初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUser 完成对onlineUser添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// DelOnlineUser del
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// GetAllOnlineUsers 返回所有在线用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

// GetOnlineUserById 根据id返回对应值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从map中取出一个值带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok {
		//查找的用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
