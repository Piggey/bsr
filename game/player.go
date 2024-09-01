package game

type player struct {
	name   string
	health uint8
	items  [8]GameItem
}

func newPlayer(name string) player {
	return player{
		name:   name,
		health: 2,
		items:  [8]GameItem{},
	}
}

func (p *player) Health() uint8 {
	return p.health
}

func (p *player) Items() [8]uint8 {
	out := [8]uint8{}
	for i := 0; i < len(p.items); i++ {
		out[i] = uint8(p.items[i])
	}

	return out
}
