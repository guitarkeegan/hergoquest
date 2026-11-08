package tui

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	b "github.com/guitarkeegan/hergoquest/internal/board"
	"github.com/guitarkeegan/hergoquest/internal/gamepack"
	"github.com/guitarkeegan/hergoquest/internal/hero"
)

type model struct {
	board  *b.Board
	form   *huh.Form
	width  int
	height int
}

// INFO: could also pass in flags for row, col, options
// in the future
func initialModel() model {
	brd := b.New(5, 6, b.WithGamePack(gamepack.Make()))
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}
	if m.form.State == huh.StateCompleted {
		// TODO: abstract
		selectedHero := m.form.Get(fkHero).(hero.HeroClass)
		h := hero.NewDefault(selectedHero)
		m.board.AddHero(h)
		// clear form
		// or set it to the next one..
		// m.form = nil
	}

	return m, cmd
}

func (m model) View() tea.View {
	var content string
	switch m.board.GameState() {
	case b.ChooseHeroes:
		content = m.form.View()
	case b.HeroTurn:
		s := m.board.View()
		content = renderMain(s)
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
