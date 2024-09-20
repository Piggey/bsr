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

func (s *Server) JoinGame(ctx context.Context, jg *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	g, ok := s.getGame(jg.GameUuid)
	if !ok {
		if jg.MaxPlayerCount == nil {
			return nil, fmt.Errorf("s.createNewGame: player_count is nil")
		}

		g = s.createNewGame(jg.GameMode, jg.GameUuid, *jg.MaxPlayerCount)
	}
	g.Lock()
	defer g.Unlock()
	defer s.games.Store(jg.GameUuid, g)

	playerUuid, err := g.AddPlayer(jg.PlayerName)
	if err != nil {
		return nil, fmt.Errorf("g.AddPlayer: %w", err)
	}

	s.logger.Info("player joined game", slog.String("gameUuid", jg.GameUuid), slog.String("player", jg.PlayerName))

	return &pb.JoinGameResponse{
		GameUuid:    jg.GameUuid,
		PlayerUuid:  playerUuid,
		GameStarted: g.PlayerCount() == g.MaxPlayerCount(),
	}, nil
}

func (s *Server) createNewGame(mode pb.GameMode, id string, maxPlayerCount uint32) *game.Game {
	g := game.NewGame(mode, maxPlayerCount)
	s.games.Store(id, g)
	return g
}

func (s *Server) getGame(gameUuid string) (g *game.Game, ok bool) {
	gameAny, ok := s.games.Load(gameUuid)
	if !ok {
		return nil, false
	}

	return gameAny.(*game.Game), true
}
