package jiaoxue

import (
	"log"
)

type CmdChat struct {
}

func NewCmdChat() *CmdChat {
	log.Println("newCmdChat")

	live_chat := new(CmdChat)

	return live_chat
}

func (this *CmdChat) Start(user *User) bool {
	log.Println("CmdChat Start")

	user.client_socket.On("chat message", func(msg string) {
		this.on_chat_message(user, "chat message", msg)
	})

	return true
}

func (this *CmdChat) Stop(user *User) bool {
	log.Println("CmdChat Stop")

	return true
}

func (this *CmdChat) on_chat_message(user *User, cmd_type string, msg string) {
	log.Println("on_chat_message", user.client_socket.Id(), msg)

	relay_chat := GetContext().relay_manager.Relay("chat")
	relay_chat.Relay(user, cmd_type, msg)
}
