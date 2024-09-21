package game

import pb "github.com/Piggey/bsr/proto"

type player struct {
	name   string
	health uint32
	items  [8]GameItem
}

func newPlayer(name string) player {
	return player{
		name:   name,
		health: 2,
		items:  [8]GameItem{},
	}
}

func (p *player) Health() uint32 {
	return p.health
}

func (p *player) Items() []*pb.GameItem {
	out := make([]*pb.GameItem, 8)

	for i := 0; i < len(p.items); i++ {
		out[i] = &pb.GameItem{
			Id:   uint32(p.items[i]),
			Name: p.items[i].String(),
		}
	}

	return out
}
