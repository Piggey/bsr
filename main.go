package main

import (
	"fmt"

	"github.com/Piggey/bsr/server"
)

func main() {
	fmt.Println("siema witam")
	srv := server.NewServer("127.0.0.1:5000")
	fmt.Printf("srv: %v\n", srv)
}
