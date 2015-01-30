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
}

func (this *LiveServer) Start() bool {
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

func (this *LiveServer) Stop() bool {
	log.Println("StopLiveServer")
	return true
}

func (this *LiveServer) on_error(client_socket socketio.Socket, err error) {
	log.Println("on_error", client_socket.Id(), err)
}

func (this *LiveServer) on_connected(client_socket socketio.Socket) {
	log.Println("on_connected", client_socket.Id())

	on_connected1(client_socket)
}

func on_connected1(client_socket socketio.Socket) {
	log.Println("on_connected1", client_socket.Id())

	user := GetContext().user_manager.HasUser(client_socket.Id())
	if user == nil {
		user = NewUser(client_socket)
		GetContext().user_manager.AddUser(client_socket.Id(), user)

		relay_system := GetContext().relay_manager.Relay("system")
		relay_system.Relay(user, "connection", "")
	}
}
