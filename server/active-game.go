package server

import (
	"net"

	"github.com/Piggey/bsr/game"
	"github.com/Piggey/bsr/packet"
)

type activeGame struct {
	game        *game.Game
	gamemode    packet.GameMode
	player1Addr net.Addr
	player2Addr net.Addr
}
