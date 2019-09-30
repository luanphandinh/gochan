package gochan

import "fmt"

func Client(clientName string, clientChan chan string) {
	for {
		text, _ := <-clientChan
		fmt.Printf("%s: %s\n", clientName, text)
	}
}
