package jiaoxue

import (
	"fmt"
	"log"
)

type JiaoRecord interface {
	StartRecord(user *User) bool
	StopRecord(user *User) bool
}

type RecordManager struct {
	record_handler map[string]JiaoRecord
}

func NewRecordManager() *RecordManager {
	log.Println("NewRecordManager")

	record_manager := new(RecordManager)
	record_manager.init()

	return record_manager
}

func (this *RecordManager) init() {
	this.record_handler = make(map[string]JiaoRecord)

	recordPPTHandler := NewRecordPPT()
	recordMediaHandler := NewRecordMedia()
	recordChatHandler := NewRecordChat()

	this.record_handler["ppt"] = recordPPTHandler
	this.record_handler["media"] = recordMediaHandler
	this.record_handler["chat"] = recordChatHandler
}

func (this *RecordManager) IsRecording() bool {
	return false
}

func (this *RecordManager) StartRecord(user *User) bool {
	log.Println("RecordManager StartRecord")

	for key, handler := range this.record_handler {
		fmt.Println(key, "start record")
		handler.StartRecord(user)
	}

	return true
}

func (this *RecordManager) StopRecord(user *User) bool {
	log.Println("RecordManager StopRecord")

	for key, handler := range this.record_handler {
		fmt.Println(key, "stop record")
		handler.StopRecord(user)
	}

	return true
}
