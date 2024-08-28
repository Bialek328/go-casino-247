package session

import (
    "github.com/Bialek328/go-casino-247/pkg/blackjack"
)

type Session struct {
    Dealer blackjack.Dealer
    Players []blackjack.Player
    Deck blackjack.Deck
}
