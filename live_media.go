package jiaoxue

import (
	"log"
)

type LiveMedia struct {
}

func NewLiveMedia() *LiveMedia {
	log.Println("NewLiveMedia")

	live_media := new(LiveMedia)

	return live_media
}

func (this *LiveMedia) StartLive(user *User) bool {
	log.Println("LiveMedia StartLive")

	return true
}

func (this *LiveMedia) StopLive(user *User) bool {
	log.Println("LiveMedia StopLive")

	return true
}
