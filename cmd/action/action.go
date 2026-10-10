package action

type CombatAction int

const (
	attack CombatAction = iota
	defendGood
	defendEvil
)

type CombatDie []CombatAction
