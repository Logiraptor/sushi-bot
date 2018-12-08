package main

import (
	"github.com/Logiraptor/sushi-bot/game"
	"fmt"
	"github.com/Logiraptor/sushi-bot/core"
	"math/rand"
)

func main() {
	numPlayers := 5

	state := NewStatefulUI()
	logger := NewLogUI()
	ui := NewMultiUI(state, NewPlayerUI(0, logger))
	g := game.NewGame(numPlayers, ui)

	for !state.gameOver {
		for i := 0; i < numPlayers; i++ {
			if i == 0 {
				g.PickCard(i, maxStrategy(i, state))
			} else {
				g.PickCard(i, randStrategy(i, state))
			}
		}
	}

	fmt.Printf("Score is %v", state.score)
}

func randStrategy(player int, ui *statefulUI) core.Card {
	available := ui.availableCards[player]
	return available[rand.Intn(len(available))]
}

func maxStrategy(player int, ui *statefulUI) core.Card {
	var bestCard = ui.availableCards[player][0]
	var bestScore core.Score

	for _, available := range ui.availableCards[player] {
		scoreKeeper := core.NewGame()
		round := core.NewRound()

		for p, chosen := range ui.chosenCards {
			hand := core.NewHand()
			for _, card := range chosen {
				hand.AddCard(card)
			}
			if p == player {
				hand.AddCard(available)
			}
			round.AddHand(hand)
		}

		scores := core.ScoreGame(scoreKeeper)
		if len(scores) == 0 {
			continue
		}
		score := scores[player]
		if score > bestScore {
			bestCard = available
			bestScore = score
		}
	}

	return bestCard
}
