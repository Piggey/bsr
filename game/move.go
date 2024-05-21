package game

type MoveType uint8

// move types
const (
	MoveTypeShoot   MoveType = 1
	MoveTypeUseItem MoveType = 2
)

type MoveDetail uint8

// move details
const (
	MoveDetailUseMagnifyingGlass MoveDetail = 1
	MoveDetailUseCigarettes      MoveDetail = 2
	MoveDetailUseBeer            MoveDetail = 3
	MoveDetailUseHandsaw         MoveDetail = 4
	MoveDetailUseHandcuffs       MoveDetail = 5
	MoveDetailShootSelf          MoveDetail = 6
	MoveDetailShootPlayer        MoveDetail = 7
)

type Move struct {
	playerId   uint8
	moveType   MoveType
	moveDetail MoveDetail
}
