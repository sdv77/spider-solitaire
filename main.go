package main

import (
	"fmt"
	"math/rand"
)

// Main application
func main() {
	fmt.Println(deckBuilder())

}

// Shuffle function
func shuffle(cards []string) []string {
	shuffled := append([]string(nil), cards...)
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}

// Slice of one suit values generator
func generateCardsOneSuit() []string {

	exampleCards := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	return shuffle(exampleCards)

}

// All cards is shuffled by start
func deckBuilder() []string {
	firstDeck := generateCardsOneSuit()
	secondDeck := generateCardsOneSuit()
	thirdDeck := generateCardsOneSuit()
	fourthDeck := generateCardsOneSuit()
	fifthDeck := generateCardsOneSuit()
	sixthDeck := generateCardsOneSuit()
	seventhDeck := generateCardsOneSuit()
	eighthDeck := generateCardsOneSuit()

	finalDeck := []string{}
	finalDeck = append(finalDeck, firstDeck...)
	finalDeck = append(finalDeck, secondDeck...)
	finalDeck = append(finalDeck, thirdDeck...)
	finalDeck = append(finalDeck, fourthDeck...)
	finalDeck = append(finalDeck, fifthDeck...)
	finalDeck = append(finalDeck, sixthDeck...)
	finalDeck = append(finalDeck, seventhDeck...)
	finalDeck = append(finalDeck, eighthDeck...)

	return shuffle(finalDeck)
}

// Main deck
func mainDeck() {

}

/*
 1   2   3   4   5   6   7   8   9  10
=== === === === === === === === === ===
=== === === === === === === === === ===
=== === === === === === === === === ===
=== === === === === === === === === ===
=== === === === === === === === === ===
=== ===  K   8   Q   3   A   2   3   5
4    6

*/
