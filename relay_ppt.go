package jiaoxue

import (
	"log"
)

type RelayPPT struct {
}

func NewRelayPPT() *RelayPPT {
	log.Println("NewRelayPPT")

	relay_ppt := new(RelayPPT)

	return relay_ppt
}

func (this *RelayPPT) Start(user *User) bool {
	log.Println("RelayPPT Start")

	return true
}

func (this *RelayPPT) Stop(user *User) bool {
	log.Println("RelayPPT Stop")

	return true
}

func (this *RelayPPT) Relay(user *User, cmd_type string, msg string) bool {
	log.Println("RelayPPT", cmd_type)

	user.client_socket.Emit(cmd_type, msg)
	user.client_socket.BroadcastTo(user.courseId, cmd_type, msg)

	return true
}
