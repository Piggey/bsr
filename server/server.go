package server

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Piggey/bsr/game"
	"github.com/Piggey/bsr/packet"
	"github.com/Piggey/bsr/util"
)

type Server struct {
	addr    net.Addr
	udpConn *net.UDPConn
	logger  *slog.Logger
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

	serverHandler := util.NewCustomHandler("server", udpAddr.String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(serverHandler)

	logger.Info("server created")

	return &Server{
		addr:    udpAddr,
		udpConn: udpConn,
		logger:  logger,
	}
}

func (s *Server) Close() error {
	return s.udpConn.Close()
}

func (s *Server) Listen() error {
	s.logger.Info("started listening")

	for {
		// wait for client to create a new game
		ngp, err := s.awaitCreateNewGame()
		if err != nil {
			return err
		}

		go s.startNewGame(ngp)
	}
}

func (s *Server) awaitCreateNewGame() (packet.CreateGamePacket, error) {
	ngp := packet.CreateGamePacket{}
	err := binary.Read(s.udpConn, binary.BigEndian, &ngp)
	if err != nil {
		return ngp, err
	}

	if err := ngp.Validate(); err != nil {
		s.logger.Error("invalid packet received", slog.Any("err", err))
		return ngp, err
	}

	s.logger.Info("received create new game packet", slog.Any("packet", ngp))
	return ngp, err
}

func (s *Server) startNewGame(ngp packet.CreateGamePacket) error {
	s.logger.Info("starting new game")

	g := game.NewGame()

	fmt.Printf("g: %v\n", g)
	return nil
}
