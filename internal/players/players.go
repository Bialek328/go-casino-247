package players

import (
	"github.com/Bialek328/go-casino-247/internal/decks"
)

type Player struct {
	Name      string
	Cash      int
	Bet       int
	Hand      []decks.Card
	HandScore int
}

func (p *Player) GetHandScore() {
	// calculate the players score, adjust ace value depending on the overall score
	var score int
	var aceCount int
	for _, card := range p.Hand {
		score += card.Value
		if card.Symbol == decks.Ace {
			aceCount++
		}
	}
	if score > 21 && aceCount > 0 {
		score -= 10
		aceCount--
	}
	p.HandScore = score
}

func (p *Player) GetCard(card decks.Card) {
	p.Hand = append(p.Hand, card)
}

type Dealer struct {
	Name      string
	Hand      []decks.Card
	HandScore int
}

func (d *Dealer) GetHandScore() {
	var score int
	for _, card := range d.Hand {
		score += card.Value
	}
	d.HandScore = score
}

func (d *Dealer) GetCard(card decks.Card) {
	d.Hand = append(d.Hand, card)
}
