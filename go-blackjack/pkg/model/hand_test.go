package game

import (
	"testing"

	"github.com/o-ifeanyi/gophercises/go-card-deck"
)

func TestString(t *testing.T) {
	hand := Hand{
		{Rank: card.Ace, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Spades}, // 21
	}
	s := hand.String()
	if s != "Ace of Diamonds (1), King of Spades (10)" {
		t.Error("Expected Ace of Diamonds (1), King of Spades (10), Got", s)
	}
}

func TestDealerString(t *testing.T) {
	hand := Hand{
		{Rank: card.Ace, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Spades}, // 21
	}
	s := hand.DealerString()
	if s != "Ace of Diamonds (1), *** HIDDEN ***" {
		t.Error("Expected Ace of Diamonds (1), *** HIDDEN ***, Got", s)
	}
}

func TestScore(t *testing.T) {
	hand := Hand{
		{Rank: card.Ace, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Spades}, // 21
	}
	s := hand.Score()
	if s != 21 {
		t.Error("Expected 21, Got", s)
	}
}

func TestMinScore(t *testing.T) {
	hand := Hand{
		{Rank: card.Ace, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Spades}, // 11
	}
	s := hand.MinScore()
	if s != 11 {
		t.Error("Expected 11, Got", s)
	}
}

func TestMin(t *testing.T) {
	m := min(20, 18)
	if m != 18 {
		t.Error("Expected 18, Got", m)
	}
}
