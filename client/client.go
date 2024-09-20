package client

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	pb "github.com/Piggey/bsr/proto"
	"github.com/Piggey/bsr/util"
	"github.com/google/uuid"
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

func (c *Client) JoinGame(ctx context.Context, mode pb.GameMode) error {
	c.logger.Info("starting new game", slog.String("mode", mode.String()))

	pveMaxPlayerCount := uint32(2)

	joinGameRes, err := c.bsrc.JoinGame(ctx, &pb.JoinGameRequest{
		Version:        pb.BsrV1,
		GameUuid:       uuid.NewString(),
		GameMode:       pb.GameMode_PVE,
		PlayerName:     c.name,
		MaxPlayerCount: &pveMaxPlayerCount,
	})
	if err != nil {
		return fmt.Errorf("c.bsrc.JoinGame: %w", err)
	}

	c.logger.Info("joined game", slog.String("playerUuid", joinGameRes.PlayerUuid))
	return nil
}
