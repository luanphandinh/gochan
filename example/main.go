package main

import (
	"github.com/gorilla/websocket"
	"github.com/luanphandinh/gochan"
	"net/http"
)

func main() {
	var upgrader = websocket.Upgrader {
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r * http.Request) bool {
			return true // Disable CORS for testing
		},
	}

	server := new(gochan.Serve)
	go server.Start()

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade this HTTP connection to a WS connection:
		ws, _ := upgrader.Upgrade(w, r, nil)
		// And register a client for this connection with the uptimeServer:
		client := make(chan string, 1)
		server.Server() <- client
		// And now check for uptimes written to the client indefinitely.
		// Yes, we are lacking proper error and disconnect checking here, too:
		for {
			select {
			case text, _ := <-client:
				writer, _ := ws.NextWriter(websocket.TextMessage)
				writer.Write([]byte(text))
				writer.Close()
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
