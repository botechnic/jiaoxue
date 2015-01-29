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

func (this *User) StartRecord() bool {

	return true
}

func (this *User) StopRecord() bool {

	return true
}

func (this *User) StartLive() {
	GetContext().live_manager.StartLive(this)
}

func (this *User) StopLive() {
	GetContext().live_manager.StopLive(this)
}

func (this *User) Login(login_msg string) {
	user_id := "4"
	user_name := "s1"
	user_role := "student"
	course_id := "chat"

	this.userId = user_id
	this.userName = user_name
	this.courseId = course_id
	this.role = user_role

	this.client_socket.Join(this.courseId)
	this.StartLive()
}

func (this *User) Logout(logout_msg string) {
	this.client_socket.Leave(this.courseId)
}

func (this *User) Disconnect() {
	this.StopLive()
	this.StartRecord()
	GetContext().users.RemoveUser(this.client_socket.Id())
}
