package game

import (
	"fmt"

	deck "github.com/o-ifeanyi/gophercises/go-card-deck"
)

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs GameState) Shuffle() GameState {
	ret := gs.Clone()
	ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	return ret
}

func (gs GameState) Deal() GameState {
	ret := gs.Clone()
	var player, dealer Hand

	for i := 0; i < 2; i++ {
		player, ret.Deck = append(player, ret.Deck[0]), ret.Deck[1:]
		dealer, ret.Deck = append(dealer, ret.Deck[0]), ret.Deck[1:]
	}
	ret.Player = player
	ret.Dealer = dealer
	ret.State = StatePlayerTurn
	return ret
}

func (gs GameState) Clone() GameState {
	newDeck := make([]deck.Card, len(gs.Deck))
	player := make(Hand, len(gs.Player))
	dealer := make(Hand, len(gs.Dealer))

	copy(newDeck, gs.Deck)
	copy(player, gs.Player)
	copy(dealer, gs.Dealer)
	return GameState{
		Deck:   newDeck,
		State:  gs.State,
		Player: player,
		Dealer: dealer,
	}
}

func (gs GameState) EndGame() GameState {
	fmt.Println("Player:", gs.Player.String())
	fmt.Printf("Dealer: %s\n\n", gs.Dealer.DealerString())
	ret := gs.Clone()
	pScore := ret.Player.Score()
	dScore := ret.Dealer.Score()
	switch {
	case pScore > 21:
		fmt.Println("==== You Busted! ====")
	case dScore > 21:
		fmt.Println("==== Dealer Busted! ====")
	case pScore > dScore:
		fmt.Println("==== You Win! ====")
	case dScore > pScore:
		fmt.Println("==== You Loose! ====")
	case pScore == dScore:
		fmt.Println("==== Draw! ====")
	}

	fmt.Printf("\nPlayer: %d\n", pScore)
	fmt.Printf("Dealer: %d\n\n", dScore)
	return ret
}
