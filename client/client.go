package client

import (
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/packet/binary"
	"github.com/Piggey/bsr/util"
)

type Client struct {
	addr    net.Addr
	udpConn *net.UDPConn
	logger  *slog.Logger
}

func NewClient(serverAddr string) *Client {
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		panic(err)
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}

	clientHandler := util.NewCustomHandler("client", udpConn.LocalAddr().String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(clientHandler)

	logger.Info("client connected")

	return &Client{
		addr:    udpConn.LocalAddr(),
		udpConn: udpConn,
		logger:  logger,
	}
}

func (c *Client) Close() error {
	c.logger.Info("connection closed")
	return c.udpConn.Close()
}

func (c *Client) StartNewGame(mode packet.GameMode) error {
	ngp := packet.NewCreateGamePacket(mode)

	return binary.Write(c.udpConn, binary.BigEndian, ngp)
}
