package monster

type MonsterClass int

const (
	skeleton MonsterClass = iota
	zombie
)

type Monster struct {
	Typ    MonsterClass
	hp     int
	mp     int
	attack int
	defend int
}

var monsters = map[MonsterClass]Monster{
	skeleton: {
		Typ:    skeleton,
		hp:     1,
		mp:     0,
		attack: 2,
		defend: 2,
	},
	zombie: {
		Typ:    zombie,
		hp:     1,
		mp:     0,
		attack: 2,
		defend: 3,
	},
}

func NewDefaultMonster(mc MonsterClass) *Monster {
	m, ok := monsters[mc]
	if !ok {
		return nil
	}

	return &m
}
