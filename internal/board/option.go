package board

type MonsterPlacement struct {
	Monster Combatant
	Coord   Coord
}

type GamePack interface {
	Monsters() []MonsterPlacement
}

func WithGamePack(gp GamePack) BoardOption {
	return func(b *Board) {
		for _, placement := range gp.Monsters() {
			b.PlaceMonster(placement.Monster, placement.Coord)
		}
	}
}

type BoardOption func(*Board)

// Doors can be open or closed.
// Will create a wall if it doesn't exist.
// Will fail silently if door is none.
func AddDoors(doors []edge) BoardOption {
	return func(b *Board) {
		for _, door := range doors {
			if door.door != none {
				b.walls[door] = true
			}
		}
	}
}
