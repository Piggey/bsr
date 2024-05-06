package packet

type GameMode byte

const (
	GameModePvP GameMode = 1
	GameModePvE GameMode = 2
)

type CreateGamePacket struct {
	Magic   [3]byte  // "bsr"
	Version byte     // protocol version
	Mode    GameMode // pvp, pve
}

func NewCreateGamePacket(mode GameMode) CreateGamePacket {
	return CreateGamePacket{
		Magic:   MagicBytes,
		Version: CurrentVersion,
		Mode:    mode,
	}
}
