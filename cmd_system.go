package jiaoxue

import (
	"log"
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
	log.Println("NewCmdSystem")

	cmd_system := new(CmdSystem)

	return cmd_system
}

func (this *CmdSystem) Start(user *User) bool {
	log.Println("CmdSystem Start")

	user.client_socket.On("login", func(msg string) {
		this.on_login(user, "login", msg)
	})

	user.client_socket.On("logout", func(msg string) {
		this.on_logout(user, "logout", msg)
	})

	user.client_socket.On("disconnection", func() {
		this.on_disconnection(user, "disconnection")
	})

	return true
}

func (this *CmdSystem) Stop(user *User) bool {
	log.Println("CmdSystem Stop")

	return true
}

func (this *CmdSystem) on_login(user *User, cmd_type string, msg string) {
	log.Println("CmdSystem on_login", user.client_socket.Id(), msg)

	if user != nil {
		//user.Login(&login_msg)

		relay_system := GetContext().relay_manager.Relay("system")
		relay_system.Relay(user, cmd_type, msg)
	}

}

func (this *CmdSystem) on_logout(user *User, cmd_type string, msg string) {
	log.Println("CmdSystem on_logout", user.client_socket.Id())

	if user != nil {
		//user.Logout()

		relay_system := GetContext().relay_manager.Relay("system")
		relay_system.Relay(user, cmd_type, msg)
	}
}

func (this *CmdSystem) on_disconnection(user *User, cmd_type string) {
	log.Println("CmdSystem on_disconnection", user.client_socket.Id())

	if user != nil {
		relay_system := GetContext().relay_manager.Relay("system")
		relay_system.Relay(user, cmd_type, "")

		//user.Disconnect()
	}
}
