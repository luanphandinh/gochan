package gochan

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	channel chan string
	socket  *websocket.Conn
}

type Server struct {
	channel     chan string
	clients     map[string]*Client
	clientCount int
}

func (srv *Server) Channel() chan string {
	if srv.channel == nil {
		srv.channel = make(chan string, 4)
	}

	return srv.channel
}

func (srv *Server) Clients() map[string]*Client {
	if srv.clients == nil {
		srv.clients = make(map[string]*Client)
	}

	return srv.clients
}

func (srv *Server) AttachClient(name string, socket *websocket.Conn) chan string {
	serveClients := srv.Clients()
	if client := serveClients[name]; client != nil {
		return client.channel
	}

	channel := make(chan string, 1)
	client := &Client{channel, socket}
	serveClients[name] = client
	srv.clientCount += 1

	return client.channel
}

func (srv *Server) Broadcast(msg string) {
	for _, client := range srv.Clients() {
		client.channel <- msg
	}
}

func (srv *Server) Run() {
	for {
		select {
		case msg, _ := <-srv.Channel():
			srv.Broadcast(msg)
		}
	}
}
