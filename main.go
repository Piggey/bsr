package main

import (
	"fmt"

	"github.com/Piggey/bsr/client"
	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/server"
	"github.com/alecthomas/kong"
)

var args struct {
	Client struct {
		ServerAddr string  `help:"server address"`
		Addr       *string `help:"optional client address, defaults to first available port"`
		Mode       string  `help:"against other player (pvp) or against ai (pve)" default:"pve"`
		// Pvp        struct {
		// 	OpponentAddr string `help:"opponents address"`
		// } `default:"1" help:"play against other client"`
	} `cmd:"" help:"run as client"`
	Server struct {
		Addr string `help:"server address" default:":5000"`
	} `cmd:"" help:"run as server"`
}

func main() {
	ctx := kong.Parse(&args)

	fmt.Printf("args: %v\n", args)

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
		srv := server.NewServer(args.Server.Addr)
		defer srv.Close()

		err := srv.Listen()
		if err != nil {
			panic(err)
		}

	case "client":
		client := client.NewClient(args.Client.ServerAddr)
		defer client.Close()

		err := client.StartNewGame(packet.GameModePvE) // for now
		if err != nil {
			panic(err)
		}
	}
}
