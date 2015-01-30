package jiaoxue

import (
	"log"
)

type RelayMedia struct {
}

func NewRelayMedia() *RelayMedia {
	log.Println("NewRelayMedia")

	relay_media := new(RelayMedia)

	return relay_media
}

func (this *RelayMedia) Start(user *User) bool {
	log.Println("RelayMedia Start")

	return true
}

func (this *RelayMedia) Stop(user *User) bool {
	log.Println("RelayMedia Stop")

	return true
}

func (this *RelayMedia) Relay(user *User, cmd_type string, msg string) bool {
	return true
}
