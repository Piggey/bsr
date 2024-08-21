package main

import (
	"log"

	"github.com/Piggey/bsr/client"
	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/server"
	"github.com/alecthomas/kong"
)

var args struct {
	Client struct {
		ServerAddr string  `help:"server address" required:""`
		Addr       *string `help:"optional client address, defaults to first available port"`
		Pvp        struct {
			GameId uint8 `help:"game id to connect players to the same game" required:""`
		} `cmd:""`
		Pve struct{} `cmd:""`
	} `cmd:"" help:"run as client"`
	Server struct {
		Addr string `help:"server address" default:":5000"`
	} `cmd:"" help:"run as server"`
}

func main() {
	ctx := kong.Parse(&args)

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
		srv, err := server.NewServer(args.Server.Addr)
		if err != nil {
			log.Fatalf("server.NewServer: %v", err)
		}
		defer srv.Close()

		err = srv.Listen()
		if err != nil {
			log.Fatalf("srv.Listen: %v", err)
		}

	case "client":
		client, err := client.NewClient(args.Client.ServerAddr)
		if err != nil {
			log.Fatalf("client.NewClient: %v", err)
		}
		defer client.Close()

		err = client.StartNewGame(packet.GameModePvE) // for now
		if err != nil {
			log.Fatalf("client.StartNewGame: %v", err)
		}

	case "client pvp":
		panic("unimplemented")
	}
}
