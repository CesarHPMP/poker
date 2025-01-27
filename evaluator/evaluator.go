package evaluator

import (
	"fmt"
	"sort"
)

type Card struct {
	Rank int    // Rank: 2-14 (11=Jack, 12=Queen, 13=King, 14=Ace)
	Suit string // Suit: "C", "D", "H", "S"
}

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

var HandRankNames = []string{
	"High Card", "One Pair", "Two Pair", "Three of a Kind", "Straight",
	"Flush", "Full House", "Four of a Kind", "Straight Flush", "Royal Flush",
}

// EvaluateBestHand takes 7 cards and determines the best possible hand ranking.
func EvaluateBestHand(cards []Card) (HandRank, error) {
	if len(cards) != 7 {
		return HighCard, fmt.Errorf("exactly 7 cards are required to evaluate the best hand")
	}

	// Generate all 5-card combinations
	combinations := generateCombinations(cards, 5)

	// Track the best hand rank
	bestRank := HighCard

	// Evaluate all combinations
	for _, combination := range combinations {
		rank := evaluateFiveCardHand(combination)
		if rank > bestRank {
			bestRank = rank
		}
	}

	return bestRank, nil
}

// EvaluateHandProbabilities calculates the probabilities of all hand ranks given 7 cards.
func EvaluateHandProbabilities(cards []Card) (map[HandRank]float64, error) {
	if len(cards) != 7 {
		return nil, fmt.Errorf("exactly 7 cards are required to evaluate hand probabilities")
	}

	// Generate all 5-card combinations
	combinations := generateCombinations(cards, 5)

	// Count occurrences of each hand rank
	rankCounts := make(map[HandRank]int)
	totalCombos := len(combinations)

	for _, combination := range combinations {
		rank := evaluateFiveCardHand(combination)
		rankCounts[rank]++
	}

	// Convert counts to probabilities
	probabilities := make(map[HandRank]float64)
	for rank, count := range rankCounts {
		probabilities[rank] = float64(count) / float64(totalCombos)
	}

	return probabilities, nil
}

// evaluateFiveCardHand determines the rank of a given 5-card hand.
func evaluateFiveCardHand(cards []Card) HandRank {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	// Check for Straight Flush or Royal Flush
	if isFlush(cards) && isStraight(cards) {
		if cards[0].Rank == 10 {
			return RoyalFlush
		}
		return StraightFlush
	}

	// Check for Four of a Kind
	if hasNOfAKind(cards, 4) {
		return FourOfAKind
	}

	// Check for Full House
	if hasFullHouse(cards) {
		return FullHouse
	}

	// Check for Flush
	if isFlush(cards) {
		return Flush
	}

	// Check for Straight
	if isStraight(cards) {
		return Straight
	}

	// Check for Three of a Kind
	if hasNOfAKind(cards, 3) {
		return ThreeOfAKind
	}

	// Check for Two Pair
	if hasTwoPair(cards) {
		return TwoPair
	}

	// Check for One Pair
	if hasNOfAKind(cards, 2) {
		return OnePair
	}

	// Default to High Card
	return HighCard
}

// Helper Functions

// generateCombinations generates all k-card combinations from a set of cards.
func generateCombinations(cards []Card, k int) [][]Card {
	n := len(cards)
	var combinations [][]Card

	var helper func(start int, combo []Card)
	helper = func(start int, combo []Card) {
		if len(combo) == k {
			comboCopy := make([]Card, k)
			copy(comboCopy, combo)
			combinations = append(combinations, comboCopy)
			return
		}
		for i := start; i < n; i++ {
			helper(i+1, append(combo, cards[i]))
		}
	}

	helper(0, []Card{})
	return combinations
}

func isFlush(cards []Card) bool {
	suit := cards[0].Suit
	for _, card := range cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isStraight(cards []Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i+1].Rank != cards[i].Rank+1 {
			return false
		}
	}
	return true
}

func hasNOfAKind(cards []Card, n int) bool {
	rankCount := make(map[int]int)
	for _, card := range cards {
		rankCount[card.Rank]++
	}
	for _, count := range rankCount {
		if count == n {
			return true
		}
	}
	return false
}

func hasFullHouse(cards []Card) bool {
	rankCount := make(map[int]int)
	for _, card := range cards {
		rankCount[card.Rank]++
	}
	hasThree := false
	hasTwo := false
	for _, count := range rankCount {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func hasTwoPair(cards []Card) bool {
	rankCount := make(map[int]int)
	for _, card := range cards {
		rankCount[card.Rank]++
	}
	pairCount := 0
	for _, count := range rankCount {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}
