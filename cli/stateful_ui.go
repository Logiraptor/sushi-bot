package main

import (
	"github.com/Logiraptor/sushi-bot/core"
)

type statefulUI struct {
	availableCards map[int][]core.Card
	chosenCards    map[int][]core.Card
	score          map[int]core.Score
	discardPile    []core.Card
	deck           []core.Card
	gameOver       bool
}

func NewStatefulUI() *statefulUI {
	return &statefulUI{
		availableCards: make(map[int][]core.Card),
		chosenCards:    make(map[int][]core.Card),
		score:          make(map[int]core.Score),
		discardPile:    nil,
	}
}

func (l *statefulUI) PlayerAvailableCardsChanged(player int, hand []core.Card) {
	l.availableCards[player] = hand
}

func (l *statefulUI) PlayerChosenCardsChanged(player int, hand []core.Card) {
	l.chosenCards[player] = hand
}

func (l *statefulUI) ScoreChanged(player int, score core.Score) {
	l.score[player] = score
}

func (l *statefulUI) RoundOver() {}

func (l *statefulUI) GameOver() {
	l.gameOver = true
}

func (l *statefulUI) DeckChanged(deck []core.Card) {
	l.deck = deck
}