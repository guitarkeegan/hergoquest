package dice

import "testing"

func TestRollCombatDice(t *testing.T) {

	rl := New()

	tests := []struct {
		input int
	}{
		{2},
		{3},
		{4},
	}

	for _, test := range tests {
		res, err := rl.RollCombat(test.input)
		if err != nil {
			t.Errorf("returned error on input: %d\n", test.input)
		}

		sum := res.attack + res.evilDefend + res.goodDefend
		if sum != test.input {
			t.Errorf("input does not equal total: input: %d, output sum: %d", test.input, sum)
		}
	}
}
