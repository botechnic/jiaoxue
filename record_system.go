package jiaoxue

import (
	"log"
)

type RecordSystem struct {
	isRecording bool
}

func NewRecordSystem() *RecordSystem {
	log.Println("NewRecordSystem")

	record_media := new(RecordSystem)
	record_media.init()

	return record_media
}

func (this *RecordSystem) init() {
	this.isRecording = false
}

func (this *RecordSystem) Start(user *User) bool {
	log.Println("RecordMedia Start")

	this.isRecording = true

	return true
}

func (this *RecordSystem) Stop(user *User) bool {
	log.Println("RecordMedia Stop")

	this.isRecording = false

	return true
}
