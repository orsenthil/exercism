package camicia

import "strings"

type Outcome struct {
	finishes bool
	cards    int
	tricks   int
}

func SimulateGame(playerA, playerB []string) Outcome {
	seen := map[string]bool{}
	pile := []string{}

	var currentPlayer string

	var currentState string

	currentPlayer = "A"
	var cards int = 0
	var penaltyDue int
	var tricks int
	var card string

	for {

		// 1. Save the current state of the game, check if we have seen this before and loop is detected.
		currentState = strings.Join(fingerprint(playerA), "") + "|" + strings.Join(fingerprint(playerB), "") + currentPlayer

		if penaltyDue == 0 && seen[currentState] {
			return Outcome{
				finishes: false,
				cards:    cards,
				tricks:   tricks,
			}
		}

		if penaltyDue == 0 {
			seen[currentState] = true
		}

		// 2. Check if the game is finished.

		// 3. Play the cards and handle the penalities.
		switch currentPlayer {
		case "A":
			if len(playerA) == 0 {
				playerB, pile = collectPile(playerB, pile)
				tricks++
				return Outcome{
					finishes: true,
					cards:    cards,
					tricks:   tricks,
				}
			}
			card, playerA, pile = playCard(playerA, pile)
			cards++
			if penaltyValue(card) > 0 {
				penaltyDue = penaltyValue(card)
				currentPlayer = "B"
			} else {
				if penaltyDue > 0 {
					penaltyDue--
					if penaltyDue == 0 {
						playerB, pile = collectPile(playerB, pile)
						tricks++
						currentPlayer = "B"
						if len(playerA) == 0 {
							return Outcome{
								finishes: true,
								cards:    cards,
								tricks:   tricks,
							}
						}
					}
				} else if penaltyDue == 0 {
					currentPlayer = "B"
				}
			}

		case "B":
			if len(playerB) == 0 {
				playerA, pile = collectPile(playerA, pile)
				tricks++
				return Outcome{
					finishes: true,
					cards:    cards,
					tricks:   tricks,
				}
			}
			card, playerB, pile = playCard(playerB, pile)
			cards++
			if penaltyValue(card) > 0 {
				penaltyDue = penaltyValue(card)
				currentPlayer = "A"
			} else {
				if penaltyDue > 0 {
					penaltyDue--
					if penaltyDue == 0 {
						playerA, pile = collectPile(playerA, pile)
						tricks++
						currentPlayer = "A"
						if len(playerB) == 0 {
							return Outcome{
								finishes: true,
								cards:    cards,
								tricks:   tricks,
							}
						}
					}
				} else if penaltyDue == 0 {
					currentPlayer = "A"
				}
			}

		}

	}

}

func playCard(deck, pile []string) (card string, newDeck []string, newPile []string) {
	card = deck[0]
	newDeck = deck[1:]
	newPile = append(pile, card)
	return
}

func collectPile(deck, pile []string) (newDeck, newPile []string) {
	newDeck = append(deck, pile...)
	newPile = []string{}
	return
}

func fingerprint(cards []string) []string {
	result := []string{}
	for _, card := range cards {
		if card == "J" || card == "Q" || card == "K" || card == "A" {
			result = append(result, card)
		} else {
			result = append(result, "N")
		}
	}
	return result
}

func penaltyValue(card string) int {
	switch card {
	case "J":
		return 1
	case "Q":
		return 2
	case "K":
		return 3
	case "A":
		return 4
	}
	return 0
}
