package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"slices"

	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/util"
)

type Server struct {
	conn    net.PacketConn
	logger  *slog.Logger
	session *session
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
		p, addr, err := s.packetReadFrom()
		if err != nil {
			s.logger.Error("couldnt read packet", slog.Any("err", err))
			continue
		}

		outPackets, err := s.handlePacket(p, addr)
		if err != nil {
			s.logger.Error("couldnt handle packet", slog.Any("err", err))
			continue
		}

		for addr, packets := range outPackets {
			for _, p := range packets {
				err := s.packetWriteTo(addr, p)
				if err != nil {
					s.logger.Error("could write packet", slog.Any("type", p.Type()), slog.Any("addr", addr))
					continue
				}
			}
		}
	}
}

func (s *Server) packetReadFrom() (packet.Packet, net.Addr, error) {
	// read header
	buf := make([]byte, packet.HeaderSize)

	n, addr, err := s.conn.ReadFrom(buf)
	if err != nil {
		return nil, nil, fmt.Errorf("header: conn.ReadFrom: %w", err)
	}
	if n != packet.HeaderSize {
		return nil, nil, fmt.Errorf("header: conn.ReadFrom: n != packet.HeaderSize")
	}

	header, err := packet.UnmarshalHeader(buf)
	if err != nil {
		return nil, nil, fmt.Errorf("packet.UnmarshalHeader: %w", err)
	}

	// read packet
	buf = make([]byte, header.Size)
	n, addr2, err := s.conn.ReadFrom(buf)
	if err != nil {
		return nil, nil, fmt.Errorf("packet: conn.ReadFrom: %w", err)
	}
	if n != int(header.Size) {
		return nil, nil, fmt.Errorf("packet: conn.ReadFrom: n != int(header.Size)")
	}
	if addr != addr2 {
		return nil, nil, fmt.Errorf("packet: conn.ReadFrom: addr != addr2")
	}

	p, err := packet.Unmarshal(buf, header.Type)
	if err != nil {
		return nil, nil, fmt.Errorf("packet.Unmarshal: %w", err)
	}

	return p, addr, nil
}

func (s *Server) packetWriteTo(addr net.Addr, outp packet.Packet) error {
	outpBytes, err := packet.Marshal(outp)
	if err != nil {
		return fmt.Errorf("packet.Marshal: %w", err)
	}

	ph := packet.Header{
		Type: outp.Type(),
		Size: len(outpBytes),
	}
	phBytes, err := packet.MarshalHeader(ph)
	if err != nil {
		return fmt.Errorf("packet.MarshalHeader: %w", err)
	}

	n, err := s.conn.WriteTo(slices.Concat(phBytes, outpBytes), addr)
	if err != nil {
		return fmt.Errorf("conn.WriteTo: %w", err)
	}
	if len(outpBytes) + len(phBytes) != n {
		return fmt.Errorf("len(outpBytes) + len(phBytes) != n")
	}

	return nil
}
