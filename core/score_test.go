package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScoreHandIndividually(t *testing.T) {
	type testCase struct {
		cards         []Card
		expectedScore Score
		name          string
	}
	var testCases = []testCase{
		{
			name:          "Empty Hand Scores 0",
			cards:         []Card{},
			expectedScore: 0,
		},
		{
			name:          "Chopsticks Scores 0",
			cards:         []Card{Chopsticks},
			expectedScore: 0,
		},
		{
			name:          "1 Sashimi Scores 0",
			cards:         []Card{Sashimi},
			expectedScore: 0,
		},
		{
			name:          "2 Sashimi Scores 0",
			cards:         []Card{Sashimi, Sashimi},
			expectedScore: 0,
		},
		{
			name:          "3 Sashimi Scores 10",
			cards:         []Card{Sashimi, Sashimi, Sashimi},
			expectedScore: 10,
		},
		{
			name:          "6 Sashimi Scores 20",
			cards:         []Card{Sashimi, Sashimi, Sashimi, Sashimi, Sashimi, Sashimi},
			expectedScore: 20,
		},
		{
			name:          "5 Sashimi Scores 10",
			cards:         []Card{Sashimi, Sashimi, Sashimi, Sashimi, Sashimi},
			expectedScore: 10,
		},
		{
			name:          "1 Tempura Scores 0",
			cards:         []Card{Tempura},
			expectedScore: 0,
		},
		{
			name:          "2 Tempura Scores 5",
			cards:         []Card{Tempura, Tempura},
			expectedScore: 5,
		},

		{
			name:          "3 Tempura Scores 5",
			cards:         []Card{Tempura, Tempura, Tempura},
			expectedScore: 5,
		},
		{
			name:          "4 Tempura Scores 10",
			cards:         []Card{Tempura, Tempura, Tempura, Tempura},
			expectedScore: 10,
		},
		{
			name:          "1 Dumpling Scores 1",
			cards:         []Card{Dumplings},
			expectedScore: 1,
		},
		{
			name:          "2 Dumplings Scores 3",
			cards:         []Card{Dumplings, Dumplings},
			expectedScore: 3,
		},
		{
			name:          "3 Dumplings Scores 6",
			cards:         []Card{Dumplings, Dumplings, Dumplings},
			expectedScore: 6,
		},
		{
			name:          "4 Dumplings Scores 10",
			cards:         []Card{Dumplings, Dumplings, Dumplings, Dumplings},
			expectedScore: 10,
		},
		{
			name:          "5 Dumplings Scores 15",
			cards:         []Card{Dumplings, Dumplings, Dumplings, Dumplings, Dumplings},
			expectedScore: 15,
		},
		{
			name:          "6 Dumplings Scores 15",
			cards:         []Card{Dumplings, Dumplings, Dumplings, Dumplings, Dumplings, Dumplings},
			expectedScore: 15,
		},
		{
			name:          "Egg Nigiri Scores 1",
			cards:         []Card{EggNigiri},
			expectedScore: 1,
		},
		{
			name:          "2 Egg Nigiri Scores 2",
			cards:         []Card{EggNigiri, EggNigiri},
			expectedScore: 2,
		},
		{
			name:          "Salmon Nigiri Scores 2",
			cards:         []Card{SalmonNigiri},
			expectedScore: 2,
		},
		{
			name:          "2 Salmon Nigiri Scores 4",
			cards:         []Card{SalmonNigiri, SalmonNigiri},
			expectedScore: 4,
		},
		{
			name:          "Squid Nigiri Scores 3",
			cards:         []Card{SquidNigiri},
			expectedScore: 3,
		},
		{
			name:          "2 Squid Nigiri Scores 6",
			cards:         []Card{SquidNigiri, SquidNigiri},
			expectedScore: 6,
		},
		{
			name:          "Wasabi Egg Nigiri Scores 3",
			cards:         []Card{Wasabi, EggNigiri},
			expectedScore: 3,
		},
		{
			name:          "Wasabi  2 Egg Nigiri Scores 4",
			cards:         []Card{Wasabi, EggNigiri, EggNigiri},
			expectedScore: 4,
		},
		{
			name:          "Wasabi Salmon Nigiri Scores 6",
			cards:         []Card{Wasabi, SalmonNigiri},
			expectedScore: 6,
		},
		{
			name:          "Wasabi 2 Salmon Nigiri Scores 8",
			cards:         []Card{Wasabi, SalmonNigiri, SalmonNigiri},
			expectedScore: 8,
		},
		{
			name:          "Wasabi Squid Nigiri Scores 9",
			cards:         []Card{Wasabi, SquidNigiri},
			expectedScore: 9,
		},
		{
			name:          "Wasabi 2 Squid Nigiri Scores 12",
			cards:         []Card{Wasabi, SquidNigiri, SquidNigiri},
			expectedScore: 12,
		},
		{
			name:          "2 Wasabi 2 Squid Nigiri Scores 12",
			cards:         []Card{Wasabi, Wasabi, SquidNigiri, SquidNigiri},
			expectedScore: 18,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var hand = NewHand()
			for _, card := range tc.cards {
				hand.AddCard(card)
			}
			assert.Equal(t, tc.expectedScore, scoreIndividualHand(hand), tc.name)
		})
	}
}

