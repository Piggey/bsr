package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/util"
)

type Server struct {
	conn   net.PacketConn
	logger *slog.Logger
}

func NewServer(network, addr string) (*Server, error) {
	conn, err := net.ListenPacket(network, addr)
	if err != nil {
		return nil, fmt.Errorf("net.ListenPacket: %w", err)
	}

	serverHandler := util.NewSlogHandler("server", conn.LocalAddr().String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(serverHandler)
	logger.Info("server created")

	return &Server{
		conn:   conn,
		logger: logger,
	}, nil
}

func (s *Server) Close() error {
	return s.conn.Close()
}

func (s *Server) Listen() error {
	for {
		packetType, addr, err := s.readPacketType()
		if err != nil {
			return fmt.Errorf("s.readPacketType: %w", err)
		}

		switch packetType {
		case packet.PacketTypeProtocolHandshakeRequest:
			err = s.handleProtocolHandshake(addr)
		}
		if err != nil {
			return fmt.Errorf("%s: %w", packetType, err)
		}
	}
}

func (s *Server) readPacketType() (packet.PacketType, net.Addr, error) {
	buf := make([]byte, 1)
	n, addr, err := s.conn.ReadFrom(buf)
	if err != nil {
		return 0, nil, fmt.Errorf("conn.ReadFrom: %w", err)
	}
	if n != 1 {
		return 0, nil, fmt.Errorf("read %d bytes instead of 1", n)
	}

	return packet.PacketType(buf[0]), addr, nil
}

func (s *Server) handleProtocolHandshake(addr net.Addr) error {
	pphReq, err := packet.ReadPacket[*packet.PacketProtocolHandshakeRequest](s.conn)
	if err != nil {
		return fmt.Errorf("packet.ReadPacket[*packet.PacketProtocolHandshake]: %w", err)
	}

	pphRes := packet.PacketProtocolHandshakeResponse{}
	switch pphReq.ProtocolVersion {
	case packet.ProtoV1:
		pphRes.Status = packet.HandshakeStatusOK
	default:
		pphRes.Status = packet.HandshakeStatusProtocolNotSupported
	}

	err = packet.SendPacket(s.conn, addr, &pphRes)
	if err != nil {
		return fmt.Errorf("packet.SendPacket: %w", err)
	}

	return nil
}
