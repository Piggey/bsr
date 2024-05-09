package packet

import "fmt"

type GameMode uint8

const (
	GameModePvP GameMode = 1
	GameModePvE GameMode = 2
)

type CreateGamePacket struct {
	magic   [3]byte  // "bsr"
	version byte     // protocol version
	Mode    GameMode // pvp, pve
}

func NewCreateGamePacket(mode GameMode) CreateGamePacket {
	return CreateGamePacket{
		magic:   MagicBytes,
		version: CurrentVersion,
		Mode:    mode,
	}
}

func (p CreateGamePacket) Validate() error {
	if p.magic != MagicBytes {
		return fmt.Errorf("invalid magic bytes")
	}

	if p.version != CurrentVersion {
		return fmt.Errorf("invalid version")
	}

	if p.Mode > 2 {
		return fmt.Errorf("invalid game mode")
	}

	return nil
}
