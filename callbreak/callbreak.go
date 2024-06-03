// start the program
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// firt create a Card object using golang Struct
type Card struct {
	Suit string
	Rank string
}

// now create a Deck of cards using Card struct
// this function should return a deck which is a slice of Cards or Array of cards
func createDeck() []Card {
	var deck []Card
	suits := []string{"Club", "Diamond", "Heart", "Spade"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range suits {
		for _, rank := range ranks {
			// initialize card
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// now Shuffle the deck using a method call shuffleDeck
func shuffleDeck(deck []Card) []Card {
	// first create the seed for random number
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range deck {
		newPos := seed.Intn(len(deck) - 1)
		deck[i], deck[newPos] = deck[newPos], deck[i]
	}
	return deck
}

// now deal the deck among four players equally
func dealDeck(deck []Card) [4][]Card {
	// hands is a 2D array of 4 people with array of cards with each of them
	var hands [4][]Card
	for i, card := range deck {
		hands[i%4] = append(hands[i%4], card)

	}
	return hands
}
func main() {
	fmt.Println("Starting the program for callbreak")
	hands := dealDeck(shuffleDeck(createDeck()))
	for i, hand := range hands {
		fmt.Printf("Player %d's hand: \n", i+1)
		for _, card := range hand {
			fmt.Printf("%s of %s \n", card.Rank, card.Suit)
		}
	}

}