func TestScoreRound(t *testing.T) {
	tests := []struct {
		name           string
		hands          [][]Card
		expectedScores []Score
	}{
		{
			name: "Most Maki Rolls Scores +6, Second gets +3 (1)",
			hands: [][]Card{
				{MakiRoll1},
				{MakiRoll2},
			},
			expectedScores: []Score{3, 6},
		},
		{
			name: "Most Maki Rolls Scores +6, Second gets +3 (2)",
			hands: [][]Card{
				{MakiRoll1, MakiRoll1, MakiRoll1},
				{MakiRoll2},
			},
			expectedScores: []Score{6, 3},
		},
		{
			name: "Most Maki Rolls Scores +6, Second gets +3 (2)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
			},
			expectedScores: []Score{3, 6},
		},
		{
			name: "Most Maki Rolls Scores +6, Second gets +3 (3)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll1, MakiRoll1},
			},
			expectedScores: []Score{3, 6, 0},
		},
		{
			name: "Second Highest Maki Splits 3 (1)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
			},
			expectedScores: []Score{6, 1, 1},
		},
		{
			name: "Second Highest Maki Splits 3 (2)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
			},
			expectedScores: []Score{6, 1, 1, 1},
		},
		{
			name: "Second Highest Maki Splits 3 (3)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
			},
			expectedScores: []Score{6, 0, 0, 0, 0},
		},
		{
			name: "Highest Maki Splits 6, No second highest (1)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
			},
			expectedScores: []Score{3, 3, 0, 0},
		},
		{
			name: "Highest Maki Splits 6, No second highest (2)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
			},
			expectedScores: []Score{2, 2, 2, 0, 0},
		},
		{
			name: "Highest Maki Splits 6, No second highest (3)",
			hands: [][]Card{
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll3, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
				{MakiRoll3, MakiRoll2, MakiRoll1},
			},
			expectedScores: []Score{1, 1, 1, 1, 0, 0},
		},
		{
			name:           "No Maki Scores 0 (1)",
			hands:          [][]Card{{}, {}, {}, {}, {}, {}},
			expectedScores: []Score{0, 0, 0, 0, 0, 0},
		},
		{
			name: "No Maki Scores 0",
			hands: [][]Card{
				{},
				{},
				{},
				{MakiRoll3},
			},
			expectedScores: []Score{0, 0, 0, 6},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var round = NewRound()
			for _, cards := range test.hands {
				var hand = NewHand()
				for _, card := range cards {
					hand.AddCard(card)
				}
				round.AddHand(hand)
			}
			var scores = make([]Score, len(test.hands))
			scoreMaki(round, scores)
			assert.Equal(t, test.expectedScores, scores)
		})
	}
}

func TestScoreGame(t *testing.T) {
	tests := []struct {
		name           string
		game           [3][][]Card
		expectedScores []Score
	}{
		{
			name: "Most Puddings Gets +6, Least -6 (1)",
			game: [3][][]Card{
				{
					{Pudding},
					{},
					{},
				},
				{
					{Pudding},
					{Pudding},
					{},
				},
				{
					{},
					{},
					{},
				},
			},
			expectedScores: []Score{6, 0, -6},
		},
		{
			name: "Most Puddings Gets +6, Least -6 (2)",
			game: [3][][]Card{
				{
					{Pudding},
					{Pudding},
					{},
					{},
				},
				{
					{Pudding},
					{Pudding},
					{},
					{},
				},
				{
					{},
					{},
					{},
					{},
				},
			},
			expectedScores: []Score{3, 3, -3, -3},
		},
		{
			name: "Complete Scoring",
			game: [3][][]Card{
				{
					{Pudding},
					{Pudding},
					{Dumplings},
					{Chopsticks, MakiRoll3},
				},
				{
					{Pudding},
					{Pudding, MakiRoll2},
					{Tempura, Tempura},
					{},
				},
				{
					{},
					{},
					{},
					{Sashimi, Sashimi, Sashimi},
				},
			},
			expectedScores: []Score{3, 9, 3, 13},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var game = NewGame()
			for _, roundHands := range test.game {
				var round = NewRound()
				for _, cards := range roundHands {
					var hand = NewHand()
					for _, card := range cards {
						hand.AddCard(card)
					}
					round.AddHand(hand)
				}
				game.AddRound(round)
			}
			assert.Equal(t, test.expectedScores, ScoreGame(game))
		})
	}
}
