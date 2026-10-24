// what is this?
// pieces that don't move
package piece

type Piece int

const (
	// furnature
	Table Piece = iota
	Throne
	AlchemistsBench
	TreasureChest
	Tomb
	SorcerersTable
	Bookcase
	Rack
	Fireplace
	WeaponsRack
	Cupboard

	// door states...
	OpenDoor
	ClosedDoor

	// obsticals
	Stairs
	BlockedSquare
	PitTrap
	SecretDoor
	FallingBlockTrap
)
