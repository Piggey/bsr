package server

import (
	"net"

	"github.com/Piggey/bsr/game"
)

type session struct {
	protoVer  uint8
	game      *game.Game
	playerIds map[net.Addr]int
	mode      game.Mode
	playersNo uint8
}
