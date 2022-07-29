package main

import (
	"flag"
	"fmt"

	game "github.com/o-ifeanyi/gophercises/go-blackjack/pkg/model"
)

var numOfRounds int

func init() {
	flag.IntVar(&numOfRounds, "rounds", 3, "The number of game rounds before program ends")
	flag.Parse()
}

func main() {
	var gs game.GameState
	gs = gs.Shuffle()

	for i := 1; i <= numOfRounds; i++ {
		fmt.Printf("\n\nROUND  %d\n\n", i)
		gs = gs.Deal()
		var input string

		for gs.State == game.StatePlayerTurn {
			fmt.Println("Player:", gs.Player.String())
			fmt.Println("Dealer:", gs.Dealer.DealerString())
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs.Player, gs.Deck = append(gs.Player, gs.Deck[0]), gs.Deck[1:]
				if gs.Player.Score() > 21 {
					gs.State = game.StateDealerTurn
				}
			case "s":
				gs.State = game.StateDealerTurn
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		for gs.State == game.StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs.Dealer, gs.Deck = append(gs.Dealer, gs.Deck[0]), gs.Deck[1:]
			} else {
				gs.State = game.StateHandOver
			}
		}

		gs = gs.EndGame()
	}
}
