package client

import (
	"context"
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

func (c *Client) NewGame(mode pb.GameMode) error {
	ctx := context.Background()
	c.logger.Info("starting new game", slog.String("mode", mode.String()))

	cg, err := c.bsrc.CreateGame(ctx, &pb.CreateGameRequest{
		Version:    pb.BsrProtoV1,
		PlayerName: c.name,
		Mode:       mode,
	})
	if err != nil {
		return fmt.Errorf("bsrc.CreateGame: %w", err)
	}

	_ = cg
	return nil
}
