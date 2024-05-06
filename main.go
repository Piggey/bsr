package main

import (
	"github.com/Piggey/bsr/client"
	"github.com/Piggey/bsr/server"
	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&Args)

	switch ctx.Command() {
	case "server":
		srv := server.NewServer(Args.Server.Addr)
		defer srv.Close()

		srv.Listen()

	case "client":
		client := client.NewClient(Args.Client.ServerAddr)
		defer client.Close()
	}
}
