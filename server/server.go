package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"sync"

	pb "github.com/Piggey/bsr/proto"
	"github.com/Piggey/bsr/util"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBsrServer
	lis    net.Listener
	srv    *grpc.Server
	games  sync.Map
	logger *slog.Logger
}

func NewServer(addr string) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}

	serverHandler := util.NewSlogHandler("server", lis.Addr().String(), os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(serverHandler)

	logger.Info("new server created")

	srv := grpc.NewServer()
	s := Server{
		lis:    lis,
		srv:    srv,
		games:  sync.Map{},
		logger: logger,
	}
	pb.RegisterBsrServer(srv, &s)

	return &s, nil
}

func (s *Server) Close() error {
	s.srv.GracefulStop()
	return s.lis.Close()
}

func (s *Server) Listen() error {
	s.logger.Info("started listening")
	return s.srv.Serve(s.lis)
}
