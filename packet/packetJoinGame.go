package packet

import "github.com/Piggey/bsr/game"

// TODO: go-validator on this hoe
type JoinGameReq struct {
	ProtocolVersion uint8
	GameMode        game.Mode
	PlayersNo       uint8
}

func (JoinGameReq) Type() Type {
	return TypeJoinGameReq
}

type JoinGameRes struct {
	Code     StatusCode
	PlayerId int
}

func (JoinGameRes) Type() Type {
	return TypeJoinGameRes
}
