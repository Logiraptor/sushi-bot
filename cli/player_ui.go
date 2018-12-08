package main

import (
	"github.com/Logiraptor/sushi-bot/core"
	"github.com/Logiraptor/sushi-bot/game"
)

type playerUI struct {
	player int
	ui     game.UI
}

func NewPlayerUI(player int, ui game.UI) *playerUI {
	return &playerUI{player, ui}
}

func (l *playerUI) PlayerAvailableCardsChanged(player int, hand []core.Card) {
	if l.player == player {
		l.ui.PlayerAvailableCardsChanged(player, hand)
	}
}

func (l *playerUI) PlayerChosenCardsChanged(player int, hand []core.Card) {
	l.ui.PlayerChosenCardsChanged(player, hand)
}

func (l *playerUI) ScoreChanged(player int, score core.Score) {
	l.ui.ScoreChanged(player, score)
}

func (l *playerUI) RoundOver() {
	l.ui.RoundOver()
}

func (l *playerUI) GameOver() {
	l.ui.GameOver()
}

func (l *playerUI) DeckChanged(deck []core.Card) {
	l.ui.DeckChanged(deck)
}
