package poker

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// HandRank represents the rank of a poker hand
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

// Card represents a playing card with a rank and suit
type Card struct {
	Rank int
	Suit rune
}

// Hand represents a poker hand with a list of cards and its rank
type Hand struct {
	Cards       []Card // The cards in the hand
	cardsSorted []Card
	HandRank    HandRank // The rank of the hand (e.g., Full House, Flush)
}

// NewHand creates a new Hand from a string and determines its rank
func NewHand(handStr string) (Hand, error) {
	cards := strings.Fields(handStr)
	if len(cards) != 5 {
		return Hand{}, errors.New("invalid hand")
	}

	handCards := make([]Card, 5)
	for i, cardStr := range cards {
		rank, suit, err := parseCard(cardStr)
		if err != nil {
			return Hand{}, err
		}
		handCards[i] = Card{Rank: rank, Suit: suit}
	}

	hand := Hand{Cards: handCards}
	hand.sortHand()
	hand.HandRank = hand.determineHandRank()
	return hand, nil
}

// sortHand sorts the cards in descending order and fills the CardsSorted field
func (h *Hand) sortHand() {
	// Make a copy of the cards and sort them
	sortedCards := make([]Card, len(h.Cards))
	copy(sortedCards, h.Cards)
	sort.Slice(sortedCards, func(i, j int) bool {
		return sortedCards[i].Rank > sortedCards[j].Rank
	})

	// Assign the sorted cards to the CardsSorted field
	h.cardsSorted = sortedCards
}

// BestHand takes a list of poker hands in string format, evaluates them, and returns the best hand(s)
func BestHand(hands []string) ([]string, error) {
	if len(hands) == 0 {
		return nil, errors.New("no hands provided")
	}

	var bestHands []Hand
	var bestHandStrs []string

	// Evaluate each hand and find the best hands
	for _, handStr := range hands {
		hand, err := NewHand(handStr)
		if err != nil {
			return nil, err
		}

		if len(bestHands) == 0 || compareHands(hand, bestHands[0]) {
			// If bestHands is empty or the current hand is better
			// than the best hand so far:
			// Update bestHands to only include the current hand.
			// Update bestHandStrs to only include the string representation
			// of the current hand.
			bestHands = []Hand{hand}
			bestHandStrs = []string{handStr}
		} else if !compareHands(bestHands[0], hand) {
			// If the current hand is equal to the best hand so far:
			// Append the current hand to bestHands.
			// Append the string representation of the current hand
			// to bestHandStrs.
			bestHands = append(bestHands, hand)
			bestHandStrs = append(bestHandStrs, handStr)
		}
	}

	return bestHandStrs, nil
}

// extractRank extracts and validates the rank from a card string
func extractRank(cardStr string) (int, error) {
	// Clean up rankStr to remove any invalid characters
	reg := regexp.MustCompile(`[^0-9A-Za-z]`)
	cleanRankStr := reg.ReplaceAllString(cardStr, "")

	// Map of card ranks
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

	// Determine rank from the cleaned rank string
	rank, exists := rankMap[cleanRankStr]
	if !exists {
		return 0, errors.New("invalid card rank")
	}

	return rank, nil
}

// extractSuit extracts and validates the suit from a card string
func extractSuit(cardStr string) (rune, error) {
	// Valid suits
	validSuits := map[byte]bool{
		164: true,
		162: true,
		161: true,
		167: true,
		100: true,
		97:  true,
		98:  true,
		103: true}

	// Count the number of suit runes in the card string
	suitCount := countSuits(cardStr, validSuits)
	if suitCount != 1 {
		return 0, errors.New("suitCount != 1")
	}

	// Extract the suit (last character)
	suit, err := getSuit(cardStr)
	if err != nil {
		return 0, err
	}

	// Validate suit
	if !validSuits[byte(suit)] {
		return 0, errors.New("!validSuits[byte(suit)]")
	}

	return suit, nil
}

