package jiaoxue

import (
	"log"
)

type LiveChat struct {
}

func NewLiveChat() *LiveChat {
	log.Println("newLiveChat")

	live_chat := new(LiveChat)

	return live_chat
}

func (this *LiveChat) StartLive(user *User) bool {
	log.Println("LiveChat StartLive")

	user.client_socket.On("chat message", func(msg string) {
		this.on_chat_message(user, msg)
	})

	return true
}

func (this *LiveChat) StopLive(user *User) bool {
	log.Println("LiveChat StopLive")

	return true
}

func (this *LiveChat) on_chat_message(user *User, msg string) {
	log.Println("on_chat_message", user.client_socket.Id(), msg)

	user.client_socket.Emit("chat message", msg)
	user.client_socket.BroadcastTo("chat", "chat message", msg)
}
