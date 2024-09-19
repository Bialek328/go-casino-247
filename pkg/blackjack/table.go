package blackjack

import (

	"github.com/google/uuid"
)

type Table struct {
    ID string `json:"id"`
    Dealer *Player `json:"dealer"`
    Deck *Deck `json:"-"`
}

func NewTable() *Table {
    table := &Table{
        ID: uuid.New().String(),
        Dealer: NewPlayer("Dealer", true),
        Deck: NewDeck(6),
    }
    return table
}

func (t *Table) ClearTable() {
   t.Deck = NewDeck(6)
   t.Dealer = NewPlayer("Dealer", true)
}
