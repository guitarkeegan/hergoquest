package view

import (
	"fmt"

	"github.com/guitarkeegan/musings/cmd/board"
)

func Render(board board.Board) {
	fmt.Println(board.View())
}
