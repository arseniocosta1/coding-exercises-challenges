package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {

	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if handScore == 22 {
		return "P"
	}

	if handScore == 21 {
		if dealerScore >= 10 {
			return "S"
		}
		return "W"
	}

	if 17 <= handScore && handScore <= 20 {
		return "S"
	}

	if 12 <= handScore && handScore <= 16 {
		if dealerScore >= 7 {
			return "H"
		}
		return "S"
	}

	return "H"
}
