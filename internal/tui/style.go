package tui

import (
	"charm.land/lipgloss/v2"
)

// or maybe more general like render box
// combine h and w to align other menus
func renderBoard(board string) string {
	s := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(2, 2)

	return s.Render(board)
}

func renderMain(board string) string {
	return renderBoard(board)
}
