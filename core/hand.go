package core

type Score int

func scoreIndividualHand(hand *Hand) Score {
	var totalScore Score = 0

	totalScore += scoreSashimi(hand)
	totalScore += scoreTempura(hand)
	totalScore += scoreDumplings(hand)
	totalScore += scoreNigiri(hand)

	return totalScore
}

func scoreNigiri(hand *Hand) Score {
	var nigiriScore Score
	var wasabiCount int
	for _, card := range hand.nigiriAndWasabi {
		var cardValue Score
		switch card {
		case EggNigiri:
			cardValue = 1
		case SalmonNigiri:
			cardValue = 2
		case SquidNigiri:
			cardValue = 3
		}
		if card != Wasabi && wasabiCount > 0 {
			cardValue *= 3
			wasabiCount--
		}
		if card == Wasabi {
			wasabiCount++
		}
		nigiriScore += cardValue
	}
	return nigiriScore
}

func scoreSashimi(hand *Hand) Score {
	return Score(10 * (hand.numSashimi / 3))
}

func scoreTempura(hand *Hand) Score {
	return Score(5 * (hand.numTempura / 2))
}

func scoreDumplings(hand *Hand) Score {
	switch {
	case hand.numDumplings == 1:
		return 1
	case hand.numDumplings == 2:
		return 3
	case hand.numDumplings == 3:
		return 6
	case hand.numDumplings == 4:
		return 10
	case hand.numDumplings >= 5:
		return 15
	default:
		return 0
	}
}

type Card int

//go:generate stringer -type Card

const (
	Chopsticks   Card = iota
	Tempura
	Sashimi
	Dumplings
	SquidNigiri   // 3 points
	SalmonNigiri  // 2 points
	EggNigiri     // 1 point
	Wasabi
	MakiRoll1
	MakiRoll2
	MakiRoll3
	Pudding
)

type Hand struct {
	numSashimi      int
	numPuddings     int
	numTempura      int
	numDumplings    int
	numMaki         int
	numChopsticks   int
	nigiriAndWasabi []Card
	all             []Card
}

func (h *Hand) AddCard(card Card) {
	switch card {
	case Sashimi:
		h.numSashimi++
	case Tempura:
		h.numTempura++
	case Dumplings:
		h.numDumplings++
	case MakiRoll1:
		h.numMaki ++
	case MakiRoll2:
		h.numMaki += 2
	case MakiRoll3:
		h.numMaki += 3
	case Pudding:
		h.numPuddings++
	case Chopsticks:
		h.numChopsticks++
	case EggNigiri, SalmonNigiri, SquidNigiri, Wasabi:
		h.nigiriAndWasabi = append(h.nigiriAndWasabi, card)
	}
	h.all = append(h.all, card)
}

func (h *Hand) Materialize() []Card {
	return h.all
}

func add(out []Card, card Card, repeat int) []Card {
	for i := 0; i < repeat; i++ {
		out = append(out, card)
	}
	return out
}

func (h *Hand) HasChopsticks() bool {
	return h.numChopsticks > 0
}

func (h *Hand) UseChopsticks() {
	h.numChopsticks--
	for i := range h.all {
		if h.all[i] == Chopsticks {
			h.all = append(h.all[:i], h.all[i+1:]...)
			return
		}
	}
}

func NewHand() *Hand {
	return &Hand{}
}
