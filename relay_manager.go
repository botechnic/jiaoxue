package jiaoxue

import (
	"log"
)

type JiaoRelay interface {
	Start(user *User) bool
	Stop(user *User) bool
}

type RelayManager struct {
	relay_handler map[string]JiaoRelay
}

func NewRelayManager() *RelayManager {
	log.Println("NewRelayManager")

	relay_manager := new(RelayManager)
	relay_manager.relay_handler = make(map[string]JiaoRelay)
	relay_manager.init()

	return relay_manager
}

func (this *RelayManager) Register(cmd_key string, cmd JiaoCmd) {
	log.Println("RelayManager Register", cmd_key)

	if this.relay_handler != nil {
		this.relay_handler[cmd_key] = cmd
	}
}

func (this *RelayManager) UnRegister(cmd_key string) {
	log.Println("RelayManager UnRegister", cmd_key)

	if this.relay_handler != nil {
		delete(this.relay_handler, cmd_key)
	}
}

func (this *RelayManager) init() {
	this.Register("system", NewCmdSystem())
	this.Register("ppt", NewCmdPPT())
	this.Register("media", NewCmdMedia())
	this.Register("chat", NewCmdChat())
	this.Register("record", NewCmdRecord())
}

func (this *RelayManager) uninit() {
	this.UnRegister("system")
	this.UnRegister("ppt")
	this.UnRegister("media")
	this.UnRegister("chat")
	this.UnRegister("record")
}

func (this *RelayManager) Start(user *User) bool {
	log.Println("RelayManager Start")

	for _, handler := range this.relay_handler {
		handler.Start(user)
	}

	return true
}

func (this *RelayManager) Stop(user *User) bool {
	log.Println("RelayManager Stop")

	for _, handler := range this.relay_handler {
		handler.Stop(user)
	}

	return true
}
