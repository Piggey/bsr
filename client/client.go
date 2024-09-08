package client

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/util"
)

type Client struct {
	conn   net.Conn
	logger *slog.Logger
}

func NewClient(network, srvAddr string) (*Client, error) {
	conn, err := net.Dial(network, srvAddr)
	if err != nil {
		return nil, fmt.Errorf("net.Dial: %w", err)
	}

	clientHandler := util.NewSlogHandler("client", conn.LocalAddr().String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(clientHandler)
	logger.Info("client created")

	return &Client{
		conn:   conn,
		logger: logger,
	}, nil
}

func (c *Client) Close() error {
	c.logger.Info("client closing")
	return c.conn.Close()
}
