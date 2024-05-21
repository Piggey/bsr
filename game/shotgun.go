package game

import "math/rand"

type shotgun struct {
	shellsLeft uint8
	chamber    uint8 // 1 for live, 0 for blank
	dmg        uint8
}

func newShotgun() shotgun {
	shells := uint8(rand.Int()%(8-2+1) + 2) // [2; 8]
	chamber := setChamber(shells)

	return shotgun{
		shellsLeft: shells,
		chamber:    chamber,
		dmg:        1,
	}
}

// returns dmg dealt, 0 when blank, 1 when normal, 2 when handsaw used
func (s *shotgun) Shoot() uint8 {
	shell := s.chamber & 1

	s.chamber >>= 1
	s.shellsLeft -= 1

	return shell * s.dmg
}

func (s *shotgun) LiveShells() uint8 {
	lives := uint8(0)
	chamber := s.chamber
	for i := 0; i < int(s.shellsLeft); i++ {
		if chamber&1 != 0 {
			lives += 1
		}

		chamber >>= 1
	}

	return lives
}

func (s *shotgun) BlankShells() uint8 {
	blanks := uint8(0)
	chamber := s.chamber
	for i := 0; i < int(s.shellsLeft); i++ {
		if chamber&1 == 0 {
			blanks += 1
		}

		chamber >>= 1
	}

	return blanks
}

func setChamber(shells uint8) uint8 {
	chamber := uint8(0)

	for i := 0; i < int(shells); i++ {
		isLive := rand.Int()%2 == 0 // 50% change for live
		if isLive {
			chamber += 1
		}
		chamber <<= 1
	}

	return chamber
}
