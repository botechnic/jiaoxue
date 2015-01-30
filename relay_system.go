package jiaoxue

import (
	"encoding/json"
	"log"
)

type RelaySystem struct {
}

func NewRelaySystem() *RelaySystem {
	log.Println("NewRelaySystem")

	relay_system := new(RelaySystem)

	return relay_system
}

func (this *RelaySystem) Start(user *User) bool {
	log.Println("RelaySystem Start")

	return true
}

func (this *RelaySystem) Stop(user *User) bool {
	log.Println("RelaySystem Stop")

	return true
}

func (this *RelaySystem) Relay(user *User, cmd_type string, msg string) bool {
	switch cmd_type {
	case "login":
		this.login(user, cmd_type, msg)
	case "logout":
		this.logout(user, cmd_type, msg)
	case "disconnect":
		this.disconnect(user, cmd_type, msg)
	}

	return true
}

func (this *RelaySystem) StartRecord(user *User) {
	log.Println("User StartRecord")

	GetContext().record_manager.Start(user)
}

func (this *RelaySystem) StopRecord(user *User) {
	log.Println("User StopRecord")

	GetContext().record_manager.Stop(user)
}

func (this *RelaySystem) StartLive(user *User) {
	log.Println("User StartLive")

	GetContext().cmd_manager.Start(user)
}

func (this *RelaySystem) StopLive(user *User) {
	log.Println("User StopLive")

	GetContext().cmd_manager.Stop(user)
}

func (this *RelaySystem) send_user_count(user *User, cmd_type string, msg string) {
	user_count := GetContext().user_manager.UserCount(user.courseId)
	user_count_msg := new(UserCountMsg)
	user_count_msg.CourseId = user.courseId
	user_count_msg.UserCount = user_count
	_, err := json.Marshal(user_count_msg)
	if err != nil {
		panic(err.Error())
	}
	//user.client_socket.Emit("user count", user_count)
	//user.client_socket.BroadcastTo(user.courseId, "user count", user_count)
	GetContext().live_server.server.BroadcastTo(user.courseId, "user count", user_count)
	log.Println("user count", user_count)
}

func (this *RelaySystem) login(user *User, cmd_type string, msg string) {
	log.Println("User Login", msg)

	var login_msg LoginMsg
	err := json.Unmarshal([]byte(msg), &login_msg)

	if err != nil {
		log.Fatalln("parse login message error :%v", err)
	} else {
		user.userId = login_msg.UserName
		user.userName = login_msg.UserName
		user.courseId = login_msg.CourseId
		user.role = login_msg.Role

		user.client_socket.Join(user.courseId)
		GetContext().user_manager.UserCountAdd(user.courseId)
		this.StartLive(user)

		this.send_user_count(user, cmd_type, msg)
	}

}

func (this *RelaySystem) logout(user *User, cmd_type string, msg string) {
	log.Println("User Logout")

	user.client_socket.Leave(user.courseId)
	GetContext().user_manager.UserCountSub(user.courseId)
	this.StopLive(user)
	this.StopRecord(user)

	user.courseId = "0"

	this.send_user_count(user, cmd_type, msg)
}

func (this *RelaySystem) disconnect(user *User, cmd_type string, msg string) {
	log.Println("User Disconnect")

	if user.courseId != "0" {
		this.logout(user, cmd_type, msg)
	}

	GetContext().user_manager.RemoveUser(user.client_socket.Id())

	this.send_user_count(user, cmd_type, msg)
}

func (this *RelaySystem) connect(user *User, cmd_type string, msg string) {
	this.StartLive(user)
}
