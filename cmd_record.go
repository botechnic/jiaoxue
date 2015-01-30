package jiaoxue

import (
	"log"
)

type CmdRecord struct {
}

func NewCmdRecord() *CmdRecord {
	log.Println("NewCmdRecord")

	cmd_record := new(CmdRecord)
	return cmd_record
}

func (this *CmdRecord) Start(user *User) bool {
	log.Println("CmdRecord Start")

	user.client_socket.On("start_record", func(msg string) {
		this.on_start_record(user, msg)
	})

	user.client_socket.On("stop_record", func(msg string) {
		this.on_stop_record(user, msg)
	})

	return true
}

func (this *CmdRecord) Stop(user *User) bool {
	log.Println("CmdRecord Stop")

	return false
}

func (this *CmdRecord) on_start_record(user *User, start_record_msg string) {
	log.Println("CmdRecord on_start_record", start_record_msg)

	user.StartRecord()
}

func (this *CmdRecord) on_stop_record(user *User, stop_record_msg string) {
	log.Println("CmdRecord on_stop_record", stop_record_msg)

	user.StopRecord()
}
