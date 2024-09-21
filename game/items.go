package game

type GameItem uint32

const (
	ItemNothing         GameItem = 0
	ItemMagnifyingGlass GameItem = 1
	ItemCigarettes      GameItem = 2
	ItemBeer            GameItem = 3
	ItemHandsaw         GameItem = 4
	ItemHandcuffs       GameItem = 5
)

func (gi GameItem) String() (str string) {
	switch gi {
	case ItemNothing:
		str = "Nothing"
	case ItemMagnifyingGlass:
		str = "MagnifyingGlass"
	case ItemCigarettes:
		str = "Cigarettes"
	case ItemBeer:
		str = "Beer"
	case ItemHandsaw:
		str = "Handsaw"
	case ItemHandcuffs:
		str = "Handcuffs"
	default:
		str = "Unknown GameItem"
	}

	return str
}
