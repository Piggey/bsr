package main

import (
	"context"
	"log"
	"time"

	"github.com/Piggey/bsr/client"
	bsr "github.com/Piggey/bsr/proto"
	"github.com/Piggey/bsr/server"
	"github.com/alecthomas/kong"
)

var args struct {
	Client struct {
		ServerAddr string  `help:"server address" required:""`
		Addr       *string `help:"optional client address, defaults to first available port"`
		Name       string  `help:"client name" default:"player1"`
		Pvp        struct {
			GameUuid uint8 `help:"game id to connect players to the same game" required:""`
		} `cmd:""`
		Pve struct{} `cmd:""`
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

	case "client pve":
		c, err := client.NewClient(args.Client.Name, args.Client.ServerAddr)
		if err != nil {
			log.Fatalf("client.NewClient: %v", err)
		}
		defer c.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = c.JoinGame(ctx, bsr.GameMode_PVE)
		if err != nil {
			log.Fatalf("client.StartNewGame: %v", err)
		}

	case "client pvp":
		panic("unimplemented")
	}
}
