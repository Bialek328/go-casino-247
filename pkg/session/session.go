package session

import (
    "fmt"
    "github.com/Bialek328/go-casino-247/pkg/blackjack"
    "github.com/google/uuid"
)

type Session struct {
    ID string `json:"id"`
    Table *blackjack.Table `json:"table"`
    Players []*blackjack.Player `json:"players"`
    Turn int `json:"turn"` // index of players whose turn it is
}

var Sessions = map[string]*Session{}

func NewSession() *Session {
    session := &Session{
        ID: uuid.NewString(),
        Table: blackjack.NewTable(),
        Players: []*blackjack.Player{},
        Turn: 0,
    }
    return session
}

func (s *Session) InitialDeal() {
    for i := 0; i < 2; i++ {
        for _, v := range(s.Players) {
            card := s.Table.Deck.DealCard()
            v.GetCard(card)
        }
        card := s.Table.Deck.DealCard()
        s.Table.Dealer.GetCard(card)
    }
}

func (s *Session) AddPlayer(p *blackjack.Player) error {
    if len(s.Players) == cap(s.Players) {
        err := fmt.Errorf("Cannot add player, table is full")
        return err
    } else {
        s.Players = append(s.Players, p)
    }
    return nil
}

func (s *Session) RemovePlayer(id string) error {
    for i, val := range(s.Players) {
        if val.ID == id {
            s.Players = append(s.Players[:i], s.Players[i+1:]...)
            return nil
        }
    }
    err := fmt.Errorf("Cannot remove player. No such player at the table")
    return err
}
