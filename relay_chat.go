package jiaoxue

import (
	"log"
)

type RelayChat struct {
}

func NewRelayChat() *RelayChat {
	log.Println("NewRelayChat")

	relay_chat := new(RelayChat)

	return relay_chat
}

func (this *RelayChat) Start(user *User) bool {
	log.Println("RelayChat Start")
	return true
}

func (this *RelayChat) Stop(user *User) bool {
	log.Println("RelayChat Stop")

	return true
}

func (this *RelayChat) Relay(user *User, cmd_type string, msg string) bool {
	log.Println("RelayChat", cmd_type, user.client_socket.Id(), msg)

	user.client_socket.Emit(cmd_type, msg)
	user.client_socket.BroadcastTo(user.courseId, cmd_type, msg)

	return true
}
