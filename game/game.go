package game

import (
	"fmt"
	"sync"

	pb "github.com/Piggey/bsr/proto"
	"github.com/google/uuid"
)

type Game struct {
	mode           pb.GameMode
	Round          uint8
	Shotgun        shotgun
	players        map[string]player
	maxPlayerCount uint32
	done           bool
	sync.Mutex
}

func NewGame(mode pb.GameMode, maxPlayerCount uint32) *Game {
	return &Game{
		mode:           mode,
		Round:          0,
		Shotgun:        newShotgun(),
		players:        map[string]player{},
		maxPlayerCount: maxPlayerCount,
	}
}

func (g *Game) AddPlayer(name string) (string, error) {
	if len(g.players) == int(g.maxPlayerCount) {
		return "", fmt.Errorf("enough players joined")
	}

	playerUuid := uuid.NewString()
	g.players[playerUuid] = newPlayer(name)

	return playerUuid, nil
}

func (g *Game) MaxPlayerCount() int {
	return int(g.maxPlayerCount)
}

func (g *Game) PlayerCount() int {
	return len(g.players)
}

func (g *Game) Done() bool {
	return g.done
}
