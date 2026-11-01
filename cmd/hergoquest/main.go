package main

import (
	"context"
	"fmt"
	"os"

	"github.com/daved/clic"
	"github.com/guitarkeegan/hergoquest/internal/tui"
)

// play game in one room
// new game only
// choose character
// player moves first
// can attack, move, be still
// monster will move to attack on each turn
// when monster is killed, the game is over

func runCLI(ctx context.Context) error {
	tui.Run()
	return nil
}

func main() {

	root := clic.NewFromFunc(runCLI, "hergo")

	cmd, err := root.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}

	cmd.Handle(context.Background())
	cmd.Usage()

}
