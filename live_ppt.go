package jiaoxue

import (
	"log"
)

type LivePPT struct {
}

func NewLivePPT() *LivePPT {
	log.Println("NewLivePPT")

	live_ppt := new(LivePPT)

	return live_ppt
}

func (this *LivePPT) StartLive(user *User) bool {
	log.Println("LivePPT StartLIve")

	return true
}

func (this *LivePPT) StopLive(user *User) bool {
	log.Println("LivePPT StopLIve")

	return true
}