func getSuit(cardStr string) (rune, error) {
	suitStr := cardStr[len(cardStr)-1:]
	if len(suitStr) != 1 {
		return 0, errors.New("len(suitStr) != 1")
	}
	return rune(suitStr[0]), nil
}

// countSuits counts the number of suit runes in a card string
func countSuits(cardStr string, validSuits map[byte]bool) int {
	suitCount := 0
	for _, char := range cardStr {
		if validSuits[byte(char)] {
			suitCount++
		}
	}
	return suitCount
}

// parseCard converts a card string to a Card struct with rank and suit
func parseCard(cardStr string) (int, rune, error) {
	// Validate card length
	if len(cardStr) < 2 {
		return 0, 0, errors.New("invalid card length")
	}

	rank, err := extractRank(cardStr)
	if err != nil {
		return 0, 0, err
	}

	suit, err := extractSuit(cardStr)
	if err != nil {
		return 0, 0, err
	}

	return rank, suit, nil
}

// determineHandRank determines the rank of the hand
func (h Hand) determineHandRank() HandRank {
	if h.isStraightFlush() {
		return StraightFlush
	} else if h.isFourOfAKind() {
		return FourOfAKind
	} else if h.isFullHouse() {
		return FullHouse
	} else if h.isFlush() {
		return Flush
	} else if h.isStraight() {
		return Straight
	} else if h.isThreeOfAKind() {
		return ThreeOfAKind
	} else if h.isTwoPair() {
		return TwoPair
	} else if h.isOnePair() {
		return OnePair
	} else {
		return HighCard
	}
}

// isStraightFlush checks if the hand forms a straight flush
func (h Hand) isStraightFlush() bool {
	return h.isStraight() && h.isFlush()
}

// isFourOfAKind checks if the hand forms four of a kind
func (h Hand) isFourOfAKind() bool {
	counts := h.rankCounts()
	for _, count := range counts {
		if count == 4 {
			return true
		}
	}
	return false
}

