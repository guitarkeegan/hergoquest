package board

import (
	"errors"
	"math"
	"strings"

	"github.com/guitarkeegan/hergoquest/internal/dice"
	"github.com/guitarkeegan/hergoquest/internal/log"
)

var DIRECTIONS = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type Coord struct {
	Row, Col int
}

type Piece int

const (
	vacant Piece = iota
	closedDoor
	openDoor
)

func (p Piece) Glyph() string {
	switch p {
	case vacant:
		return "[]"
	case closedDoor:
		return "--"
	case openDoor:
		return "()"
	default:
		return "??"
	}
}

func (p Piece) AvailableMovement() int {
	return 0
}

type Occupier interface {
	Glyph() string
	AvailableMovement() int
}

type Combatant interface {
	Occupier
	Attack() int
	Defend() int
	IsDead() bool
	// this is really for monsters, might need a
	// better way to handle this
	IsVisible() bool
	TakeDamage(int)
	RemainingHP() int
	Moved(...bool) bool
	// attacked, searched for traps, treasure, cast a spell
	Acted(...bool) bool
	// TODO: will need to think on this one
	// Spell() string
}

type GameState string

const (
	ChooseHeroes GameState = "chooseHeroes"
	HeroTurn     GameState = "hero"
	MonsterTurn  GameState = "monster"
)

type Board struct {
	cells              [][]Occupier
	walls              map[edge]bool
	heroes             []Combatant
	monsters           []Combatant
	combatantPositions map[Combatant]Coord
	gameState          GameState
	dice               *dice.Rollist
}

func (b *Board) GameState() GameState {
	return b.gameState
}

func (brd *Board) View() string {
	rows := make([]string, 0, brd.Rows())
	for i := range brd.Rows() {
		cells := make([]string, 0, brd.Cols())
		for j := range brd.Cols() {
			cells = append(cells, brd.Get(Coord{Row: i, Col: j}).Glyph())
		}
		rows = append(rows, strings.Join(cells, " "))
	}
	return strings.Join(rows, "\n")
}

func less(a, b Coord) bool {
	if a.Row != b.Row {
		return a.Row < b.Row
	}
	return a.Col < b.Col
}

type Door int

const (
	none Door = iota
	closed
	open
)

// edge is the border shared by two adjacent cells, a always being the lesser coord
type edge struct {
	a, b Coord
	door Door
}

func (e *edge) Door() Door {
	return e.door
}

type edgeOption func(*edge)

func DoorOption(d Door) edgeOption {
	return func(e *edge) {
		e.door = d
	}
}

// Edges must be created with newEdge so that
// the order of the coordinances are always sorted
func newEdge(c1, c2 Coord, opt ...edgeOption) edge {
	if less(c2, c1) {
		c1, c2 = c2, c1
	}
	e := edge{a: c1, b: c2}
	for _, f := range opt {
		f(&e)
	}
	return e
}

func (b *Board) IsWall(c1, c2 Coord) bool {
	return b.walls[newEdge(c1, c2)]
}

// turn is derived from the heros available movements
// INFO: what if a hero decided not to move?
// we must make them exhast their movemnts in that case.
func (b *Board) Turn() GameState {
	for _, h := range b.heroes {
		if h.AvailableMovement() > 0 {
			return HeroTurn
		}
	}
	return MonsterTurn
}

var (
	ErrOutOfBounds       = errors.New("coord is out of bounds")
	ErrNoOccupant        = errors.New("no occupant to move")
	ErrNonCombatant      = errors.New("not a combatant interface")
	ErrCombatantNotFound = errors.New("combatant not found")
)

func (b *Board) constructWall(c1, c2 Coord) {
	b.walls[newEdge(c1, c2)] = true
}

// TODO: maybe hand a room over, and get back the walls?
func (b *Board) constructOuterWalls(cells [][]Occupier) {

	// visit each cell
	// check all 4 directions around it
	// if out of bounds, add to the map
	for i := range cells {
		for j := range cells[i] {
			for _, dir := range DIRECTIONS {
				nextRow, nextCol := dir[0]+i, dir[1]+j
				if nextRow < 0 || nextCol < 0 {
					// make left or top wall
					b.constructWall(Coord{Row: nextRow, Col: nextCol}, Coord{Row: i, Col: j})
				} else if nextRow >= len(cells) || nextCol >= len(cells[i]) {
					// make right or bottom wall
					b.constructWall(Coord{Row: i, Col: j}, Coord{Row: nextRow, Col: nextCol})
				}
			}
		}
	}
}

