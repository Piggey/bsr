package game

type player struct {
	lives uint8
	items []GameItem
}

func newPlayer() player {
	return player{
		lives: 2,
		items: []GameItem{},
	}
}
