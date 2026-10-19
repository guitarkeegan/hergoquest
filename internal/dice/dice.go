package dice

import (
	"math/rand"
	"time"
)

// responsible for returning numbers for 1 or 2 dice
// rolling the action dice and returning results
// a struct returning the dice functions?
// can be called from wherever

type combat int

type InputError struct{}

type CombatResult struct {
	Attack, GoodDefend, EvilDefend int
}

func (ie InputError) Error() string {
	return "dice: input error"
}

const (
	attack combat = iota
	goodDefend
	evilDefend
)

type Rollist struct {
	rand *rand.Rand
}

func (r *Rollist) rollNumberDie() int {
	return r.rand.Intn(6) + 1
}

func (r *Rollist) rollCombatDie() combat {
	roll := r.rand.Intn(6) + 1
	if roll >= 1 && roll <= 3 {
		return attack
	} else if roll >= 4 && roll <= 5 {
		return goodDefend
	} else {
		return evilDefend
	}
}

// number of the red number die to roll. can only be 1 or 2.
// if the number passed is too large or small, we return the
// number passed with the error.
func (r *Rollist) RollNumber(numberOfDie int) (int, error) {
	if numberOfDie <= 0 || numberOfDie > 2 {
		return numberOfDie, InputError{}
	}
	var res int
	for range numberOfDie {
		res += r.rollNumberDie()
	}
	return res, nil
}

func (r *Rollist) RollCombat(numberOfDie int) (*CombatResult, error) {
	cr := &CombatResult{}
	if numberOfDie <= 0 || numberOfDie > 6 {
		return nil, InputError{}
	}
	for range numberOfDie {
		switch r.rollCombatDie() {
		case attack:
			cr.Attack++
		case goodDefend:
			cr.GoodDefend++
		case evilDefend:
			cr.EvilDefend++
		default:
			panic("RollCombat: default should never fire\n")
		}
	}
	return cr, nil
}

func New() *Rollist {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Rollist{
		rand: randGen,
	}
}
