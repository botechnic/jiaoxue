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
