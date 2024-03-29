package blackjack

import (
	"fmt"

	"github.com/euller88/deck"
)

// Player is an abstraction of what a player in a blackjack game may be, a human inputing commands or an AI executing tasks
type Player interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Summary(hand [][]deck.Card, dealer []deck.Card)
}

// HumanPlayer returns a Player "object" ready to receive human inputs
func HumanPlayer() Player {
	return &humanPlayer{}
}

// TestPlayer returns a Player "object" to test some behaviours
func TestPlayer() Player {
	return &dealer{}
}

// humanPlayer is the implementation of the Player interface to receive interactions from a real person, whatever it may be
type humanPlayer struct{}

// Bet returns whatever the player whats to put at stake in the game
func (pl humanPlayer) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled.")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

// Play is the core function of that takes an input and decides what kind of move a player should make
func (pl humanPlayer) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand, (d)ouble, s(p)lit")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

// Summary gives the final state of a player
func (pl humanPlayer) Summary(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer:", dealer)
	fmt.Println("")
}

// Dealer is the implementation of the Player interface that represents a dealer in a blackjack game
type dealer struct{}

// Bet is the implementation of the Bet method of the Player interface, this actually does nothing
func (dl dealer) Bet(shuffled bool) int {
	// noop
	return 0
}

// Play is the implementation of the Play method of the Player interface
func (dl dealer) Play(hand []deck.Card, dealer deck.Card) Move {
	ds := Score(hand...)
	if ds <= 16 || (ds == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

// Summary is the implementation of the Summary method of the Player interface, this actually does nothing
func (dl dealer) Summary(hand [][]deck.Card, dealer []deck.Card) {
	// noop
}
