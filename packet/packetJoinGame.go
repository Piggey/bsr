package packet

import "github.com/Piggey/bsr/game"

type JoinGameReq struct {
	ProtocolVersion uint8
	GameMode        game.Mode
	NoPlayers       uint8
}

func (JoinGameReq) Type() Type {
	return TypeJoinGameReq
}

type JoinGameRes struct {
	Code     StatusCode
	PlayerId uint8
}

func (JoinGameRes) Type() Type {
	return TypeJoinGameRes
}
