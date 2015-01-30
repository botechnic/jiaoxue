package jiaoxue

import (
	"log"
)

type RelaySystem struct {
}

func NewRelaySystem() *RelaySystem {
	log.Println("NewRelaySystem")

	relay_system := new(RelaySystem)

	return relay_system
}

func (this *RelaySystem) Start(user *User) bool {
	log.Println("RelaySystem Start")

	return true
}

func (this *RelaySystem) Stop(user *User) bool {
	log.Println("RelaySystem Stop")

	return true
}
