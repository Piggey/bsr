package main

import (
	"fmt"

	"github.com/Piggey/bsr/client"
	"github.com/Piggey/bsr/server"
)

func main() {
	fmt.Println("siema witam")
	srv := server.NewServer(":5000")
	defer srv.Close()

	srv.Listen()
	fmt.Printf("srv: %v\n", srv)

	client := client.NewClient("127.0.0.1:5000")
	defer client.Close()

	fmt.Printf("client: %v\n", client)
}
