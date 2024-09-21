package game

import (
	"fmt"
	"sync"

	pb "github.com/Piggey/bsr/proto"
	"github.com/google/uuid"
)

type Game struct {
	Round          uint32
	Shotgun        shotgun
	players        map[string]player
	maxPlayerCount uint32
	done           bool
	sync.Mutex
}

func NewGame(maxPlayerCount uint32) *Game {
	return &Game{
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

func (g *Game) ToGameState() *pb.GameState {
	pbPlayers := make(map[string]*pb.Player, len(g.players))
	for playerUuid, p := range g.players {
		pbPlayers[playerUuid] = &pb.Player{
			Health: p.health,
			Items:  p.Items(),
		}
	}

	return &pb.GameState{
		Round: g.Round,
		Shotgun: &pb.Shotgun{
			ShellsLeft: g.Shotgun.shellsLeft,
			Dmg:        g.Shotgun.dmg,
		},
		Players: pbPlayers,
		Done:    g.done,
	}
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
