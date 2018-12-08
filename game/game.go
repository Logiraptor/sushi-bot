package game

import (
	"github.com/Logiraptor/sushi-bot/core"
	"fmt"
)

var numCards = map[int]int{
	2: 10, 3: 9, 4: 8, 5: 7,
}

type Game struct {
	numPlayers int
	deck       *core.Deck
	players    []*core.Hand

	played map[int]struct{}
	available  [][]core.Card

	score *core.Game
	ui    UI
}

type UI interface {
	PlayerAvailableCardsChanged(player int, hand []core.Card)
	PlayerChosenCardsChanged(player int, hand []core.Card)
	RoundOver()
	GameOver()

	ScoreChanged(player int, score core.Score)
	DeckChanged([]core.Card)
}

func NewGame(numPlayers int, ui UI) *Game {
	deck := core.NewDeck()
	game := &Game{
		numPlayers: numPlayers,
		deck:       deck,
		score:      core.NewGame(),
		ui:         ui,
	}
	game.initRound()
	return game
}

func (g *Game) initRound() {
	cardsPerPlayer := numCards[g.numPlayers]
	g.players = make([]*core.Hand, g.numPlayers)
	g.available = make([][]core.Card, g.numPlayers)
	for i := range g.available {
		g.players[i] = core.NewHand()
		g.ui.PlayerChosenCardsChanged(i, g.players[i].Materialize())
		g.setAvailable(i, g.deck.Draw(cardsPerPlayer))
	}
	g.played = make(map[int]struct{})
}

func (g *Game) PickCard(player int, card core.Card) {
	g.mustMarkPlayed(player)
	g.mustConsume(player, card)

	g.players[player].AddCard(card)
	g.ui.PlayerChosenCardsChanged(player, g.players[player].Materialize())

	g.tryToPassHands()
}

func (g *Game) PlayChopsticks(player int, card1, card2 core.Card) {
	g.mustMarkPlayed(player)
	g.mustConsume(player, card1)
	g.mustConsume(player, card2)
	if !g.players[player].HasChopsticks() {
		panic(fmt.Sprintf("Illegal play: player %d does not have chopsticks", player))
	}

	g.players[player].UseChopsticks()
	g.ui.PlayerChosenCardsChanged(player, g.players[player].Materialize())
	g.players[player].AddCard(card1)
	g.ui.PlayerChosenCardsChanged(player, g.players[player].Materialize())
	g.players[player].AddCard(card2)
	g.ui.PlayerChosenCardsChanged(player, g.players[player].Materialize())

	g.setAvailable(player, append(g.available[player], core.Chopsticks))
	g.tryToPassHands()
}

func (g *Game) tryToPassHands() {
	if len(g.played) != g.numPlayers {
		return
	}

	g.played = make(map[int]struct{})
	first := g.available[0]
	for i := 0; i < len(g.available)-1; i++ {
		g.setAvailable(i, g.available[i+1])
	}
	g.setAvailable(len(g.available)-1, first)

	for player, score := range core.ScoreGame(g.score) {
		g.ui.ScoreChanged(player, score)
	}

	if len(g.available[0]) == 0 {
		g.endRound()
	}
}

func (g *Game) setAvailable(player int, available []core.Card){
	g.available[player] = available
	g.ui.PlayerAvailableCardsChanged(player, available)
}

func (g *Game) endRound() {
	round := core.NewRound()
	for _, h := range g.players {
		round.AddHand(h)
	}
	g.score.AddRound(round)
	g.ui.RoundOver()

	for player, score := range core.ScoreGame(g.score) {
		g.ui.ScoreChanged(player, score)
	}

	if !g.score.Done() {
		g.initRound()
	} else {
		g.ui.GameOver()
	}
}

func (g *Game) mustMarkPlayed(player int) {
	if _, ok := g.played[player]; ok {
		panic(fmt.Sprintf("Illegal play: player %d has already played", player))
	}
	g.played[player] = struct{}{}
}

func (g *Game) mustConsume(player int, card core.Card) {
	for i, c := range g.available[player] {
		if c == card {
			g.setAvailable(player, append(g.available[player][:i], g.available[player][i+1:]...))
			return
		}
	}
	panic(fmt.Sprintf("Illegal play: no card %v available", card))
}
