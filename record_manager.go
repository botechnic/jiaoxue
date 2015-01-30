package jiaoxue

import (
	"log"
)

type JiaoRecorder interface {
	Start(user *User) bool
	Stop(user *User) bool
}

type RecordManager struct {
	record_handler map[string]JiaoRecorder
}

func NewRecordManager() *RecordManager {
	log.Println("NewRecordManager")

	record_manager := new(RecordManager)
	record_manager.init()

	return record_manager
}

func (this *RecordManager) Register(cmd_key string, cmd JiaoCmd) {
	log.Println("RecordManager Register", cmd_key)

	if this.record_handler != nil {
		this.record_handler[cmd_key] = cmd
	}
}

func (this *RecordManager) UnRegister(cmd_key string) {
	log.Println("RecordManager UnRegister", cmd_key)

	if this.record_handler != nil {
		delete(this.record_handler, cmd_key)
	}
}

func (this *RecordManager) init() {
	this.Register("system", NewRecordSystem())
	this.Register("ppt", NewRecordPPT())
	this.Register("media", NewRecordMedia())
	this.Register("chat", NewRecordChat())
}

func (this *RecordManager) uninit() {
	this.UnRegister("system")
	this.UnRegister("ppt")
	this.UnRegister("media")
	this.UnRegister("chat")
}

func (this *RecordManager) IsRecording() bool {
	return false
}

func (this *RecordManager) Start(user *User) bool {
	log.Println("RecordManager Start")

	for _, handler := range this.record_handler {
		handler.Start(user)
	}

	return true
}

func (this *RecordManager) Stop(user *User) bool {
	log.Println("RecordManager Stop")

	for _, handler := range this.record_handler {
		handler.Stop(user)
	}

	return true
}
