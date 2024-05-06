package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"

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
		_, addr, err := s.udpConn.ReadFrom(buf)
		if err != nil {
			return err
		}

		fmt.Printf("received packet from addr: %v\n", addr)

		// create new go routine for a game
	}
}

func (s *Server) StartNewGame() {
	panic("to be implemented")
}
