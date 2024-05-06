package game

import "math/rand"

type shotgun struct {
	shellsLeft uint8
	chamber    uint8 // 1 for live, 0 for blank
}

func newShotgun() shotgun {
	shells := uint8(rand.Intn(8) + 1)
	chamber := uint8(rand.Intn(256))

	return shotgun{
		shellsLeft: shells,
		chamber:    chamber,
	}
}
