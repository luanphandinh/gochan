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

	server := new(gochan.Server)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade this HTTP connection to a WS connection:
		ws, _ := upgrader.Upgrade(w, r, nil)
		client := make(chan string, 1)
		server.AttachClient(r.Header.Get("Origin"), client)
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
