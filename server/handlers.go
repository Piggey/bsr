package server

import (
	"net"

	"github.com/Piggey/bsr/packet"
)

func (s *Server) handlePacket(p packet.Packet, addr net.Addr) {
	switch v := p.(type) {
	case packet.JoinGameReq:
		s.handleJoinGameReq(v, addr)
	default:
		panic("lol")
	}
}

func (s *Server) handleJoinGameReq(jg packet.JoinGameReq, addr net.Addr) (map[net.Addr]packet.JoinGameRes, error) {
	if jg.ProtocolVersion != packet.ProtoV1 {
		return map[net.Addr]packet.JoinGameRes{
			addr: packet.JoinGameRes{Code: packet.StatusProtocolNotSupported}
		}, nil
	}

	if s.session != nil {
		if jg.GameMode != s.session.mode || jg.NoPlayers != s.session.noPlayers {
			return map[net.Addr]packet.JoinGameRes{
				addr: packet.JoinGameRes{Code: packet.StatusIncorrectGameDetails}
			}
		}

		s.session.playerIds[addr] = len(s.session.playerIds)
		
	}
}
