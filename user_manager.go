package jiaoxue

import (
	"log"
)

type UserManager struct {
	UserIdUserMap map[string]*User
}

func NewUserManager() *UserManager {
	log.Println("NewUserManager")

	user_manager := new(UserManager)
	user_manager.UserIdUserMap = make(map[string]*User)
	return user_manager
}

func (this *UserManager) UserCount() int {
	log.Println("UserManager UserCount")

	return len(this.UserIdUserMap)
}

func (this *UserManager) AddUser(userId string, user *User) {
	log.Println("UserManager AddUser", userId)

	this.UserIdUserMap[userId] = user
}

func (this *UserManager) RemoveUser(userId string) {
	log.Println("UserManager RemoveUser", userId)

	delete(this.UserIdUserMap, userId)
}

func (this *UserManager) HasUser(userId string) *User {
	log.Println("UserManager HasUser", userId)

	user, _ := this.UserIdUserMap[userId]

	return user
}