// TODO: maybe add a config for the start of games...
func New(rows, cols int, opt ...BoardOption) *Board {
	cells := make([][]Occupier, rows)
	for i := range cells {
		cells[i] = make([]Occupier, cols)
		for j := range cells[i] {
			cells[i][j] = vacant
		}
	}
	board := &Board{
		cells:              cells,
		heroes:             []Combatant{},
		monsters:           []Combatant{},
		combatantPositions: make(map[Combatant]Coord),
		gameState:          ChooseHeroes,
		dice:               dice.New(),
		walls:              make(map[edge]bool),
	}
	board.constructOuterWalls(cells)
	for _, f := range opt {
		f(board)
	}
	// TODO: add rooms
	// TODO: add doors
	// TODO: furnature
	return board
}

// returns the number of rows on the board
func (b *Board) Rows() int {
	return len(b.cells)
}

// returns the number of cols on the board
func (b *Board) Cols() int {
	if len(b.cells) == 0 {
		return 0
	}
	return len(b.cells[0])
}

func (b *Board) inBounds(coord Coord) bool {
	return coord.Row >= 0 && coord.Row < b.Rows() &&
		coord.Col >= 0 && coord.Col < b.Cols()
}

func (b *Board) Remove(coord Coord) {
	b.cells[coord.Row][coord.Col] = vacant
}

// TODO: changes this to AddHeros
func (b *Board) AddHero(c Combatant) {
	// TODO: how to determine plaement
	// maybe a check here if the first place is taken, etc.
	b.placeHero(c, Coord{Row: 0, Col: 0})
}

func (b *Board) placeHero(c Combatant, coord Coord) {
	b.cells[coord.Row][coord.Col] = c
	b.heroes = append(b.heroes, c)
	b.combatantPositions[c] = coord
	// INFO: toriaizu
	b.gameState = HeroTurn
}

func (b *Board) PlaceMonster(c Combatant, coord Coord) {
	b.cells[coord.Row][coord.Col] = c
	b.monsters = append(b.monsters, c)
	b.combatantPositions[c] = coord
}

func (b *Board) Place(o Occupier, coord Coord) {
	b.cells[coord.Row][coord.Col] = o
}

func (b *Board) PlaceCombatant(c Combatant, coord Coord) {
	b.cells[coord.Row][coord.Col] = c
}

// returns the Occupier given a coord, to get the coord
// of a combatant, use GetCombatant
func (b *Board) Get(coord Coord) Occupier {
	return b.cells[coord.Row][coord.Col]
}

func (b *Board) GetCombatant(c Combatant) (Coord, error) {
	if coord, ok := b.combatantPositions[c]; ok {
		return coord, nil
	}
	// TODO: add to errors
	return Coord{}, ErrCombatantNotFound
}

// move to algos pkg
// TODO: will need to update when traps, doors, etc. exist
func isValidPath[T any](matrix [][]T, start, end [2]int, movesRemaining int) bool {
	ROWS := len(matrix)
	COLS := len(matrix[0])

	type Step struct {
		pos   [2]int
		moves int
	}

	q := []Step{
		{pos: start, moves: 0},
	}

	seen := make(map[[2]int]bool)
	seen[start] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, dir := range DIRECTIONS {
			nextRow := cur.pos[0] + dir[0]
			nextCol := cur.pos[1] + dir[1]

			next := [2]int{nextRow, nextCol}

			if nextRow < 0 || nextRow >= ROWS ||
				nextCol < 0 || nextCol >= COLS ||
				seen[next] {
				continue
			}

			nextMoves := cur.moves + 1

			if nextMoves > movesRemaining {
				continue
			}

			if next == end {
				return true
			}

			seen[next] = true

			q = append(q, Step{
				pos:   next,
				moves: nextMoves,
			})
		}
	}

	return false
}

