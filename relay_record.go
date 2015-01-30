package jiaoxue

import (
	"log"
)

type RelayRecord struct {
}

func NewRelayRecord() *RelayRecord {
	log.Println("NewRelayRecord")

	relay_record := new(RelayRecord)

	return relay_record
}

func (this *RelayRecord) Start(user *User) bool {
	log.Println("RelayRecord Start")

	return true
}

func (this *RelayRecord) Stop(user *User) bool {
	log.Println("RelayRecord Stop")

	return true
}

func (this *RelayRecord) Relay(user *User, cmd_type string, msg string) bool {
	return true
}
