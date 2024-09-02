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
	s.logger.Info("server closing")
	return s.conn.Close()
}

func (s *Server) Listen() error {
	s.logger.Info("started listening")

	for {
		packetType, addr, err := s.readPacketType()
		if err != nil {
			return fmt.Errorf("s.readPacketType: %w", err)
		}

		s.logger.Info("received packet", slog.String("packet", packetType.String()))

		switch packetType {
		case packet.PacketTypeProtocolHandshakeRequest:
			err = s.handleProtocolHandshake(addr)
		}
		if err != nil {
			return fmt.Errorf("%s: %w", packetType, err)
		}
	}
}

func (s *Server) handleProtocolHandshake(addr net.Addr) error {
	pphReq := packet.PacketProtocolHandshakeRequest{}
	_, err := s.readPacket(&pphReq)
	if err != nil {
		return fmt.Errorf("s.readPacket: %w", err)
	}

	pphRes := packet.PacketProtocolHandshakeResponse{}
	switch pphReq.ProtocolVersion {
	case packet.ProtoV1:
		pphRes.Status = packet.HandshakeStatusOK
	default:
		pphRes.Status = packet.HandshakeStatusProtocolNotSupported
	}

	err = s.writePacket(&pphRes, addr)
	if err != nil {
		return fmt.Errorf("s.writePacket: %w", err)
	}

	return nil
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

func (s *Server) readPacket(p packet.Packet) (net.Addr, error) {
	data := make([]byte, p.Size())
	n, addr, err := s.conn.ReadFrom(data)
	if err != nil {
		return nil, fmt.Errorf("conn.ReadFrom: %w", err)
	}
	if len(data) != n {
		return nil, fmt.Errorf("len(data) != n")
	}

	err = p.FromBytes(data)
	if err != nil {
		return nil, fmt.Errorf("p.FromBytes: %w", err)
	}

	return addr, nil
}

func (s *Server) writePacket(p packet.Packet, addr net.Addr) error {
	data, err := p.ToBytes()
	if err != nil {
		return fmt.Errorf("p.ToBytes: %w", err)
	}

	n, err := s.conn.WriteTo(data, addr)
	if err != nil {
		return fmt.Errorf("conn.WriteTo: %w", err)
	}
	if len(data) != n {
		return fmt.Errorf("len(data) != n")
	}

	return nil
}
