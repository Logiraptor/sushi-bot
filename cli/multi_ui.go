package main

import (
	"github.com/Logiraptor/sushi-bot/core"
	"github.com/Logiraptor/sushi-bot/game"
)

type multiUI struct {
	uis []game.UI
}

func NewMultiUI(uis ...game.UI) *multiUI {
	return &multiUI{uis: uis}
}

func (l *multiUI) PlayerAvailableCardsChanged(player int, hand []core.Card) {
	for _, ui := range l.uis {
		ui.PlayerAvailableCardsChanged(player, hand)
	}
}

func (l *multiUI) PlayerChosenCardsChanged(player int, hand []core.Card) {
	for _, ui := range l.uis {
		ui.PlayerChosenCardsChanged(player, hand)
	}
}

func (l *multiUI) ScoreChanged(player int, score core.Score) {
	for _, ui := range l.uis {
		ui.ScoreChanged(player, score)
	}
}

func (l *multiUI) RoundOver() {
	for _, ui := range l.uis {
		ui.RoundOver()
	}
}

func (l *multiUI) GameOver() {
	for _, ui := range l.uis {
		ui.GameOver()
	}
}

func (l *multiUI) DeckChanged(deck []core.Card) {
	for _, ui := range l.uis {
		ui.DeckChanged(deck)
	}
}
