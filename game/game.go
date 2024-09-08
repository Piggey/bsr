package game

import "fmt"

type Game struct {
	Round   uint8
	Shotgun shotgun
	players map[int]player
	done    bool
}

func NewGame() *Game {
	return &Game{
		Round:   0,
		Shotgun: newShotgun(),
	}
}

func (g *Game) AddPlayer(playerId int) error {
	if _, found := g.players[playerId]; found {
		return fmt.Errorf("player %d already exists", playerId)
	}

	g.players[playerId] = newPlayer()
	return nil
}

func (g *Game) Done() bool {
	return g.done
}
