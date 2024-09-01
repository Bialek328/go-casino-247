package blackjack

import (
    "github.com/google/uuid"
)

type Table struct {
    ID string
    Dealer *Dealer
    Players []Player
    Deck *Deck
}

func NewTable() *Table {
    table := &Table{
        ID: uuid.New().String(),
        Dealer: NewDealer(),
        Deck: NewDeck(6),
    }
    return table
}

func (t *Table) StartGame() {

}

func (t *Table) AddPlayer() error {
    return nil
}

func (t *Table) RemovePlayer() error {
    return nil
}

func (t *Table) ClearTable() {

}
