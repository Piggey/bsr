package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"sync"

	"github.com/Piggey/bsr/game"
	pb "github.com/Piggey/bsr/proto"
	"github.com/Piggey/bsr/util"
	"github.com/google/uuid"
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
	s.logger.Info("closing")

	s.srv.GracefulStop()
	return s.lis.Close()
}

func (s *Server) Listen() error {
	s.logger.Info("started listening")
	return s.srv.Serve(s.lis)
}

func (s *Server) CreateGame(ctx context.Context, cg *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	s.logger.Info("creating a new game")

	gameUuid := uuid.NewString()
	g := game.NewGame()
	playerUuid := g.AddPlayer(cg.PlayerName)

	s.games.Store(gameUuid, g)

	return &pb.CreateGameResponse{
		Version:    pb.BsrProtoV1,
		GameUuid:   gameUuid,
		PlayerUuid: playerUuid,
	}, nil
}

func (s *Server) JoinGame(ctx context.Context, jg *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	g, ok := s.getGame(jg.GameUuid)
	if !ok {
		return nil, fmt.Errorf("game %s does not exist", jg.GameUuid)
	}
	g.Lock()
	defer g.Unlock()

	playerUuid := g.AddPlayer(jg.PlayerName)
	s.games.Store(jg.GameUuid, g)

	s.logger.Info("player joined game", slog.String("gameUuid", jg.GameUuid), slog.String("player", jg.PlayerName))

	return &pb.JoinGameResponse{
		Version:    pb.BsrProtoV1,
		GameUuid:   jg.GameUuid,
		PlayerUuid: playerUuid,
	}, nil
}

func (s *Server) getGame(gameUuid string) (g *game.Game, ok bool) {
	gameAny, ok := s.games.Load(gameUuid)
	if !ok {
		return nil, false
	}

	return gameAny.(*game.Game), true
}
