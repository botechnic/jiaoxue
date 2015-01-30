package jiaoxue

import (
	"log"
)

type RecordPPT struct {
	isRecording bool
}

func NewRecordPPT() *RecordPPT {
	log.Println("NewRecordPPT")

	record_ppt := new(RecordPPT)
	record_ppt.init()

	return record_ppt
}

func (this *RecordPPT) init() {
	this.isRecording = false
}

func (this *RecordPPT) Start(user *User) bool {
	log.Println("RecordPPT StartRecord")

	this.isRecording = true
	return true
}

func (this *RecordPPT) Stop(user *User) bool {
	log.Println("RecordPPT StopRecord")

	this.isRecording = false
	return true
}
