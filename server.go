package gochan

import (
	"fmt"
)

type Serve struct {
	server      chan chan string
	clients     []chan string
	clientCount int
}

func (serve *Serve) Server() chan chan string {
	if serve.server == nil {
		serve.server = make(chan chan string, 4)
	}

	return serve.server
}

func (serve *Serve) Clients() []chan string {
	if serve.clients == nil {
		serve.clients = make([]chan string, 0)
	}

	return serve.clients
}

func (serve *Serve) attachClient(clients chan string) {
	serveClients := serve.Clients()

	serve.clients = append(serveClients, clients)
	serve.clientCount = len(serve.clients)
}

func (serve *Serve) Start() {
	server := serve.Server()

	for {
		select {
		case client := <- server:
			serve.attachClient(client)
			for _, client := range serve.Clients() {
				client <- fmt.Sprintf("%d clients connected.", serve.clientCount)
			}
		}
	}
}
