package jiaoxue

import (
	"log"
)

type LiveRecord struct {
}

func NewLiveRecord() *LiveRecord {
	live_record := new(LiveRecord)
	return live_record
}

func (this *LiveRecord) StartLive(user *User) bool {
	log.Println("LiveRecord StartLive")

	user.client_socket.On("start_record", func(msg string) {
		this.on_start_record(user, msg)
	})

	user.client_socket.On("stop_record", func(msg string) {
		this.on_stop_record(user, msg)
	})

	return true
}

func (this *LiveRecord) StopLive(user *User) bool {
	log.Println("LiveRecord StopLive")

	return false
}

func (this *LiveRecord) on_start_record(user *User, start_record_msg string) {
	GetContext().record_manager.StartRecord(user)
}

func (this *LiveRecord) on_stop_record(user *User, stop_record_msg string) {
	GetContext().record_manager.StopRecord(user)
}
