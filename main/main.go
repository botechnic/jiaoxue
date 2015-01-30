package main

import (
	"github.com/botechnic/jiaoxue"
)

func main() {
	jiaoxue.GetContext()

	live_server := jiaoxue.NewLiveServer()
	live_server.Start()
}
