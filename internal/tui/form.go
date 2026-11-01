package tui

import (
	"charm.land/huh/v2"
	"github.com/guitarkeegan/hergoquest/internal/hero"
)

// form keys
const (
	fkHero = "hero"
)

func chooseHeroForm() *huh.Form {
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[hero.HeroClass]().Title("Choose your hero.").Key(fkHero).Options(
				huh.NewOption("Barbarian", hero.Barbarian),
				huh.NewOption("Dwarf", hero.Dwarf),
				huh.NewOption("Wizard", hero.Wizard),
				huh.NewOption("Elf", hero.Elf),
			),
		),
	)

	return f
}
