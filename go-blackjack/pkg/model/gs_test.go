package game

import (
	"testing"

	"github.com/o-ifeanyi/gophercises/go-card-deck"
)

var gs GameState

func TestShuffle(t *testing.T) {
	gs = gs.Shuffle()
	if len(gs.Deck) != 156 {
		t.Error("Expected 3 deck of cards with length of 156, Got", len(gs.Deck))
	}
	if len(gs.Player) != 0 || len(gs.Dealer) != 0 {
		t.Error("Expected Player and Dealer to be empty, Got", len(gs.Player), len(gs.Dealer))
	}
	if gs.State != StatePlayerTurn {
		t.Error("Expected State to be StatePlayerTurn, Got", gs.State)
	}
}

func TestDeal(t *testing.T) {
	if len(gs.Deck) != 156 {
		t.Error("Expected 3 deck of cards with length of 156, Got", len(gs.Deck))
	}
	if len(gs.Player) != 0 || len(gs.Dealer) != 0 {
		t.Error("Expected Player and Dealer to be empty, Got", len(gs.Player), len(gs.Dealer))
	}

	gs = gs.Deal()

	if len(gs.Deck) != 152 {
		t.Error("Expected cards with length of 152, Got", len(gs.Deck))
	}
	if len(gs.Player) != 2 || len(gs.Dealer) != 2 {
		t.Error("Expected Player and Dealer to have 2 cards each Got", len(gs.Player), len(gs.Dealer))
	}
}

func TestClone(t *testing.T) {
	if len(gs.Deck) != 152 {
		t.Error("Expected cards with length of 152, Got", len(gs.Deck))
	}
	if len(gs.Player) != 2 || len(gs.Dealer) != 2 {
		t.Error("Expected Player and Dealer to have 2 cards each Got", len(gs.Player), len(gs.Dealer))
	}

	gs = gs.Clone()

	if len(gs.Deck) != 152 {
		t.Error("Expected cards with length of 152, Got", len(gs.Deck))
	}
	if len(gs.Player) != 2 || len(gs.Dealer) != 2 {
		t.Error("Expected Player and Dealer to have 2 cards each Got", len(gs.Player), len(gs.Dealer))
	}
}

func ExamplePlayerWin() {
	pWin, dLoose := Hand{
		{Rank: card.Ace, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Diamonds}, // 21
	},
		Hand{
			{Rank: card.King, Suit: card.Diamonds}, // 10
		}
	gs.Player, gs.Dealer = pWin, dLoose

	gs.EndGame()

	// Output:
	// Player: Ace of Diamonds (1), King of Diamonds (10)
	// Dealer: King of Diamonds (10), *** HIDDEN ***
	//
	// ==== You Win! ====
	//
	// Player: 21
	// Dealer: 10
}

func ExamplePlayerLoose() {
	pLoose, dWin := Hand{
		{Rank: card.Two, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Diamonds}, // 12
	},
		Hand{
			{Rank: card.King, Suit: card.Diamonds},
			{Rank: card.King, Suit: card.Diamonds}, // 20
		}
	gs.Player, gs.Dealer = pLoose, dWin

	gs.EndGame()

	// Output:
	// Player: Two of Diamonds (2), King of Diamonds (10)
	// Dealer: King of Diamonds (10), *** HIDDEN ***
	//
	// ==== You Loose! ====
	//
	// Player: 12
	// Dealer: 20
}

func ExamplePlayerBust() {
	p, d := Hand{
		{Rank: card.King, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Diamonds}, // 30
	},
		Hand{
			{Rank: card.King, Suit: card.Diamonds},
			{Rank: card.King, Suit: card.Diamonds}, // 20
		}
	gs.Player, gs.Dealer = p, d

	gs.EndGame()

	// Output:
	// Player: King of Diamonds (10), King of Diamonds (10), King of Diamonds (10)
	// Dealer: King of Diamonds (10), *** HIDDEN ***
	//
	// ==== You Busted! ====
	//
	// Player: 30
	// Dealer: 20
}

func ExampleDealerBust() {
	p, d := Hand{
		{Rank: card.King, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Diamonds}, // 20
	},
		Hand{
			{Rank: card.King, Suit: card.Diamonds},
			{Rank: card.King, Suit: card.Diamonds},
			{Rank: card.King, Suit: card.Diamonds}, // 30
		}
	gs.Player, gs.Dealer = p, d

	gs.EndGame()

	// Output:
	// Player: King of Diamonds (10), King of Diamonds (10)
	// Dealer: King of Diamonds (10), *** HIDDEN ***
	//
	// ==== Dealer Busted! ====
	//
	// Player: 20
	// Dealer: 30
}

func ExampleDraw() {
	p, d := Hand{
		{Rank: card.King, Suit: card.Diamonds},
		{Rank: card.King, Suit: card.Diamonds}, // 20
	},
		Hand{
			{Rank: card.King, Suit: card.Diamonds},
			{Rank: card.King, Suit: card.Diamonds}, // 20
		}
	gs.Player, gs.Dealer = p, d

	gs.EndGame()

	// Output:
	// Player: King of Diamonds (10), King of Diamonds (10)
	// Dealer: King of Diamonds (10), *** HIDDEN ***
	//
	// ==== Draw! ====
	//
	// Player: 20
	// Dealer: 20
}
