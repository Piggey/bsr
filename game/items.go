package game

type GameItem uint8

const (
	ItemMagnifyingGlass GameItem = 1
	ItemCigarettes      GameItem = 2
	ItemBeer            GameItem = 3
	ItemHandsaw         GameItem = 4
	ItemHandcuffs       GameItem = 5
)

func (gi GameItem) String() string {
	str := ""
	switch gi {
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
