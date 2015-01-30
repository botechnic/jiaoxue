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
	log.Println("CmdChat StartLive")

	user.client_socket.On("chat message", func(msg string) {
		this.on_chat_message(user, msg)
	})

	return true
}

func (this *CmdChat) Stop(user *User) bool {
	log.Println("CmdChat StopLive")

	return true
}

func (this *CmdChat) on_chat_message(user *User, msg string) {
	log.Println("on_chat_message", user.client_socket.Id(), msg)

	user.client_socket.Emit("chat message", msg)
	user.client_socket.BroadcastTo("chat", "chat message", msg)
}
