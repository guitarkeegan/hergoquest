package monster

import (
	"log"

	"github.com/guitarkeegan/hergoquest/internal/dice"
)

type MonsterClass int

const (
	Skeleton MonsterClass = iota
	Zombie
)

type Monster struct {
	Typ           MonsterClass
	HP            int
	MP            int
	attack        int
	defend        int
	hasMoved      bool
	hasAttacked   bool
	movementQuota int
	dice          *dice.Rollist
}

// TODO: combine with hero's funcs so they can't differ
func (m *Monster) IsDead() bool {
	return m.HP <= 0
}

func (m *Monster) TakeDamage(damage int) {
	m.HP = max(m.HP-damage, 0)
}

func (m *Monster) AvailableMovement() int {
	return m.movementQuota
}

func (m *Monster) RemainingHP() int {
	return m.HP
}

func (m *Monster) Glyph() string {
	switch m.Typ {
	case Skeleton:
		return "sk"
	case Zombie:
		return "zb"
	default:
		return "??"
	}
}

var monsters = map[MonsterClass]Monster{
	Skeleton: {
		Typ:           Skeleton,
		HP:            1,
		MP:            0,
		attack:        2,
		defend:        2,
		movementQuota: 6,
		dice:          dice.New(),
	},
	Zombie: {
		Typ:           Zombie,
		HP:            1,
		MP:            0,
		attack:        2,
		defend:        3,
		movementQuota: 5,
		dice:          dice.New(),
	},
}

func NewDefaultMonster(mc MonsterClass) *Monster {
	m, ok := monsters[mc]
	if !ok {
		return nil
	}

	return &m
}

func (m *Monster) Attack() int {
	res, err := m.dice.RollCombat(m.attack)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	return res.Attack
}

func (m *Monster) Defend() int {
	res, err := m.dice.RollCombat(m.defend)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	return res.EvilDefend
}

//| Monster | Movement |
//|---|---:|
//| **Goblin** | **10** |
//| **Orc** | **8** |
//| **Dread Warrior** | **7** |
//| **Abomination** | **6** |
//| **Skeleton** | **6** |
//| **Gargoyle** | **6** |
//| **Zombie** | **5** |
//| **Mummy** | **4** |
