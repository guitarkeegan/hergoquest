package board

import h "github.com/guitarkeegan/musings/cmd/hero"

func FromHero(hero h.HeroClass) SquareState {
	switch hero {
	case h.Barbarian:
		return occupiedBarbarian
	default:
		return occupiedBarbarian
	}
}
