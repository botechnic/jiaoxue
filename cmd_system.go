package jiaoxue

import (
	"github.com/googollee/go-socket.io"
	"log"
)

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
		this.on_login(user.client_socket, msg)
	})

	user.client_socket.On("logout", func(msg string) {
		this.on_logout(user.client_socket, msg)
	})

	user.client_socket.On("disconnection", func() {
		this.on_disconnection(user.client_socket)
	})

	return true
}

func (this *CmdSystem) Stop(user *User) bool {
	log.Println("CmdSystem Stop")

	return true
}

func (this *CmdSystem) on_login(client_socket socketio.Socket, msg string) {
	log.Println("CmdSystem on_login", client_socket.Id(), msg)

	user := GetContext().user_manager.HasUser(client_socket.Id())
	if user != nil {
		user.Login(msg)
	}
}

func (this *CmdSystem) on_logout(client_socket socketio.Socket, msg string) {
	log.Println("CmdSystem on_logout", client_socket.Id(), msg)

	user := GetContext().user_manager.HasUser(client_socket.Id())
	if user != nil {
		user.Logout(msg)
	}
}

func (this *CmdSystem) on_disconnection(client_socket socketio.Socket) {
	log.Println("CmdSystem on_disconnection", client_socket.Id())

	user := GetContext().user_manager.HasUser(client_socket.Id())
	if user != nil {
		user.Disconnect()
	}
}
