package game

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type Game struct {
	Id                  uuid.UUID
	Round               uint8
	Shotgun             shotgun
	Player1             player
	Player2             player
	CurrentTurnPlayerId uint8
	done                bool
}

func NewGame() *Game {
	return &Game{
		Id:                  uuid.New(),
		Round:               0,
		Shotgun:             newShotgun(),
		Player1:             newPlayer(0),
		Player2:             newPlayer(1),
		CurrentTurnPlayerId: 0,
	}
}

func (g *Game) PlayerMove(m Move) error {
	if err := g.validateMove(m); err != nil {
		return fmt.Errorf("invalid move: %v", err)
	}

	if gameDone := g.move(m); gameDone {
		g.done = true
	}

	return nil
}

func (g *Game) Done() bool {
	return g.done
}

func (g *Game) move(m Move) bool {
	if m.moveType == MoveTypeShoot {
		return g.moveShoot(m)
	}

	return g.moveUseItem(m)
}

func (g *Game) validateMove(m Move) error {
	if m.moveType == MoveTypeShoot {
		return g.validateMoveShoot(m)
	}

	return g.validateMoveUseItem(m)
}

func (g *Game) moveShoot(m Move) bool {
	dmg := g.Shotgun.Shoot()

	var playerShot player
	switch struct {
		playerId uint8
		detail   MoveDetail
	}{m.playerId, m.moveDetail} {
	case struct {
		playerId uint8
		detail   MoveDetail
	}{0, MoveDetailShootSelf}, struct {
		playerId uint8
		detail   MoveDetail
	}{1, MoveDetailShootPlayer}:
		playerShot = g.Player1
	case struct {
		playerId uint8
		detail   MoveDetail
	}{1, MoveDetailShootSelf}, struct {
		playerId uint8
		detail   MoveDetail
	}{0, MoveDetailShootPlayer}:
		playerShot = g.Player2
	}

	playerShot.health -= dmg
	return false
}

func (g *Game) moveUseItem(m Move) bool {
	return false
}

func (g *Game) validateMoveShoot(m Move) error {
	// its players turn
	if g.CurrentTurnPlayerId != m.playerId {
		return fmt.Errorf("its not player's %d turn (current turn: player %d)", m.playerId, g.CurrentTurnPlayerId)
	}

	// shotgun not empty
	if g.Shotgun.shellsLeft < 1 {
		return fmt.Errorf("shotgun is empty")
	}

	// valid move detail
	if m.moveDetail != MoveDetailShootSelf && m.moveDetail != MoveDetailShootPlayer {
		return fmt.Errorf("invalid move detail %d", m.moveDetail)
	}

	return nil
}

func (g *Game) validateMoveUseItem(m Move) error {
	// it's players turn
	if g.CurrentTurnPlayerId != m.playerId {
		return fmt.Errorf("its not player's %d turn (current turn: player %d)", m.playerId, g.CurrentTurnPlayerId)
	}

	// valid move detail
	if m.moveDetail < MoveDetailUseMagnifyingGlass || m.moveDetail > MoveDetailUseHandcuffs {
		return fmt.Errorf("invalid move detail %d", m.moveDetail)
	}

	// player has that item
	playerItems := g.Player1.items
	if m.playerId == 1 {
		playerItems = g.Player2.items
	}

	gameItem := GameItem(m.moveDetail)

	if !slices.Contains(playerItems[:], gameItem) {
		return fmt.Errorf("player %d does not have GameItem %s", m.playerId, gameItem)
	}

	return nil
}
