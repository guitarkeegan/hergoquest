package view

import (
	"fmt"
	"strings"

	"github.com/guitarkeegan/musings/cmd/board"
)

func Render(board board.Board) {
	b := strings.Builder{}
	for i := range board {
		for j := range board[i] {
			b.WriteString(board[i][j].Glyph())
			b.WriteString(" ")
		}
		b.WriteString("\n")
	}
	fmt.Print(b.String())
}
