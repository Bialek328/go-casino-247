package blackjack

import (
    "fmt"
	"testing"
    "reflect"
)

func TestNewDeck(t *testing.T) {
    singleDecksN := 1
    deck := NewDeck(singleDecksN)
    if reflect.TypeOf(deck) != reflect.TypeOf(Deck{}) {
        t.Errorf("wrong type of deck")
    }
    if len(deck.Cards) != (singleDecksN * 52) {
        t.Errorf("Cards number in deck not matching")
    }
}

func TestDealCard(t *testing.T) {
    deck := NewDeck(1)
    card := deck.DealCard()
    if reflect.TypeOf(card) != reflect.TypeOf(Card{}) {
        t.Errorf("Error dealing card")
    }
    if len(deck.Cards) != 51 {
        t.Errorf("Card not substracted from deck")
    }
}

func TestCreateSingleDeckCards(t *testing.T) {
    cards := createSingleDeckCards()
    if len(cards) != 52 {
        t.Errorf("Cards length is not 52")
    }
}

func TestShuffleDeck(t *testing.T) {
    deck := NewDeck(1)
    firstCard := deck.Cards[0]
    fmt.Println(firstCard)
    deck.Shuffle()
}
