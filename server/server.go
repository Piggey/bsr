package server

import (
	"log/slog"
	"net"
	"net/netip"
	"os"

	"github.com/Piggey/bsr/util"
)

type Server struct {
	udpAddr *net.UDPAddr
	udpConn *net.UDPConn
	logger  *slog.Logger
}

func NewServer(addr string) *Server {
	addrPort := netip.MustParseAddrPort(addr)
	udpAddr := net.UDPAddrFromAddrPort(addrPort)

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}

	serverHandler := util.NewCustomHandler("server", udpAddr.String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(serverHandler)

	logger.Info("ok dziala?", slog.Bool("kutas", true))

	return &Server{
		udpAddr: udpAddr,
		udpConn: udpConn,
		logger:  logger,
	}
}

func (s *Server) Close() error {
	return s.udpConn.Close()
}

func (s *Server) Listen() {
	panic("to be implemented")
}
