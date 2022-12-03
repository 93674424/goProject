package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goDemo/src/chatroom/common/message"
)

// MyUserDao 在服务器启动后就初始化一个userDao实例
// 做成全局变量，在需要与redis操作时直接使用即可
var (
	MyUserDao *UserDao
)

// UserDao 定义一个UserDao结构体完成对User结构体各种操作
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定的id到redis中查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//表示在users哈希中不存在对应id
		if err == redis.ErrNil {
			err = ErrorUserNoExists
		}
		return
	}
	user = &User{}
	//将res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// Login 完成登录校验
// 完成对用户验证 用户名和密码正确返回User实例
// 用户名或密码错误返回对应错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从userDao连接池中取一根连接
	conn := this.pool.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//这时用户名获取到了，判断密码是否正确
	if user.UserPwd != userPwd {
		err = ErrorUserPwd
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	//先从userDao连接池中取一根连接
	conn := this.pool.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ErrorUserExists
		return
	}

	//这时id未注册，进行注册操作
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
