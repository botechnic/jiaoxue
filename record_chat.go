package jiaoxue

import (
	"log"
)

type RecordChat struct {
	isRecording bool
}

func NewRecordChat() *RecordChat {
	log.Println("NewRecordChat")

	record_chat := new(RecordChat)
	record_chat.init()

	return record_chat
}

func (this *RecordChat) init() {
	this.isRecording = false
}

func (this *RecordChat) Start(user *User) bool {
	log.Println("RecordChat Start")

	this.isRecording = true
	return true
}

func (this *RecordChat) Stop(user *User) bool {
	log.Println("RecordChat Stop")

	this.isRecording = false
	return true
}
