package game

import (
	"sync"

	"github.com/google/uuid"
)

type Game struct {
	Round   uint8
	Shotgun shotgun
	Players map[string]player
	done    bool
	sync.Mutex
}

func NewGame() *Game {
	return &Game{
		Round:   0,
		Shotgun: newShotgun(),
		Players: map[string]player{},
	}
}

func (g *Game) AddPlayer(name string) string {
	playerUuid := uuid.NewString()
	g.Players[playerUuid] = newPlayer(name)

	return playerUuid
}

func (g *Game) Done() bool {
	return g.done
}
