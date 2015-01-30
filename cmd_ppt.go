package jiaoxue

import (
	"log"
)

type CmdPPT struct {
}

func NewCmdPPT() *CmdPPT {
	log.Println("NewCmdPPT")

	cmd_ppt := new(CmdPPT)

	return cmd_ppt
}

func (this *CmdPPT) Start(user *User) bool {
	log.Println("CmdPPT Start")

	return true
}

func (this *CmdPPT) Stop(user *User) bool {
	log.Println("CmdPPT Stop")

	return true
}
