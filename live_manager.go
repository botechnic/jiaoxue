package jiaoxue

import (
	"log"
)

type JiaoLive interface {
	StartLive(user *User) bool
	StopLive(user *User) bool
}

type LiveManager struct {
	live_handler map[string]JiaoLive
}

func NewLiveManager() *LiveManager {
	log.Println("NewLiveManager")

	live_manager := new(LiveManager)
	live_manager.init()

	return live_manager
}

func (this *LiveManager) init() {
	live_ppt := NewLivePPT()
	live_media := NewLiveMedia()
	live_chat := NewLiveChat()

	this.live_handler["ppt"] = live_ppt
	this.live_handler["media"] = live_media
	this.live_handler["chat"] = live_chat
}

func (this *LiveManager) StartLive(user *User) bool {
	log.Println("LiveManager StartLive")

	for _, handler := range this.live_handler {
		handler.StartLive(user)
	}

	return true

}

func (this *LiveManager) StopLive(user *User) bool {
	log.Println("LiveManager StopLive")

	for _, handler := range this.live_handler {
		handler.StopLive(user)
	}

	return true
}
