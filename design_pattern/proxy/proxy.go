package proxy

import (
	"fmt"
	"log"
	"time"
)

type IUser interface {
	Login(username, password string)
}

type User struct{}

func (u *User) Login(username, password string) {
	fmt.Println("login...")
}

type UserProxy struct {
	user *User
}

func (u *UserProxy) Login(username, password string) {
	// before 一些统计逻辑
	start := time.Now()

	u.user.Login(username, password)

	// after 一些统计逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}
