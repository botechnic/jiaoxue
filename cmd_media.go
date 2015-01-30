package jiaoxue

import (
	"log"
)

type CmdMedia struct {
}

func NewCmdMedia() *CmdMedia {
	log.Println("NewCmdMedia")

	cmd_media := new(CmdMedia)

	return cmd_media
}

func (this *CmdMedia) Start(user *User) bool {
	log.Println("CmdMedia Start")

	return true
}

func (this *CmdMedia) Stop(user *User) bool {
	log.Println("CmdMedia Stop")

	return true
}
