package packet

type JoinGameReq struct {
	ProtocolVersion uint8
	SessionId       uint8
	GameMode        uint8
	NoPlayers       uint8
}

func (JoinGameReq) Type() Type {
	return TypeJoinGameReq
}
