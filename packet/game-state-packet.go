package packet

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
