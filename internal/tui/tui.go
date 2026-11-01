package tui

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	b "github.com/guitarkeegan/hergoquest/internal/board"
	"github.com/guitarkeegan/hergoquest/internal/hero"
)

// seems like i'd import board and store it here
type model struct {
	board *b.Board
	form  *huh.Form
}

// TODO: what to pass in?
func initialModel() model {
	brd := b.New(5, 6)
	return model{
		board: brd,
		form:  chooseHeroForm(),
	}
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}
	if m.form.State == huh.StateCompleted {
		selectedHero := m.form.Get(fkHero).(hero.HeroClass)
		h := hero.NewDefault(selectedHero)
		m.board.AddHero(h)
	}

	return m, cmd
}

func (m model) View() tea.View {
	var content string
	switch m.board.GameState() {
	case b.ChooseHeroes:
		content = m.form.View()
	case b.HeroTurn:
		content = m.board.View()
	case b.MonsterTurn:
		// TODO:
	default:
		content = "default case"
	}

	v := tea.NewView(content)
	v.AltScreen = true
	return v
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
