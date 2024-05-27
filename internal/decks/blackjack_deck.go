package decks

import "strconv"

var Colors = [4]string{"clubs", "diamonds", "hearts", "spades"}
var Symbols = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

const aceValue = 11
const figureValue = 10

type Card struct {
	Color  string
	Symbol string
	Value  int
}

type BlackJackDeck struct {
	Cards []Card
}

func InitDeck() BlackJackDeck {
	var cards []Card
	var figures = []string{"J", "Q", "K"}
	for _, color := range Colors {
		for _, symbol := range Symbols {
			if stringInArray(symbol, figures) {
				cards = append(cards, Card{Color: color, Symbol: symbol, Value: figureValue})
			} else {
				num, err := strconv.Atoi(symbol)
				if err == nil {
					cards = append(cards, Card{Color: color, Symbol: symbol, Value: num})
				} else if symbol == "A" {
					cards = append(cards, Card{Color: color, Symbol: symbol, Value: aceValue})
				}
			}
		}
	}
	deck := BlackJackDeck{Cards: cards}
	return deck
}

func stringInArray(str string, array []string) bool {
	for _, s := range array {
		if str == s {
			return true
		}
	}
	return false
}
