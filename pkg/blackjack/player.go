package blackjack

import (
    "github.com/google/uuid"
)

type Hand struct {
    Cards []Card `json:"cards"`
}

type Player struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Cash      int `json:"cash"`
    BetAmount       int `json:"bet"`
    Hand `json:"hand"`       
    HandScore int `json:"handScore"`
    IsDealer bool `json:"isDealer"`
}

func NewPlayer(name string, isDealer bool) *Player {
    player := &Player{
        ID: uuid.New().String(),
        Name: name,
    }
    return player
}

func calculateHandScore(hand []Card) int {
	// calculate the players score, adjust ace value depending on the overall score
	var score int
	var aceCount int
	for _, card := range hand{
		score += card.Value
		if card.Symbol == Ace {
			aceCount++
		}
	}
	for score > 21 && aceCount > 0 {
		score -= 10
		aceCount--
	}
    return score
}

func (p *Player) GetHandScore() int {
    score := calculateHandScore(p.Cards)
    p.HandScore = score
    return score
}

func (p *Player) GetCard(card Card) {
	p.Cards = append(p.Cards, card)
}

func (p *Player) ClearHand() {
    p.Cards = []Card{}
}

func (p* Player) Bet(amount int) {
    p.BetAmount= amount
    p.Cash = p.Cash - amount
}

func (p *Player) Stand() {

}

func (h *Hand) ClearHand() {
    h.Cards = []Card{}
}
