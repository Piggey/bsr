package packet

import "encoding/json"

type GameStatePacket struct {
	magic         [3]uint8 // "bsr"
	version       uint8
	Round         uint8
	Player1Health uint8
	Player1Items  [8]uint8
	Player2Health uint8
	Player2Items  [8]uint8
	ShotgunLive   uint8
	ShotgunBlank  uint8
	PlayerTurn    uint8
}

func (p *GameStatePacket) FromBytes(b []byte) error {
	return json.Unmarshal(b, p)
}

func (p *GameStatePacket) ToBytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *GameStatePacket) Validate() error {
	// TODO: implementacja
	panic("unimplemented")
}
