package jiaoxue

import (
	"log"
)

type CmdPPT struct {
}

func NewCmdPPT() *CmdPPT {
	log.Println("NewCmdPPT")

	cmd_ppt := new(CmdPPT)

	return cmd_ppt
}

func (this *CmdPPT) Start(user *User) bool {
	log.Println("CmdPPT Start")

	user.client_socket.On("prev", func(msg string) {
		this.on_prev(user, "prev", msg)
	})

	user.client_socket.On("next", func(msg string) {
		this.on_next(user, "next", msg)
	})

	return true
}

func (this *CmdPPT) Stop(user *User) bool {
	log.Println("CmdPPT Stop")

	return true
}

func (this *CmdPPT) on_prev(user *User, cmd_type string, msg string) {
	log.Println("CmdPPT on_prev", user.client_socket.Id())

}

func (this *CmdPPT) on_next(user *User, cmd_type string, msg string) {
	log.Println("CmdPPT on_next", user.client_socket.Id())

	relay_ppt := GetContext().relay_manager.Relay("ppt")
	relay_ppt.Relay(user, cmd_type, msg)
}
