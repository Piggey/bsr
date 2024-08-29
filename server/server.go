package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/game"
	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/util"
)

type Server struct {
	conn        net.PacketConn
	logger      *slog.Logger
	activeGames map[uint32]activeGame
}

func NewServer(addr string) (*Server, error) {
	conn, err := net.ListenPacket("udp", addr)
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
		conn:        conn,
		logger:      logger,
		activeGames: map[uint32]activeGame{},
	}, nil
}

func (s *Server) Close() error {
	s.logger.Info("connection closed")
	return s.conn.Close()
}

func (s *Server) Listen() error {
	s.logger.Info("started listening")

	for {
		// wait for client to create a new game
		ngp := packet.CreateGamePacket{}
		clientAddr, err := s.ReadFrom(&ngp)
		if err != nil {
			return fmt.Errorf("s.ReadFrom: %w", err)
		}

		go s.startNewGame(clientAddr, &ngp)
	}
}

func (s *Server) WriteTo(p packet.Packet, addr net.Addr) error {
	payload, err := p.ToBytes()
	if err != nil {
		return fmt.Errorf("p.ToBytes: %w", err)
	}

	s.logger.Debug("sending packet", slog.Int("packet length", len(payload)), slog.String("receiver addr", addr.String()))

	_, err = s.conn.WriteTo(payload, addr)
	if err != nil {
		return fmt.Errorf("conn.WriteTo: %w", err)
	}

	return nil
}

func (s *Server) ReadFrom(p packet.Packet) (net.Addr, error) {
	buf := make([]byte, 1024)

	n, addr, err := s.conn.ReadFrom(buf)
	if err != nil {
		return nil, fmt.Errorf("conn.ReadFrom: %w", err)
	}
	buf = buf[:n]
	s.logger.Debug("read packet", slog.Int("packet length", n), slog.String("sender addr", addr.String()))

	err = p.FromBytes(buf)
	if err != nil {
		return nil, fmt.Errorf("p.FromBytes: %w", err)
	}

	return addr, nil
}

func (s *Server) startNewGame(hostAddr net.Addr, ngp *packet.CreateGamePacket) error {
	g := game.NewGame()
	s.logger.Info("started new game", slog.Any("game id", g.Id))

	ag := activeGame{
		game:        g,
		gamemode:    ngp.Mode,
		player1Addr: hostAddr,
		player2Addr: nil,
	}

	s.activeGames[g.Id] = ag

	// send game state packet to all players
	gsp := packet.GameStatePacket{
		Round:         g.Round,
		Player1Health: g.Player1.Health(),
		Player1Items:  g.Player1.Items(),
		Player2Health: g.Player1.Health(),
		Player2Items:  g.Player2.Items(),
		ShotgunLive:   g.Shotgun.LiveShells(),
		ShotgunBlank:  g.Shotgun.BlankShells(),
		PlayerTurn:    g.CurrentTurnPlayerId,
	}

	err := s.WriteTo(&gsp, hostAddr)
	if err != nil {
		return fmt.Errorf("s.WriteTo: %w", err)
	}

	return nil
}
