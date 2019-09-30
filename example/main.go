package main

import (
	"github.com/luanphandinh/gochan"
	"time"
)

func main() {
	server := new(gochan.Serve)
	go server.Start()

	client1Chan := make(chan string, 4)
	client2Chan := make(chan string, 4)
	server.Server() <- client1Chan
	server.Server() <- client2Chan

	go gochan.Client("Client 1", client1Chan)
	go gochan.Client("Client 2", client2Chan)

	time.Sleep(time.Second)
}
