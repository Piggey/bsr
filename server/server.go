package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/game"
	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/packet/binary"
	"github.com/Piggey/bsr/util"
	"github.com/google/uuid"
)

type Server struct {
	addr        net.Addr
	udpConn     *net.UDPConn
	logger      *slog.Logger
	activeGames map[uuid.UUID]activeGame
}

func NewServer(addr string) *Server {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}

	serverHandler := util.NewSlogHandler("server", udpAddr.String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(serverHandler)

	logger.Info("server created")

	return &Server{
		addr:        udpAddr,
		udpConn:     udpConn,
		logger:      logger,
		activeGames: map[uuid.UUID]activeGame{},
	}
}

func (s *Server) Close() error {
	s.logger.Info("connection closed")
	return s.udpConn.Close()
}

func (s *Server) Listen() error {
	s.logger.Info("started listening")

	for {
		// wait for client to create a new game
		ngp, hostAddr, err := s.awaitCreateNewGame()
		if err != nil {
			return fmt.Errorf("awaitCreateNewGame: %v", err)
		}

		go s.startNewGame(hostAddr, ngp)
	}
}

func (s *Server) awaitCreateNewGame() (packet.CreateGamePacket, net.Addr, error) {
	ngp, addr, err := readPacket[packet.CreateGamePacket](s.udpConn)
	if err != nil {
		return ngp, addr, fmt.Errorf("readAndValidate: %v", err)
	}

	return ngp, addr, nil
}

func (s *Server) startNewGame(hostAddr net.Addr, ngp packet.CreateGamePacket) error {
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

	err := binary.Write(s.udpConn, binary.BigEndian, gsp)
	if err != nil {
		return fmt.Errorf("udp write: %v", err)
	}

	return nil
}

func readPacket[T packet.Packet](conn *net.UDPConn) (T, net.Addr, error) {
	var p T
	addr, err := binary.ReadFrom(conn, binary.BigEndian, &p)
	if err != nil {
		return p, nil, fmt.Errorf("udp read: %v", err)
	}

	if err := p.Validate(); err != nil {
		return p, addr, fmt.Errorf("packet validate: %v", err)
	}

	return p, addr, nil
}
