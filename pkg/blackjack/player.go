package blackjack

import (
    "github.com/google/uuid"
)

type Hand struct {
    Cards []Card
}

type Player struct {
    ID        string
    Name      string
    Cash      int
    Bet       int
    Hand      
    HandScore int
}

type Dealer struct {
	Name      string
	Hand      
	HandScore int
}

func NewPlayer(name string) *Player {
    player := &Player{
        ID: uuid.New().String(),
        Name: name,
    }
    return player
}

func NewDealer() *Dealer {
    dealer := &Dealer{}
    dealer.Name = "Dealer"
    return dealer
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

func (d *Dealer) GetHandScore() int {
    score := calculateHandScore(d.Cards)
    d.HandScore = score
    return score
}

func (d *Dealer) GetCard(card Card) {
	d.Cards = append(d.Cards, card)
}

func (d *Dealer) ClearHand() {
    d.Cards = []Card{}
}   

func (h *Hand) ClearHand() {
    h.Cards = []Card{}
}
