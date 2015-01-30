package jiaoxue

import (
	"log"
)

type Context struct {
	user_manager   *UserManager
	cmd_manager    *CmdManager
	relay_manager  *RelayManager
	record_manager *RecordManager
}

var jiao_context *Context

func GetContext() *Context {
	if jiao_context == nil {
		log.Println("NewContext")

		jiao_context = new(Context)
		jiao_context.user_manager = NewUserManager()
		jiao_context.cmd_manager = NewCmdManager()
		jiao_context.relay_manager = NewRelayManager()
		jiao_context.record_manager = NewRecordManager()
	}

	return jiao_context
}
