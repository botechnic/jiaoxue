package jiaoxue

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

type LiveServer struct {
	server *socketio.Server
}

func NewLiveServer() *LiveServer {
	log.Println("NewLiveServer")

	live_server := new(LiveServer)
	live_server.init()

	return live_server
}

func (this *LiveServer) init() {
	server, err := socketio.NewServer(nil)
	this.server = server

	if err != nil {
		log.Fatal(err)
	}

	users := NewUsers()
	GetContext().users = users
}

func (this *LiveServer) StartLiveServer() bool {
	log.Println("StartLiveServer")

	ipport := ":5000"
	static_path := "/home/pony/goworkspace/jiaoxue/bin/asset"

	if this.server != nil {
		this.server.On("error", func(client_socket socketio.Socket, err error) {
			this.on_error(client_socket, err)
		})

		this.server.On("connection", func(client_socket socketio.Socket) {
			this.on_connected(client_socket)
		})

		http.Handle("/socket.io/", this.server)
		http.Handle("/", http.FileServer(http.Dir(static_path)))

		log.Printf("Serving at %v ...\n", ipport)
		log.Fatal(http.ListenAndServe(ipport, nil))
	}
	return true
}

func (this *LiveServer) StopLiveServer() bool {
	log.Println("StopLiveServer")
	return true
}

func (this *LiveServer) on_error(client_socket socketio.Socket, err error) {
	log.Println("on_error:", client_socket.Id(), err)
}

func (this *LiveServer) on_connected(client_socket socketio.Socket) {
	log.Println("on_connected", client_socket.Id())

	this.on_connected1(client_socket)

	this.bind_system(client_socket)
}

func (this *LiveServer) on_connected1(client_socket socketio.Socket) {
	user := GetContext().users.HasUser(client_socket.Id())
	if user == nil {
		user = NewUser(client_socket)
		GetContext().users.AddUser(client_socket.Id(), user)
	}
}

func (this *LiveServer) bind_system(client_socket socketio.Socket) {
	client_socket.On("login", func(msg string) {
		this.on_login(client_socket, msg)
	})

	client_socket.On("logout", func(msg string) {
		this.on_logout(client_socket, msg)
	})

	client_socket.On("disconnection", func() {
		this.on_disconnection(client_socket)
	})
}

func (this *LiveServer) on_login(client_socket socketio.Socket, msg string) {
	log.Println("on_login", client_socket.Id(), msg)

	user := GetContext().users.HasUser(client_socket.Id())
	if user != nil {
		user.Login(msg)
	}
}

func (this *LiveServer) on_logout(client_socket socketio.Socket, msg string) {
	log.Println("on_logout", client_socket.Id(), msg)

	user := GetContext().users.HasUser(client_socket.Id())
	if user != nil {
		user.Logout(msg)
	}
}

func (this *LiveServer) on_disconnection(client_socket socketio.Socket) {
	log.Println("on_disconnection", client_socket.Id())

	user := GetContext().users.HasUser(client_socket.Id())
	if user != nil {
		user.Disconnect()
	}
}
