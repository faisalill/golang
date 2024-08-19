package main

import "log"

func main() {
	server := NewApiServer(":8000")

	server.Run()

	log.Println("Server Starting at: ", server.serverAddress)
}
