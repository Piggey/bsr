package packet

type Type uint8

const (
	TypeJoinGameReq Type = 1
	TypeJoinGameRes Type = 2
)

func (t Type) String() (out string) {
	switch t {
	case TypeJoinGameReq:
		out = "JoinGameRequest"
	case TypeJoinGameRes:
		out = "JoinGameResponse"
	default:
		out = "Unknown"
	}

	return out
}
