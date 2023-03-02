package main

import (
	"project/iCredidentials/cmd/server"
)

func main() {

	server := server.InitServer()
	
	server.Start("localhost:8080")
}
