package gamepack

import (
	b "github.com/guitarkeegan/hergoquest/internal/board"
	m "github.com/guitarkeegan/hergoquest/internal/monster"
)

var OneSkeleton = []b.MonsterPlacement{
	{
		Monster: m.NewDefaultMonster(m.Skeleton),
		Coord:   b.Coord{Row: 2, Col: 2},
	},
}

type Pack struct {
	monsters []b.MonsterPlacement
}

func (p Pack) Monsters() []b.MonsterPlacement {
	return p.monsters
}

func Make() Pack {
	return Pack{
		monsters: OneSkeleton,
	}
}
