package game

type Game struct {
	round uint8
}

func NewGame() *Game {
	return &Game{round: 0}
}
