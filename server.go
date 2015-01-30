package jiaoxue

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

var user_manager *UserManager = NewUserManager()
var live_server *LiveServer = NewLiveServer()

type LiveServer struct {
	server *socketio.Server
}

func GetLiveServer() *LiveServer {
	return live_server
}

func NewLiveServer() *LiveServer {
	var err interface{}
	live_server := new(LiveServer)
	live_server.server, err = socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	return live_server
}

func (this *LiveServer) Start() bool {
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

func (this *LiveServer) on_error(client_socket socketio.Socket, err error) {
	log.Println("on_error", client_socket.Id(), err)
}

func (this *LiveServer) on_connected(client_socket socketio.Socket) {
	log.Println("on_connected", client_socket.Id())

	user := NewUser(client_socket)
	user_manager.AddUser(client_socket.Id(), user)
}
