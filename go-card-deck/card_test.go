package card

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Hearts, Rank: Three})
	fmt.Println(Card{Suit: Diamonds, Rank: King})
	fmt.Println(Card{Suit: Spades, Rank: Ace})
	fmt.Println(Card{Suit: Clubs, Rank: Jack})
	fmt.Println(Card{Suit: Joker, Rank: Rank(0)})

	// Output:
	// Three of Hearts
	// King of Diamonds
	// Ace of Spades
	// Jack of Clubs
	// Joker
}

func TestAbsValue(t *testing.T) {
	testStruct := []struct {
		Card
		absVal int
	}{
		{Card{Suit: Diamonds, Rank: Ace}, 1},
		{Card{Suit: Diamonds, Rank: Two}, 2},
		{Card{Suit: Diamonds, Rank: Three}, 3},
		{Card{Suit: Hearts, Rank: Ace}, 2},
		{Card{Suit: Hearts, Rank: Two}, 3},
		{Card{Suit: Hearts, Rank: Three}, 4},
		{Card{Suit: Spades, Rank: King}, 15},
		{Card{Suit: Clubs, Rank: King}, 16},
	}

	for _, v := range testStruct {
		if v.absValue() != v.absVal {
			t.Error("Expected", v.absVal, "Got", v.absValue())
		}
	}
}

func TestRankIsAnyOf(t *testing.T) {
	card := Card{Suit: Diamonds, Rank: Ace}

	if card.rankIsAnyOf(Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King) {
		t.Error("Expected Card to be of Rank Ace")
	}
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Expected 52 cards, Got", len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Diamonds}
	if cards[0] != exp {
		t.Error("Expected", exp.String(), "Got", cards[0].String())
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Diamonds}
	if cards[0] != exp {
		t.Error("Expected", exp.String(), "Got", cards[0].String())
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	count := 0
	for _, v := range cards {
		if v.Suit == Joker {
			count++
		}
	}
	if count != 4 {
		t.Error("Expected", 4, "Jokers, Got", count)
	}
}

func TestFilter(t *testing.T) {
	ranks := []Rank{Ace, Two, Three, Four}
	cards := New(Filter(ranks...))
	exp := 52 - (len(ranks) * 4)
	if len(cards) != exp {
		t.Error("Expected", exp, "cards, Got", len(cards))
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()

	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Error("Expected", first, "received", cards[0])
	}
	if cards[1] != second {
		t.Error("Expected", second, "received", cards[1])
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(4))
	exp := 52 * 4
	if len(cards) != exp {
		t.Error("Expected", exp, "cards, Got", len(cards))
	}
}
