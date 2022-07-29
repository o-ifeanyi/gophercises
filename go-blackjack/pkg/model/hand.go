package game

import (
	"fmt"
	"strings"

	deck "github.com/o-ifeanyi/gophercises/go-card-deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	hands := make([]string, len(h))
	for i, v := range h {
		hands[i] = fmt.Sprintf("%s (%d)", v.String(), min(int(v.Rank), 10))
	}
	return strings.Join(hands, ", ")
}

func (h Hand) DealerString() string {
	return fmt.Sprintf("%s (%d), *** HIDDEN ***", h[0].String(), min(int(h[0].Rank), 10))
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, v := range h {
		score += min(int(v.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
