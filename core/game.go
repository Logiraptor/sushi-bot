package core

import (
	"sort"
)

type Game struct {
	numRounds int
	rounds    [3]*Round
}

func (g *Game) AddRound(round *Round) {
	g.rounds[g.numRounds] = round
	g.numRounds++
}

func (g *Game) Done() bool {
	return g.numRounds == 3
}

func NewGame() *Game {
	return &Game{}
}

func ScoreGame(game *Game) []Score {
	numPlayers := numPlayers(game)

	scores := make([]Score, numPlayers)

	puddings := make([]struct{i int; n int}, numPlayers)

	for _, round := range game.rounds {
		if round == nil {
			continue
		}
		for i, hand := range round.hands {
			scores[i] += scoreIndividualHand(hand)

			puddings[i].i = i
			puddings[i].n += hand.numPuddings
		}

		scoreMaki(round, scores)
	}

	if !game.Done() {
		return scores
	}

	sort.Slice(puddings, func(i, j int) bool {
		return puddings[i].n < puddings[j].n
	})

	highest := puddings[len(puddings) - 1]
	numHighest := 0
	lowest := puddings[0]
	numLowest := 0

	for _, player := range puddings {
		if player.n == highest.n {
			numHighest++
		}
		if player.n == lowest.n {
			numLowest++
		}
	}

	for _, player := range puddings {
		if player.n == highest.n {
			scores[player.i] += Score(6 / numHighest)
		}
		if player.n == lowest.n {
			scores[player.i] -= Score(6 / numLowest)
		}
	}

	return scores
}

func numPlayers(game *Game) int {
	if game.rounds[0] == nil {
		return 0
	}
	return len(game.rounds[0].hands)
}