// isFullHouse checks if the hand forms a full house
func (h Hand) isFullHouse() bool {
	counts := h.rankCounts()

	hasThree := false
	hasTwo := false

	for _, count := range counts {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

// isFlush checks if the hand forms a flush
func (h Hand) isFlush() bool {
	suit := h.Cards[0].Suit
	for _, card := range h.Cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

// isStraight checks if the hand forms a straight using cardsSorted
func (h Hand) isStraight() bool {
	// Ensure the cards are sorted
	h.sortHand()

	ranks := make([]int, len(h.cardsSorted))
	for i, card := range h.cardsSorted {
		ranks[i] = card.Rank
	}

	// Check for regular straight
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]-1 {
			// Check for A-2-3-4-5 straight
			if ranks[0] == 14 && ranks[1] == 5 && ranks[2] == 4 && ranks[3] == 3 && ranks[4] == 2 {
				return true
			}
			return false
		}
	}
	return true
}

// isThreeOfAKind checks if the hand forms three of a kind
func (h Hand) isThreeOfAKind() bool {
	counts := h.rankCounts()
	for _, count := range counts {
		if count == 3 {
			return true
		}
	}
	return false
}

// isTwoPair checks if the hand forms two pairs
func (h Hand) isTwoPair() bool {
	counts := h.rankCounts()
	pairCount := 0
	for _, count := range counts {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

// isOnePair checks if the hand forms one pair
func (h Hand) isOnePair() bool {
	counts := h.rankCounts()
	for _, count := range counts {
		if count == 2 {
			return true
		}
	}
	return false
}

// rankCounts returns a map of card ranks to their respective counts in the hand
func (h Hand) rankCounts() map[int]int {
	counts := make(map[int]int)
	for _, card := range h.cardsSorted {
		counts[card.Rank]++
	}
	return counts
}

// compareHands compares two hands by rank and returns true if the first hand is higher
func compareHands(h1, h2 Hand) bool {
	if h1.HandRank != h2.HandRank {
		return h1.HandRank > h2.HandRank
	}
	result := compareCards(h1, h2)
	return result > 0
}

// compareCards compares two sets of cards and returns an integer indicating their relative rank
func compareCards(h1, h2 Hand) int {
	if h1.HandRank == FourOfAKind && h2.HandRank == FourOfAKind {
		return compareFourOfAKinds(h1, h2)
	} else if h1.HandRank == FullHouse && h2.HandRank == FullHouse {
		return compareFullHouses(h1, h2)
	} else if h1.HandRank == Straight && h2.HandRank == Straight {
		return compareStraights(h1, h2)
	} else if h1.HandRank == ThreeOfAKind && h2.HandRank == ThreeOfAKind {
		return compareThreeOfAKinds(h1, h2)
	} else if h1.HandRank == TwoPair && h2.HandRank == TwoPair {
		return compareTwoPairs(h1, h2)
	} else if h1.HandRank == OnePair && h2.HandRank == OnePair {
		return comparePairs(h1, h2)
	}

	for i := 0; i < len(h1.cardsSorted); i++ {
		rank1 := h1.cardsSorted[i].Rank
		rank2 := h2.cardsSorted[i].Rank

		if rank1 != rank2 {
			return rank1 - rank2
		}
	}
	return 0
}

// isAceLowStraight checks if the hand forms an A-2-3-4-5 straight
func (h Hand) isAceLowStraight() bool {
	ranks := make([]int, len(h.Cards))
	for i, card := range h.Cards {
		ranks[i] = card.Rank
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))
	return ranks[0] == 14 && ranks[1] == 5 && ranks[2] == 4 && ranks[3] == 3 && ranks[4] == 2
}

// compareFullHouses compares two full houses and returns an integer indicating which is better
func compareFullHouses(h1, h2 Hand) int {
	// Get the rank counts for both hands
	counts1 := h1.rankCounts()
	counts2 := h2.rankCounts()

	// Find the rank of the three of a kind and the pair for each hand
	var threeOfAKind1, pair1, threeOfAKind2, pair2 int
	for rank, count := range counts1 {
		if count == 3 {
			threeOfAKind1 = rank
		} else if count == 2 {
			pair1 = rank
		}
	}
	for rank, count := range counts2 {
		if count == 3 {
			threeOfAKind2 = rank
		} else if count == 2 {
			pair2 = rank
		}
	}

	// Compare the three of a kind ranks
	if threeOfAKind1 != threeOfAKind2 {
		return threeOfAKind1 - threeOfAKind2
	}

	// If the three of a kind ranks are equal, compare the pair ranks
	return pair1 - pair2
}

// compareStraights compares two straights and returns an integer indicating which is better
func compareStraights(h1, h2 Hand) int {
	for i := 0; i < len(h1.cardsSorted); i++ {
		rank1 := h1.cardsSorted[i].Rank
		rank2 := h2.cardsSorted[i].Rank

		// Adjust ranks if the hand forms an A-2-3-4-5 straight
		if rank1 == 14 && h1.isAceLowStraight() {
			rank1 = 1
		}
		if rank2 == 14 && h2.isAceLowStraight() {
			rank2 = 1
		}

		if rank1 != rank2 {
			return rank1 - rank2
		}
	}

	return 0
}

// compareFourOfAKinds compares two Four of a Kind hands and returns an integer indicating which is better
func compareFourOfAKinds(h1, h2 Hand) int {
	// Get the rank counts for both hands
	counts1 := h1.rankCounts()
	counts2 := h2.rankCounts()

	var fourOfAKind1, kicker1, fourOfAKind2, kicker2 int

	// Extract four of a kind and kicker for the first hand
	for rank, count := range counts1 {
		if count == 4 {
			// Identify the rank of the four of a kind
			fourOfAKind1 = rank
		} else if count == 1 {
			// Identify the kicker
			kicker1 = rank
		}
	}

	// Extract four of a kind and kicker for the second hand
	for rank, count := range counts2 {
		if count == 4 {
			// Identify the rank of the four of a kind
			fourOfAKind2 = rank
		} else if count == 1 {
			// Identify the kicker
			kicker2 = rank
		}
	}

	// Compare the four of a kinds
	if fourOfAKind1 != fourOfAKind2 {
		return fourOfAKind1 - fourOfAKind2
	}

	// Compare the kickers
	return kicker1 - kicker2
}

// compareThreeOfAKinds compares two Three of a Kind hands and returns an integer indicating which is better
func compareThreeOfAKinds(h1, h2 Hand) int {
	// Get the rank counts for both hands
	counts1 := h1.rankCounts()
	counts2 := h2.rankCounts()

	var threeOfAKind1, kicker1, kicker2, threeOfAKind2, kicker3 int

	// Extract three of a kind and kickers for the first hand
	for rank, count := range counts1 {
		if count == 3 {
			threeOfAKind1 = rank
		} else if count == 1 {
			kicker1 = rank
		} else {
			kicker2 = rank
		}
	}

	// Extract three of a kind and kickers for the second hand
	for rank, count := range counts2 {
		if count == 3 {
			threeOfAKind2 = rank
		} else if count == 1 {
			kicker3 = rank
		} else {
			kicker2 = rank
		}
	}

	// Compare the three of a kinds
	if threeOfAKind1 != threeOfAKind2 {
		return threeOfAKind1 - threeOfAKind2
	}

	// Compare the first kickers
	if kicker1 != kicker2 {
		return kicker1 - kicker2
	}

	// Compare the second kickers
	return kicker1 - kicker3
}

// extractPairsAndKickers extracts pairs and kickers from sorted cards
func extractPairsAndKickers(cardsSorted []Card) ([]int, []int) {
	pairs := []int{}
	kickers := []int{}

	i := 0
	for i < len(cardsSorted) {
		rank := cardsSorted[i].Rank
		count := 1
		if i+1 < len(cardsSorted) && cardsSorted[i+1].Rank == rank {
			count = 2
			i++ // Skip the next card as it's part of the pair
		}
		if count == 2 {
			pairs = append(pairs, rank)
		} else {
			kickers = append(kickers, rank)
		}
		i++
	}

	return pairs, kickers
}

// comparePairs compares two One Pair hands and returns an
// integer indicating which is better
func comparePairs(h1, h2 Hand) int {
	pairs1, kickers1 := extractPairsAndKickers(h1.cardsSorted)
	pairs2, kickers2 := extractPairsAndKickers(h2.cardsSorted)

	// Compare the pairs
	if pairs1[0] != pairs2[0] {
		return pairs1[0] - pairs2[0]
	}

	// Compare the kickers
	for i := 0; i < len(kickers1); i++ {
		if kickers1[i] != kickers2[i] {
			return kickers1[i] - kickers2[i]
		}
	}

	return 0
}

// compareTwoPairs compares two Two Pairs hands and returns an
// integer indicating which is better
func compareTwoPairs(h1, h2 Hand) int {
	pairs1, kickers1 := extractPairsAndKickers(h1.cardsSorted)
	pairs2, kickers2 := extractPairsAndKickers(h2.cardsSorted)

	// Compare the highest pairs
	if pairs1[0] != pairs2[0] {
		return pairs1[0] - pairs2[0]
	}

	// Compare the second highest pairs
	if pairs1[1] != pairs2[1] {
		return pairs1[1] - pairs2[1]
	}

	// Finally, compare the kickers
	for i := 0; i < len(kickers1); i++ {
		if kickers1[i] != kickers2[i] {
			return kickers1[i] - kickers2[i]
		}
	}

	return 0
}
