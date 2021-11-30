package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int
    switch card {
        case "ace":
    		value = 11
        case "king","queen","jack", "ten":
    		value = 10
    	case "two":
    		value = 2
        case "three":
    		value = 3
        case "four":
    		value = 4
        case "five":
    		value = 5
        case "six":
    		value = 6
        case "seven":
    		value = 7
        case "eight":
    		value = 8
        case "nine":
    		value = 9
        default:
    		value = 0
    }
	return value
	panic("Please implement the ParseCard function")
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
  	isblackJack := ParseCard(card1) + ParseCard(card2)
    switch {
        case isblackJack == 21:
    		return true
    	default:
    		return false
    }
	panic("Please implement the IsBlackjack function")
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string{
	if isBlackjack == true {
        switch {
            case dealerScore >= 10:
        		return "S"
            case dealerScore < 10:
        		return "W"
        }
    } else {
    	return "P"
    }
	panic("Please implement the LargeHand function")
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
    if handScore >= 12 && handScore <= 16 {
    	switch {
            case dealerScore >= 7:
        		return "H"
            case dealerScore < 7:
        		return "S"
        }
    } else if handScore	>= 17 {
    	return "S"
    } else {
    	return "H"
    }
	panic("Please implement the SmallHand function")
}

