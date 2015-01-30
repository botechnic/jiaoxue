package main

import (
	"github.com/botechnic/jiaoxue"
)

func main() {
	live_server := jiaoxue.GetLiveServer()
	live_server.Start()
}
