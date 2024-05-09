package game

import "github.com/google/uuid"

type Game struct {
	Id      uuid.UUID
	Round   uint8
	Shotgun shotgun
	Player1 player
	Player2 player
}

func NewGame() *Game {
	return &Game{
		Id:      uuid.New(),
		Round:   0,
		Shotgun: newShotgun(),
		Player1: newPlayer(0),
		Player2: newPlayer(1),
	}
}
