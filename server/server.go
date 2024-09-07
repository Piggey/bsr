package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"sync"

	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/util"
)

type Server struct {
	conn           net.PacketConn
	logger         *slog.Logger
	clientSessions sync.Map
	gameSessions   sync.Map
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
		conn:           conn,
		logger:         logger,
		clientSessions: sync.Map{},
		gameSessions:   sync.Map{},
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
			// czy chce tu haltowac caly serwer w sumie?
			s.logger.Error("couldnt read packet", slog.Any("err", err))
			continue
		}

		sessionId, ok := s.getSessionId(addr, p)
		if !ok {
			s.logger.Warn("couldnt find sessionId for client", slog.Any("addr", addr))
			continue
		}

		gameSession, ok := s.getGameSession(sessionId, p)
		if !ok {
			s.logger.Warn("couldnt find game session for sessionId", slog.Any("sessionId", sessionId))
			continue
		}
		_ = gameSession
	}
}

func (s *Server) getGameSession(sessionId uint8, p packet.Packet) (gameSession, bool) {
	panic("implement")
}

func (s *Server) getSessionId(addr net.Addr, p packet.Packet) (uint8, bool) {
	sessionId, ok := s.clientSessions.Load(addr)
	if !ok {
		switch v := p.(type) {
		case packet.JoinGameReq:
			// wasnt stored yet
			s.logger.Info("storing new session id", slog.Any("addr", addr), slog.Any("sessionId", v.SessionId))
			s.clientSessions.Store(addr, v.SessionId)
			sessionId = v.SessionId
		default:
			return 0, false
		}
	}

	return sessionId.(uint8), true
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
