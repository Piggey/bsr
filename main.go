package main

import (
	"context"
	"log"
	"time"

	"github.com/Piggey/bsr/client"
	"github.com/Piggey/bsr/server"
	"github.com/alecthomas/kong"
)

var args struct {
	Client struct {
		ServerAddr     string  `help:"server address" required:""`
		Addr           *string `help:"optional client address, defaults to first available port"`
		Name           string  `help:"client name" default:"player1"`
		MaxPlayerCount uint32  `help:"number of players in a lobby" default:"2"`
	} `cmd:"" help:"run as client"`
	Server struct {
		Addr string `help:"server address" default:":5000"`
	} `cmd:"" help:"run as server"`
}

func main() {
	ctx := kong.Parse(&args)

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
		c, err := client.NewClient(args.Client.Name, args.Client.ServerAddr)
		if err != nil {
			log.Fatalf("client.NewClient: %v", err)
		}
		defer c.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = c.JoinGame(ctx, args.Client.MaxPlayerCount)
		if err != nil {
			log.Fatalf("client.StartNewGame: %v", err)
		}
	}
}
