package hero

import (
	"github.com/guitarkeegan/hergoquest/internal/dice"
)

type Hero struct {
	Typ           HeroClass
	HP            int
	mp            int
	attack        int
	defend        int
	hasMoved      bool
	hasAttacked   bool
	movementQuota int
	dice          *dice.Rollist
}

// gets or sets the hero's movement.
// only the first argument is used
func (h *Hero) Moved(moved ...bool) bool {
	if len(moved) == 0 {
		return h.hasMoved
	}
	h.hasMoved = moved[0]
	return h.hasMoved
}

// gets or sets the heros attack status for this turn
func (h *Hero) Acted(attacked ...bool) bool {
	if len(attacked) == 0 {
		return h.hasAttacked
	}
	h.hasAttacked = attacked[0]
	return h.hasAttacked
}

func (h *Hero) AvailableMovement() int {
	return h.movementQuota
}

func (h *Hero) AddMovementQuota(moves int) {
	h.movementQuota = moves
}

// TODO:
// make the monster and isDead the same func if possible
// so they can't differ
func (h *Hero) IsDead() bool {
	return h.HP <= 0
}

func (h *Hero) TakeDamage(damage int) {
	h.HP = max(h.HP-damage, 0)
}

func (h *Hero) RemainingHP() int {
	return h.HP
}

// just return the number of combat dice to use
func (h *Hero) Attack() int {
	return h.attack
}

func (h *Hero) Defend() int {
	return h.defend
}

func (h *Hero) Glyph() string {
	switch h.Typ {
	case Barbarian:
		return "Ba"
	case Dwarf:
		return "Dw"
	default:
		return "??"
	}
}

type HeroClass string

const (
	Barbarian HeroClass = "barbarian"
	Wizard    HeroClass = "wizard"
	Elf       HeroClass = "elf"
	Dwarf     HeroClass = "dwarf"
)

type HeroOption func(*Hero)

func WithTyp(hc HeroClass) HeroOption {
	return func(h *Hero) {
		h.Typ = hc
	}
}

func NewDefault(hc HeroClass) *Hero {
	switch hc {
	case Barbarian:
		return &Hero{
			Typ:    Barbarian,
			HP:     8,
			mp:     2,
			attack: 3,
			defend: 2,
			dice:   dice.New(),
		}
	case Dwarf:
		return &Hero{
			Typ:    Dwarf,
			HP:     7,
			mp:     3,
			attack: 2,
			defend: 2,
			dice:   dice.New(),
		}
	default:
		return &Hero{
			Typ:    Barbarian,
			HP:     8,
			mp:     2,
			attack: 3,
			defend: 2,
			dice:   dice.New(),
		}
	}
}

// Create a new Hero with custom config. An invalid config will
// result in a hero with default stats.
func New(opt ...HeroOption) *Hero {
	hero := &Hero{}
	for _, f := range opt {
		f(hero)
	}
	// TODO: create range validation
	if hero.defend <= 0 || hero.attack <= 0 ||
		hero.HP <= 0 {
		// TODO: assume typ will be valid for now
		return NewDefault(hero.Typ)
	}
	return hero
}
