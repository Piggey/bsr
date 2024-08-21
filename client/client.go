package client

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/util"
)

type Client struct {
	srvConn net.Conn
	logger  *slog.Logger
}

func NewClient(serverAddr string) (*Client, error) {
	udpSrvAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("net.ResolveUDPAddr: %w", err)
	}

	srvConn, err := net.DialUDP("udp", nil, udpSrvAddr)
	if err != nil {
		return nil, fmt.Errorf("net.DialUDP: %w", err)
	}

	clientHandler := util.NewSlogHandler("client", srvConn.LocalAddr().String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(clientHandler)
	logger.Info("client connected")

	return &Client{
		srvConn: srvConn,
		logger:  logger,
	}, nil
}

func (c *Client) Close() error {
	c.logger.Info("connection closed")
	return c.srvConn.Close()
}

func (c *Client) StartNewGame(gameId uint8, mode packet.GameMode) error {
	ngp := packet.NewJoinGamePacket(gameId, mode)

	return binary.Write(c.srvConn, binary.BigEndian, ngp)
}

func (c *Client) Read(p packet.Packet) error {
	buf := make([]byte, 1024)
	n, err := c.srvConn.Read(buf)
	if err != nil {
		return fmt.Errorf("srvConn.Read: %w", err)
	}
	buf = buf[:n]

	err = p.FromBytes(buf)
	if err != nil {
		return fmt.Errorf("p.FromBytes: %w", err)
	}

	return nil
}

func (c *Client) Write(p packet.Packet) error {
	payload, err := p.ToBytes()
	if err != nil {
		return fmt.Errorf("p.ToBytes: %w", err)
	}

	_, err = c.srvConn.Write(payload)
	if err != nil {
		return fmt.Errorf("srvConn.Write: %w", err)
	}

	return nil
}
