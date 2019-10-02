package gochan

import (
	"fmt"
)

type Server struct {
	channel     chan chan string
	clients     map[string]chan string
	clientCount int
}

func (srv *Server) Channel() chan chan string {
	if srv.channel == nil {
		srv.channel = make(chan chan string, 4)
	}

	return srv.channel
}

func (srv *Server) Clients() map[string]chan string {
	if srv.clients == nil {
		srv.clients = make(map[string]chan string)
	}

	return srv.clients
}

func (srv *Server) AttachClient(name string, client chan string) {
	serveClients := srv.Clients()
	if serveClients[name] != nil {
		return
	}

	serveClients[name] = client
	srv.clientCount += 1

	for _, client := range srv.Clients() {
		client <- fmt.Sprintf("%d clients connected.", srv.clientCount)
	}
}
