package client

import (
	"fmt"
	"log/slog"
	"os"

	pb "github.com/Piggey/bsr/proto"
	"github.com/Piggey/bsr/util"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	bsrc   pb.BsrClient
	name   string
	logger *slog.Logger
}

func NewClient(name, srvAddr string) (*Client, error) {
	conn, err := grpc.NewClient(srvAddr)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	bsrc := pb.NewBsrClient(conn)

	clientHandler := util.NewSlogHandler("client", "", os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(clientHandler)

	return &Client{
		conn:   conn,
		bsrc:   bsrc,
		name:   name,
		logger: logger,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) StartNewGame() error {
	panic("unimplemented")
}
