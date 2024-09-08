package game

type Mode uint8

const (
	ModePvE Mode = 1
	ModePvP Mode = 2
)

func (m Mode) String() (out string) {
	switch m {
	case ModePvE:
		out = "PvE"
	case ModePvP:
		out = "PvP"
	default:
		out = "Unknown"
	}

	return out
}
