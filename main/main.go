package main

import (
	"github.com/botechnic/jiaoxue"
)

func main() {
	live_server := jiaoxue.NewLiveServer()
	jiaoxue.GetContext().SetLiveServer(live_server)

	live_server.Start()
}
