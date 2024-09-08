package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/Piggey/bsr/game"
	"github.com/Piggey/bsr/packet"
)

type outPackets map[net.Addr][]packet.Packet

func (s *Server) handlePacket(p packet.Packet, addr net.Addr) (outp outPackets, err error) {
	switch v := p.(type) {
	case packet.JoinGameReq:
		outp, err = s.handleJoinGame(v, addr)
	default:
		panic("lol")
	}

	return outp, err
}

func (s *Server) handleJoinGame(jg packet.JoinGameReq, addr net.Addr) (outPackets, error) {
	if jg.ProtocolVersion != packet.ProtoV1 {
		return outPackets{
			addr: { packet.JoinGameRes{Code: packet.StatusProtocolNotSupported} },
		}, nil
	}

	if s.session != nil {
		if jg.GameMode != s.session.mode || jg.PlayersNo != s.session.playersNo || jg.ProtocolVersion != s.session.protoVer {
			return outPackets{
				addr: { packet.JoinGameRes{Code: packet.StatusIncorrectGameDetails} },
			}, nil
		}

		playerId := len(s.session.playerIds)

		err := s.session.game.AddPlayer(playerId)
		if err != nil {
			return nil, fmt.Errorf("game.AddPlayer: %w", err)
		}

		s.session.playerIds[addr] = playerId
		s.logger.Info("player joined session", slog.Int("playerId", playerId))

		return outPackets{
			addr: { packet.JoinGameRes{Code: packet.StatusOk, PlayerId: playerId} },
		}, nil
	}

	playerId := 0
	g := game.NewGame()
	err := g.AddPlayer(playerId)
	if err != nil {
		return nil, fmt.Errorf("g.AddPlayer: %w", err)
	}

	s.session = &session{
		protoVer:  jg.ProtocolVersion,
		game:      g,
		playerIds: map[net.Addr]int{addr: playerId},
		mode:      jg.GameMode,
		playersNo: jg.PlayersNo,
	}
	s.logger.Info("player created new session", slog.Int("playerId", playerId))

	return outPackets{
		addr: { packet.JoinGameRes{Code: packet.StatusOk, PlayerId: playerId} },
	}, nil
}
