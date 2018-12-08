package core

type Round struct {
	hands []*Hand
}

func (r *Round) AddHand(h *Hand) {
	r.hands = append(r.hands, h)
}

func NewRound() *Round {
	return &Round{}
}

func scoreMaki(round *Round, scores []Score) {
	type placement struct {
		numMaki int
		index int
		numTies int
	}

	var (
		first = placement{0, -1, 0}
		second = placement{0, -1, 0}
	)
	for i, hand := range round.hands {
		if hand.numMaki > first.numMaki {
			second = first
			first = placement{numMaki: hand.numMaki, index: i}
		} else if hand.numMaki > 0 && second.index == -1 {
			second = placement{numMaki: hand.numMaki, index: i}
		}

		if hand.numMaki == first.numMaki {
			first.numTies++
		} else if hand.numMaki == second.numMaki {
			second.numTies++
		}
	}

	if first.numMaki == 0 {
		return
	}

	for i := range scores {
		if round.hands[i].numMaki == first.numMaki {
			scores[i] += Score(6 / first.numTies)
		}
		if second.numMaki > 0 && first.numTies == 1 && round.hands[i].numMaki == second.numMaki {
			scores[i] += Score(3 / second.numTies)
		}
	}
}
