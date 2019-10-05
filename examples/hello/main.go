package main

import "fmt"

func greeting(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	// Make a channel of string
	fmt.Println("Start")
	c := make(chan string)

	// Start a goroutine
	go greeting(c)

	// Sending data to channel
	// Here the main go routine is blocked until some other go routine read it
	c <- "Join"
	// go greeting read data, so main go routine is unblocked
	fmt.Println("Done")
}
