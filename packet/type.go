package packet

type Type uint8

const (
	TypeJoinGameReq Type = 1
)

func (t Type) String() string {
	out := ""
	switch t {
	case TypeJoinGameReq:
		out = "JoinGameRequest"
	default:
		out = "Unknown"
	}

	return out
}
