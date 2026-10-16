package view

import (
	"fmt"
	"strings"

	"github.com/guitarkeegan/hergoquest/internal/board"
)

func Render(brd *board.Board) {
	b := strings.Builder{}
	for i := range brd.Rows() {
		for j := range brd.Cols() {
			b.WriteString(brd.Get(board.Coord{Row: i, Col: j}).Glyph())
			b.WriteString(" ")
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")

	fmt.Print(b.String())
}
