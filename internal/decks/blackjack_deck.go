package decks

import (
	"math/rand"
	"strconv"
	"time"
)

// definition of colors and symbols for cards
var colors = [4]string{"clubs", "diamonds", "hearts", "spades"}
var symbols = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

const aceValue = 11
const figureValue = 10

// symbol for face cards
const (
	Jack  = "J"
	Queen = "Q"
	King  = "K"
	Ace   = "A"
)

// structure to hold card information
type Card struct {
	Color  string
	Symbol string
	Value  int
}

// structure to hold a array of cards to represent a card deck, cards values  are added according to blackjack rules
type BlackJackDeck struct {
	Cards []Card
}

func (d *BlackJackDeck) Shuffle() {
	// TODO: add an algorithm to shuffle the deck
}

// initialize a deck for blackjack
func InitBlackJackDeck() BlackJackDeck {
	var cards []Card
	for _, color := range colors {
		for _, symbol := range symbols {
			var value int
			switch symbol {
			case Jack, Queen, King:
				value = figureValue
			case Ace:
				value = aceValue
			default:
				value, _ = strconv.Atoi(symbol)
			}
			cards = append(cards, Card{Color: color, Symbol: symbol, Value: value})
		}
	}
	deck := BlackJackDeck{Cards: cards}
	return deck
}
