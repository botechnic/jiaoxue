package jiaoxue

import (
	"github.com/googollee/go-socket.io"
	"log"
)

type User struct {
	userId   string
	userName string
	courseId string
	role     string

	client_socket socketio.Socket
	isRecording   bool
}

func NewUser(client_socket socketio.Socket) *User {
	log.Println("NewUser")

	user := new(User)
	user.client_socket = client_socket

	return user
}
