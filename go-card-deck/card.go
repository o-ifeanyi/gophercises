//go:generate stringer -type=Suit,Rank
package card

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Suit
	Rank
}

// Return a formated String of the Card
func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

// Return an int which represents the absolute value of the Card
func (c Card) absValue() int {
	return int(c.Rank) + int(c.Suit)
}

// Return a bool dependeing on if the Card's Rank is any of the passed in Ranks
func (c Card) rankIsAnyOf(ranks ...Rank) bool {
	for _, r := range ranks {
		if c.Rank == r {
			return true
		}
	}
	return false
}

type Suit uint8
type Rank uint8

var suits = []Suit{Diamonds, Hearts, Spades, Clubs}

const (
	Diamonds Suit = iota
	Hearts
	Spades
	Clubs
	Joker
)

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Creates a new Deck (i.e []Card) of Cards
// New takes a variadic func parameter that would run on the newly created Deck
func New(fns ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, fn := range fns {
		cards = fn(cards)
	}
	return cards
}

// Filters the ranks from the Deck of Cards
func Filter(ranks ...Rank) func(c []Card) []Card {
	return func(cards []Card) []Card {
		var filtered []Card
		for _, c := range cards {
			if !c.rankIsAnyOf(ranks...) {
				filtered = append(filtered, c)
			}
		}
		return filtered
	}

}

// Sorts the Deck of Cards
// In order of Diamonds, Hearts, Spades, Clubs
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sorts the Deck of Cards using a custom less func
// In order of Diamonds, Hearts, Spades, Clubs
func Sort(less func(c []Card) func(i, j int) bool) func(c []Card) []Card {
	return func(c []Card) []Card {
		sort.Slice(c, less(c))
		return c
	}
}

// Adds n number of Jokers to the Deck
func Jokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Rank: Rank(i), Suit: Joker})
		}
		return cards
	}
}

var shuffleRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Shuffles the Deck of Cards
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, v := range perm {
		ret[i] = cards[v]
	}
	return ret
}

// Creates n number of Decks
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

// Determines the sorting order for DefaultSort
func Less(c []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return c[i].absValue() < c[j].absValue()
	}
}
