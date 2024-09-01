package blackjack

import (
	"fmt"

	"github.com/google/uuid"
)

type Table struct {
    ID string
    Dealer *Dealer
    Players []*Player
    Deck *Deck
}

func NewTable() *Table {
    table := &Table{
        ID: uuid.New().String(),
        Dealer: NewDealer(),
        Players: make([]*Player, 0, 8),
        Deck: NewDeck(6),
    }
    return table
}

func (t *Table) AddPlayer(p *Player) error {
    if len(t.Players) == cap(t.Players) {
        err := fmt.Errorf("Cannot add player, table is full")
        return err
    } else {
        t.Players = append(t.Players, p)
    }
    return nil
}

func (t *Table) RemovePlayer(id string) error {
    for i, val := range(t.Players) {
        if val.ID == id {
            t.Players = append(t.Players[:i], t.Players[i+1:]...)
            return nil
        }
    }
    err := fmt.Errorf("Cannot remove player. No such player at the table")
    return err
}

func (t *Table) ClearTable() {

}
