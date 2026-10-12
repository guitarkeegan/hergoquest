package hero

type Hero struct {
	Typ           HeroClass
	hp            int
	mp            int
	attack        int
	defend        int
	hasMoved      bool
	hasAttacked   bool
	movementQuota int
}

func (h *Hero) AvailableMovement() int {
	return h.movementQuota
}

func (h *Hero) AddMovementQuota(moves int) {
	h.movementQuota = moves
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
			hp:     8,
			mp:     2,
			attack: 3,
			defend: 2,
		}
	case Dwarf:
		return &Hero{
			Typ:    Dwarf,
			hp:     7,
			mp:     3,
			attack: 2,
			defend: 2,
		}
	default:
		return &Hero{
			Typ:    Barbarian,
			hp:     8,
			mp:     2,
			attack: 3,
			defend: 2,
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
		hero.hp <= 0 {
		// TODO: assume typ will be valid for now
		return NewDefault(hero.Typ)
	}
	return hero
}
