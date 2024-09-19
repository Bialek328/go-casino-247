package blackjack

import (
    "strconv"
    "math/rand"
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
    Color string `json:"color"`
    Symbol string `json:"symbol"`
    Value int `json:"value"`
}

type Deck struct {
    Cards []Card `json:"cards"`
}

func (d *Deck) Shuffle() {
	cards := d.Cards
	deck_len := len(cards) - 1
	for i := 0; i <= deck_len; i++ {
		j := rand.Intn(deck_len)
		cards[i], cards[j] = cards[j], cards[i]
	}
}

func (d *Deck) DealCard() Card {
    lastIndex := len(d.Cards) - 1
    card := d.Cards[lastIndex]
    d.Cards = d.Cards[:lastIndex]
    return card
}

func NewDeck(numberOfDecks int) *Deck {
    // numberOfDecks - deck being a singular colection of 52 cards
    // return already shuffled deck
    var all_cards []Card
    for i := 0; i < numberOfDecks; i++ {
        cards := createSingleDeckCards()
        all_cards = append(all_cards, cards...)
    }
    deck := &Deck{Cards: all_cards}
    deck.Shuffle()
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