func (b *Board) isValidMove(from, to Coord) bool {
	occupant := b.Get(from)
	return isValidPath(b.cells, [2]int{from.Row, from.Col}, [2]int{to.Row, to.Col}, occupant.AvailableMovement())
}

func (b *Board) Attack(attackCoord, defendCoord Coord) {

	// if close enough
	// WARN: allows for diagonal attack
	isAdjacent := math.Abs(float64(attackCoord.Row)-float64(defendCoord.Row)) <= 1 &&
		math.Abs(float64(attackCoord.Col)-float64(defendCoord.Col)) <= 1

	if !isAdjacent {
		announce.Log("no enemy in range")
		return
	}

	a, ok := b.Get(attackCoord).(Combatant)
	if !ok {
		//TODO: handle non combatant
		return
	}

	d, ok := b.Get(defendCoord).(Combatant)
	if !ok {
		//TODO: handle non combatant
		return
	}

	attack, err := b.dice.RollCombat(a.Attack())
	if err != nil {
		panic("handle roll combat err")
	}

	// updated attackers turn
	a.Acted(true)

	defend, err := b.dice.RollCombat(d.Defend())
	if err != nil {
		panic("handle roll combat err")
	}

	// hero or monster, derive from turn
	var difference int
	if b.gameState == HeroTurn {
		difference = attack.Attack - defend.EvilDefend
	} else {
		difference = attack.Attack - defend.GoodDefend
	}

	if difference > 0 {
		d.TakeDamage(difference)
		announce.Log("took damage, %d remaining.", d.RemainingHP())
	} else {
		announce.Log("no damage, %d remaining.", d.RemainingHP())
	}

	// TODO:  always call clean up?
	b.update()

}

// return a list of living heros/monsters
// WARN: might be better just to mark fallen as isDead()
func filterOutFallen[T comparable](fallen T, pile []T) []T {
	out := make([]T, 0, len(pile))
	for _, body := range pile {
		if fallen != body {
			out = append(out, body)
		}
	}
	return out
}

// can be generalized, but essentially clean up the
// fallen combatants
// TODO: add turn check
func (b *Board) update() {
	announce.Log("updating...")
	for _, h := range b.heroes {
		if h.IsDead() {
			if coord, err := b.GetCombatant(h); err == nil {
				b.Remove(coord)
				delete(b.combatantPositions, h)
				b.heroes = filterOutFallen(h, b.heroes)
			}
		}
	}
	for _, m := range b.monsters {
		if m.IsDead() {
			if coord, err := b.GetCombatant(m); err == nil {
				b.Remove(coord)
				delete(b.combatantPositions, m)
				b.monsters = filterOutFallen(m, b.monsters)
			}
		}
	}
	announce.Log("heros remaining: %d", len(b.heroes))
	announce.Log("monsters remaining: %d", len(b.monsters))

	switch b.gameState {
	case HeroTurn:
		var allDone bool
		for _, h := range b.heroes {
			if h.Acted() && h.Moved() {
				allDone = true
			} else {
				allDone = false
				break
			}
		}
		if allDone {
			// TODO: restart with the next turn
			b.gameState = MonsterTurn
		}
	case MonsterTurn:
		var allDone bool
		for _, m := range b.monsters {
			if m.Acted() && m.Moved() {
				allDone = true
			} else {
				allDone = false
				break
			}
		}
		if allDone {
			// TODO: restart with the next turn
			b.gameState = HeroTurn
		}
	}

	announce.Log("Turn: %s", b.gameState)

}

// Move reports whether the occupant at from could reach to, and moves it if so.
// An error is returned when the move can't be attempted at all.
// TODO: don't forget about occupied cells
func (b *Board) Move(from, to Coord) (bool, error) {
	if !b.inBounds(from) || !b.inBounds(to) {
		return false, ErrOutOfBounds
	}
	occupant := b.Get(from)
	if occupant == vacant {
		return false, ErrNoOccupant
	}
	if !b.isValidMove(from, to) {
		return false, nil
	}
	oc, ok := occupant.(Combatant)
	if !ok {
		return false, ErrNonCombatant
	}
	b.combatantPositions[oc] = to
	oc.Moved(true)
	b.Remove(from)
	b.Place(occupant, to)
	return true, nil
}
