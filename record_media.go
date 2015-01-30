package jiaoxue

import (
	"log"
	"os/exec"
)

type RecordMedia struct {
	isRecording   bool
	RecordHandler *exec.Cmd
}

func NewRecordMedia() *RecordMedia {
	log.Println("NewRecordMedia")

	record_media := new(RecordMedia)
	record_media.init()

	return record_media
}

func (this *RecordMedia) init() {
	this.isRecording = false
	this.RecordHandler = nil
}

func (this *RecordMedia) Start(user *User) bool {
	log.Println("RecordMedia StartRecord")

	this.isRecording = true

	return true
}

func (this *RecordMedia) Stop(user *User) bool {
	log.Println("RecordMedia StopRecord")

	this.isRecording = false

	return true
}
