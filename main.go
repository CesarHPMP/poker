package main

import (
	"fmt"
	"poker/evaluator"
)

func main() {
	cards := []evaluator.Card{
		{Rank: 13, Suit: "H"}, // Ace of Hearts
		{Rank: 13, Suit: "H"}, // King of Hearts
		{Rank: 12, Suit: "H"}, // Queen of Hearts
		{Rank: 11, Suit: "H"}, // Jack of Hearts
		{Rank: 10, Suit: "C"}, // 10 of Hearts
		{Rank: 2, Suit: "C"},  // 2 of Clubs
		{Rank: 3, Suit: "D"},  // 3 of Diamonds
	}

	bestRank, _ := evaluator.EvaluateBestHand(cards)
	fmt.Println("Best Hand:", evaluator.HandRankNames[bestRank])

	probabilities, _ := evaluator.EvaluateHandProbabilities(cards)
	fmt.Println("Hand Probabilities:")
	for rank, prob := range probabilities {
		fmt.Printf("%s: %.2f%%\n", evaluator.HandRankNames[rank], prob*100)
	}
}
