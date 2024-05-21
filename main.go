package main

import (
	"github.com/Piggey/bsr/client"
	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/server"
	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&Args)

	// open server
	// connect client
	// client -> server: start new game
	// server starts up game
	// server -> client: game state for new game
	// client -> server: move
	// server validates move
	// server -> client: game state

	switch ctx.Command() {
	case "server":
		srv := server.NewServer(Args.Server.Addr)
		defer srv.Close()

		srv.Listen()

	case "client":
		client := client.NewClient(Args.Client.ServerAddr)
		defer client.Close()

		err := client.StartNewGame(packet.GameModePvE) // for now
		if err != nil {
			panic(err)
		}
	}
}
