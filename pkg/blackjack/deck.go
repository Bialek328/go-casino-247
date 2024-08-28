package blackjack

import (
    "strconv"
)

const (
    Jack = "J"
    Queen = "Q"
    King = "K"
    Ace = "A"
)

// definition of colors and symbols for cards
var colors = [4]string{"clubs", "diamonds", "hearts", "spades"}
var symbols = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

const aceValue = 11
const figureValue = 10

type Card struct {
    Color string
    Symbol string
    Value int
}

type Deck struct {
    Cards []Card
}

func (d *Deck) Shuffle() {

}

func (d *Deck) DealCard() {

}

func NewDeck(numberOfDecks int) Deck {
    // numberOfDecks - deck being a singular colection of 52 cards
    // return already shuffled deck
    deck := Deck{}
    var all_cards []Cards
    for _, 
    return deck
}

func createSingleDeckCards() []Card {
    var cards []Card
    for _, color := range(colors) {
        for _, symbol := range(symbols) {
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
    return cards
}
