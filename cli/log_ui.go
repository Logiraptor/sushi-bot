package main

import (
	"github.com/Logiraptor/sushi-bot/core"
	"fmt"
)

type logUI struct {
}

func NewLogUI() *logUI {
	return &logUI{}
}

func (l *logUI) PlayerAvailableCardsChanged(player int, hand []core.Card) {
	fmt.Printf("PlayerAvailableCardsChanged(%v, %v)\n", player, hand)
}

func (l *logUI) PlayerChosenCardsChanged(player int, hand []core.Card) {
	fmt.Printf("PlayerChosenCardsChanged   (%v, %v)\n", player, hand)
}

func (l *logUI) ScoreChanged(player int, score core.Score) {
	fmt.Printf("ScoreChanged               (%v, %v)\n", player, score)
}

func (l *logUI) RoundOver() {
	fmt.Printf("RoundOver                  ()\n")
}

func (l *logUI) GameOver() {
	fmt.Printf("GameOver                   ()\n")
}

func (l *logUI) DeckChanged(deck []core.Card) {
	fmt.Printf("DeckChanged                (%v)\n", deck)
}