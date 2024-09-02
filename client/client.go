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
	conn net.Conn
	logger *slog.Logger
}

func NewClient(network, srvAddr string) (*Client, error) {
	conn, err := net.Dial(network, srvAddr)
	if err != nil {
		return nil, fmt.Errorf("net.Dial: %w", err)
	}

	clientHandler := util.NewSlogHandler("client", conn.LocalAddr().String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level: slog.LevelDebug,
	})
	logger := slog.New(clientHandler)
	logger.Info("client created")

	return &Client{
		conn: conn,
		logger: logger,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Handshake() error {
	pphReq := packet.PacketProtocolHandshakeRequest{
		ProtocolVersion: packet.ProtoV1,
	}

	err := c.writePacket(&pphReq)
	if err != nil {
		return fmt.Errorf("c.writePacket: %w", err)
	}

	// await response
	pphRes := packet.PacketProtocolHandshakeResponse{}
	err = c.readPacket(&pphRes)
	if err != nil {
		return fmt.Errorf("c.readPacket: %w", err)
	}
	if pphRes.Status != packet.HandshakeStatusOK {
		return fmt.Errorf("pphRes.Status != packet.HandshakeStatusOK")
	}

	return nil
}

func (c *Client) writePacket(p packet.Packet) error {
	data, err := p.ToBytes()
	if err != nil {
		return fmt.Errorf("p.ToBytes: %w", err)
	}

	n, err := c.conn.Write(data)
	if err != nil {
		return fmt.Errorf("conn.Write: %w", err)
	}
	if len(data) != n {
		return fmt.Errorf("len(data) != n")
	}

	return nil
}

func (c *Client) readPacket(p packet.Packet) error {
	data := make([]byte, p.Size())
	n, err := c.conn.Read(data)
	if err != nil {
		return fmt.Errorf("conn.Read: %w", err)
	}
	if len(data) != n {
		return fmt.Errorf("len(data) != n")
	}

	err = p.FromBytes(data)
	if err != nil {
		return fmt.Errorf("p.FromBytes: %w", err)
	}

	return nil
}
