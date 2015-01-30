package jiaoxue

import (
	"log"
)

type JiaoCmd interface {
	Start(user *User) bool
	Stop(user *User) bool
}

type CmdManager struct {
	cmd_handler map[string]JiaoCmd
}

func NewCmdManager() *CmdManager {
	log.Println("NewCmdManager")

	cmd_manager := new(CmdManager)
	cmd_manager.cmd_handler = make(map[string]JiaoCmd)
	cmd_manager.init()

	return cmd_manager
}

func (this *CmdManager) Register(cmd_key string, cmd JiaoCmd) {
	log.Println("CmdManager Register", cmd_key)

	if this.cmd_handler != nil {
		this.cmd_handler[cmd_key] = cmd
	}
}

func (this *CmdManager) UnRegister(cmd_key string) {
	log.Println("CmdManager UnRegister", cmd_key)

	if this.cmd_handler != nil {
		delete(this.cmd_handler, cmd_key)
	}
}

func (this *CmdManager) init() {
	this.Register("system", NewCmdSystem())
	this.Register("ppt", NewCmdPPT())
	this.Register("media", NewCmdMedia())
	this.Register("chat", NewCmdChat())
	this.Register("record", NewCmdRecord())
}

func (this *CmdManager) uninit() {
	this.UnRegister("system")
	this.UnRegister("ppt")
	this.UnRegister("media")
	this.UnRegister("chat")
	this.UnRegister("record")
}

func (this *CmdManager) Start(user *User) bool {
	log.Println("CmdManager Start")

	for _, handler := range this.cmd_handler {
		handler.Start(user)
	}

	return true
}

func (this *CmdManager) Stop(user *User) bool {
	log.Println("CmdManager Stop")

	for _, handler := range this.cmd_handler {
		handler.Stop(user)
	}

	return true
}
