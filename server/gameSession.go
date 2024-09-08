package server

import (
	"net"

	"github.com/Piggey/bsr/game"
)

type gameSession struct {
	g         *game.Game
	playerIds map[net.Addr]uint8
	mode      game.Mode
	noPlayers uint8
}
