package jiaoxue

import (
	"encoding/json"
	"log"
	"strings"
)

type LoginMsg struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	CourseId string `json:"course_id"`
}

type UserCountMsg struct {
	CourseId  string `json:"course_id"`
	UserCount int    `json:"numUsers"`
}

type CmdSystem struct {
}

func NewCmdSystem() *CmdSystem {
	cmd_system := new(CmdSystem)
	return cmd_system
}

func (this *CmdSystem) Bind(user *User) bool {
	// system
	user.client_socket.On("login", func(msg string) {
		this.on_login(user, "login", msg)
	})

	user.client_socket.On("logout", func(msg string) {
		this.on_logout(user, "logout", msg)
	})

	user.client_socket.On("disconnection", func() {
		this.on_disconnection(user, "disconnection", "")
	})

	// ppt
	user.client_socket.On("prev", func(msg string) {
		this.on_prev(user, "prev", msg)
	})

	user.client_socket.On("next", func(msg string) {
		this.on_next(user, "next", msg)
	})

	// chat
	user.client_socket.On("chat message", func(msg string) {
		this.on_chat_message(user, "chat message", msg)
	})

	// record
	user.client_socket.On("start_record", func(msg string) {
		this.on_start_record(user, "start_record", msg)
	})

	user.client_socket.On("stop_record", func(msg string) {
		this.on_stop_record(user, "stop_record", msg)
	})

	return true
}

// system
func (this *CmdSystem) on_login(user *User, cmd_type string, msg string) {
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
		user_manager.UserCountAdd(user.courseId)

		this.relay_user_count(user, cmd_type, msg)
	}
}

func (this *CmdSystem) on_logout(user *User, cmd_type string, msg string) {
	user.client_socket.Leave(user.courseId)
	user_manager.UserCountSub(user.courseId)

	this.relay_user_count(user, cmd_type, msg)
	user_manager.RemoveUser(user.client_socket.Id())
	user.courseId = "0"
}

func (this *CmdSystem) on_disconnection(user *User, cmd_type string, msg string) {
	if strings.EqualFold(user.courseId, "0") {
		this.on_logout(user, cmd_type, msg)
	}
}

// ppt
func (this *CmdSystem) on_prev(user *User, cmd_type string, msg string) {

}

func (this *CmdSystem) on_next(user *User, cmd_type string, msg string) {

}

// chat
func (this *CmdSystem) on_chat_message(user *User, cmd_type string, msg string) {

}

// record
func (this *CmdSystem) on_start_record(user *User, cmd_type string, msg string) {

}

func (this *CmdSystem) on_stop_record(user *User, cmd_type string, msg string) {

}

/**
 * relay
 */
func (this *CmdSystem) relay_system(user *User, cmd_type string, msg string) {

}

func (this *CmdSystem) relay_ppt(user *User, cmd_type string, msg string) {
	user.client_socket.Emit(cmd_type, msg)
	user.client_socket.BroadcastTo(user.courseId, cmd_type, msg)
}

func (this *CmdSystem) relay_chat(user *User, cmd_type string, msg string) {
	user.client_socket.Emit(cmd_type, msg)
	user.client_socket.BroadcastTo(user.courseId, cmd_type, msg)
}

func (this *CmdSystem) relay_user_count(user *User, cmd_type string, msg string) {
	user_count := user_manager.UserCount(user.courseId)
	user_count_msg := new(UserCountMsg)
	user_count_msg.CourseId = user.courseId
	user_count_msg.UserCount = user_count
	_, err := json.Marshal(user_count_msg)
	if err != nil {
		panic(err.Error())
	}

	live_server.server.BroadcastTo(user.courseId, "user count", user_count)
	log.Println("user count", user_count)
}

/**
 * record
 */
func (this *CmdSystem) record_system(user *User, cmd_type string, msg string) {

}

func (this *CmdSystem) record_ppt(user *User, cmd_type string, msg string) {

}

func (this *CmdSystem) record_chat(user *User, cmd_type string, msg string) {

}

func (this *CmdSystem) record_media(user *User, cmd_type string, msg string) {

}
