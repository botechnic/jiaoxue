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

func (this *User) StartRecord() {
	log.Println("User StartRecord")

	GetContext().record_manager.Start(this)
}

func (this *User) StopRecord() {
	log.Println("User StopRecord")

	GetContext().record_manager.Stop(this)
}

func (this *User) StartLive() {
	log.Println("User StartLive")

	GetContext().cmd_manager.Start(this)
}

func (this *User) StopLive() {
	log.Println("User StopLive")

	GetContext().cmd_manager.Stop(this)
}

func (this *User) Login(login_msg string) {
	log.Println("User Login", login_msg)

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
	log.Println("User Logout", logout_msg)

	this.client_socket.Leave(this.courseId)
}

func (this *User) Disconnect() {
	log.Println("User Disconnect")

	this.StopLive()
	this.StartRecord()
	GetContext().user_manager.RemoveUser(this.client_socket.Id())
}
