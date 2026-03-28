package poker

import (
	"sort"
	"strings"
)

type card struct {
	rank  string
	suit  rune
	value int
}

type hand struct {
	cards    []card
	original string
}

func parseCard(s string) card {
	rankMap := map[string]int{
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
		"A":  14,
	}
	var rank string
	var suit rune

	if len(s) == 3 {
		rank = s[:2]
		suit = rune(s[2])
	} else {
		rank = s[:1]
		suit = rune(s[1])
	}

	return card{rank: rank, suit: suit, value: rankMap[rank]}
}

func parseHand(s string) hand {
	cards := strings.Fields(s)
	h := hand{cards: make([]card, len(cards)), original: s}
	for i, cardStr := range cards {
		h.cards[i] = parseCard(cardStr)
	}
	return h
}

type Category int

const (
	HighCard Category = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type handCategory struct {
	hand     hand
	category Category
	ranks    []int
}

func categorizeHand(h hand) (Category, []int) {
	if isStraightFlush(h) {
		return StraightFlush, []int{getHighCard(h)}
	}
	if isFourOfAKind(h) {
		return FourOfAKind, getFourOfAKindRanks(h)
	}
	if isFullHouse(h) {
		return FullHouse, getFullHouseRanks(h)
	}
	if isFlush(h) {
		return Flush, getFlushRanks(h)
	}
	if isStraight(h) {
		return Straight, []int{getHighCard(h)}
	}
	if isThreeOfAKind(h) {
		return ThreeOfAKind, getThreeOfAKindRanks(h)
	}
	if isTwoPairs(h) {
		return TwoPairs, getTwoPairsRanks(h)
	}
	if isOnePair(h) {
		return OnePair, getOnePairRanks(h)
	}
	return HighCard, []int{getHighCard(h)}
}

func getOnePairRanks(h hand) []int {
	ranks := getRanks(h)
	for i := 0; i < len(ranks)-1; i++ {
		if ranks[i] == ranks[i+1] {
			return []int{ranks[i], ranks[4], ranks[3], ranks[2]}
		}
	}
	return nil
}

func isOnePair(h hand) bool {
	ranks := getRanks(h)
	for i := 0; i < len(ranks)-1; i++ {
		if ranks[i] == ranks[i+1] {
			return true
		}
	}
	return false

}

func getTwoPairsRanks(h hand) []int {
	ranks := getRanks(h)
	if ranks[0] == ranks[1] && ranks[2] == ranks[3] {
		return []int{ranks[0], ranks[2], ranks[4]}
	}
	if ranks[0] == ranks[1] && ranks[3] == ranks[4] {
		return []int{ranks[0], ranks[3], ranks[2]}
	}
	return []int{ranks[1], ranks[3], ranks[0]}

}

func isTwoPairs(h hand) bool {
	ranks := getRanks(h)
	if ranks[0] == ranks[1] && ranks[2] == ranks[3] {
		return true
	}
	if ranks[0] == ranks[1] && ranks[3] == ranks[4] {
		return true
	}
	if ranks[1] == ranks[2] && ranks[3] == ranks[4] {
		return true
	}
	return false

}

func getThreeOfAKindRanks(h hand) []int {
	ranks := getRanks(h)
	for i := 0; i < len(ranks)-2; i++ {
		if ranks[i] == ranks[i+1] && ranks[i] == ranks[i+2] {
			return []int{ranks[i]}
		}
	}
	return nil

}

func isThreeOfAKind(h hand) bool {
	ranks := getRanks(h)
	for i := 0; i < len(ranks)-2; i++ {
		if ranks[i] == ranks[i+1] && ranks[i] == ranks[i+2] {
			return true
		}
	}
	return false

}

func getFlushRanks(h hand) []int {
	ranks := getRanks(h)
	return ranks
}

func getFullHouseRanks(h hand) []int {
	ranks := getRanks(h)
	if ranks[0] == ranks[1] && ranks[1] == ranks[2] {
		return []int{ranks[0], ranks[3]}
	}
	return []int{ranks[2], ranks[3]}

}

func isFullHouse(h hand) bool {
	ranks := getRanks(h)
	if ranks[0] == ranks[1] && ranks[1] == ranks[2] && ranks[3] == ranks[4] {
		return true
	}
	if ranks[0] == ranks[1] && ranks[2] == ranks[3] && ranks[3] == ranks[4] {
		return true
	}
	return false

}

func getFourOfAKindRanks(h hand) []int {
	ranks := getRanks(h)
	for i := 0; i < len(ranks)-3; i++ {
		if ranks[i] == ranks[i+1] && ranks[i] == ranks[i+2] && ranks[i] == ranks[i+3] {
			return []int{ranks[i]}
		}
	}
	return nil

}

func isFourOfAKind(h hand) bool {
	ranks := getRanks(h)
	for i := 0; i < len(ranks)-3; i++ {
		if ranks[i] == ranks[i+1] && ranks[i] == ranks[i+2] && ranks[i] == ranks[i+3] {
			return true
		}
	}
	return false

}

func getHighCard(h hand) int {
	max := 0
	for _, c := range h.cards {
		if c.value > max {
			max = c.value
		}
	}
	return max

}

func isStraightFlush(h hand) bool {
	return isFlush(h) && isStraight(h)
}

func getRanks(h hand) []int {
	ranks := make([]int, len(h.cards))
	for i, c := range h.cards {
		ranks[i] = c.value
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ranks))) // Sort in descending order
	return ranks
}

func isStraight(h hand) bool {
	ranks := getRanks(h)
	// Check for Ace-low straight (A,2,3,4,5)
	if ranks[0] == 14 && ranks[1] == 5 && ranks[2] == 4 &&
		ranks[3] == 3 && ranks[4] == 2 {
		return true
	}

	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]-1 { // Note: -1 because ranks are descending
			return false
		}
	}
	return true
}

func isFlush(h hand) bool {
	suit := h.cards[0].suit
	for _, c := range h.cards {
		if c.suit != suit {
			return false
		}
	}
	return true
}

func BestHand(hands []string) ([]string, error) {
	if len(hands) == 0 {
		return nil, nil
	}
	categories := make([]handCategory, len(hands))
	for i, h := range hands {
		parsed := parseHand(h)
		cat, ranks := categorizeHand(parsed)
		categories[i] = handCategory{hand: parsed, category: cat, ranks: ranks}
	}

	// Find best category
	bestCat := categories[0]
	for _, cat := range categories[1:] {
		if cat.category > bestCat.category {
			bestCat = cat
		}
	}

	// Find all hands of the best category and compare their ranks
	var bestHands []string
	for _, c := range categories {
		if c.category == bestCat.category {
			if len(bestHands) == 0 {
				bestHands = append(bestHands, c.hand.original)
			} else {
				comparison := compareRanks(c.ranks, categories[0].ranks)
				if comparison > 0 {
					bestHands = []string{c.hand.original}
				} else if comparison == 0 {
					bestHands = append(bestHands, c.hand.original)
				}
			}
		}
	}
	return bestHands, nil
}

func compareRanks(a, b []int) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return 0
}
