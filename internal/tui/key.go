package tui

import "charm.land/bubbles/v2/key"

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},
		{k.Quit},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type actionKeysMap struct {
	Move         key.Binding
	Attack       key.Binding
	CastSpell    key.Binding
	SearchTraps  key.Binding
	SearchChests key.Binding
	SearchDoors  key.Binding
}

var actionKeys = actionKeysMap{
	Move: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "move"),
	),
	Attack: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "attack"),
	),
	CastSpell: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "cast spell"),
	),
	SearchTraps: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "search traps"),
	),
	SearchChests: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "search chests"),
	),
	SearchDoors: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "search doors"),
	),
}

func (k actionKeysMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k actionKeysMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Move, k.Attack, k.CastSpell},
		{k.SearchChests, k.SearchTraps, k.SearchDoors},
	}
}
