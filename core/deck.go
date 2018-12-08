package core

import (
	"math/rand"
	"fmt"
)

var cardCounts = map[Card]int {
	Tempura: 14,
	Sashimi: 14,
	Dumplings: 14,
	MakiRoll1: 6,
	MakiRoll2: 12,
	MakiRoll3: 8,
	EggNigiri: 5,
	SalmonNigiri: 10,
	SquidNigiri: 5,
	Pudding: 10,
	Wasabi: 6,
	Chopsticks: 4,
}

type Deck struct {
	used int
	cards []Card
}

func (deck *Deck) Draw(numCards int) []Card {
	result := deck.cards[deck.used: deck.used+numCards]
	deck.used += numCards
	return result
}

func NewDeck() *Deck {
	var cards = make([]Card, 0, 108)
	for card, count := range cardCounts {
		for i := 0; i < count; i++ {
			cards = append(cards, card)
		}
	}
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	fmt.Println(cards)
	return &Deck{0, cards}
}
