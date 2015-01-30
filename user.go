package jiaoxue

import (
	"github.com/googollee/go-socket.io"
)

type User struct {
	userId   string
	userName string
	courseId string
	role     string

	client_socket socketio.Socket
	isRecording   bool
}

func NewUser(client_socket socketio.Socket) *User {
	user := new(User)
	user.client_socket = client_socket

	NewCmdSystem().Bind(user)

	return user
}

type UserManager struct {
	UserIdUserMap map[string]*User
	UserCountMap  map[string]int
}

func NewUserManager() *UserManager {
	user_manager := new(UserManager)
	user_manager.UserIdUserMap = make(map[string]*User)
	user_manager.UserCountMap = make(map[string]int)
	return user_manager
}

func (this *UserManager) UserCount(course_id string) int {
	if this.UserCountMap != nil {
		count, exists := this.UserCountMap[course_id]
		if !exists {
			return 0
		} else {
			return count
		}
	}

	return 0
}

func (this *UserManager) UserCountAdd(course_id string) int {
	if this.UserCountMap != nil {
		_, exists := this.UserCountMap[course_id]
		if !exists {
			this.UserCountMap[course_id] = 1
		} else {
			this.UserCountMap[course_id] = this.UserCountMap[course_id] + 1
		}

		return this.UserCountMap[course_id]
	}

	return 0
}

func (this *UserManager) UserCountSub(course_id string) int {
	if this.UserCountMap != nil {
		_, exists := this.UserCountMap[course_id]
		if !exists {
			this.UserCountMap[course_id] = 0
		} else {
			if this.UserCountMap[course_id] > 0 {
				this.UserCountMap[course_id] = this.UserCountMap[course_id] - 1
			}
		}

		return this.UserCountMap[course_id]
	}

	return 0
}

func (this *UserManager) AddUser(userId string, user *User) {
	this.UserIdUserMap[userId] = user
}

func (this *UserManager) RemoveUser(userId string) {
	delete(this.UserIdUserMap, userId)
}

func (this *UserManager) HasUser(userId string) *User {
	user, _ := this.UserIdUserMap[userId]

	return user
}
