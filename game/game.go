package game

import "github.com/google/uuid"

type Game struct {
	Id      uuid.UUID
	round   uint8
	shotgun shotgun
	player1 player
	player2 player
}

func NewGame() *Game {
	return &Game{
		Id:      uuid.New(),
		round:   0,
		shotgun: newShotgun(),
		player1: newPlayer(),
		player2: newPlayer(),
	}
}
