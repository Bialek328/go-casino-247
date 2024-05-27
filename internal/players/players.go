package players

import "github.com/Bialek328/go-casino-247/internal/decks"

type Player struct {
	Name      string
	Cash      int
	Bet       int
	Hand      []decks.Card
	HandScore int
}

func (p *Player) GetHandScore() int {
	var score int
	for _, card := range p.Hand {
		score += card.Value
	}
	return score
}

type Dealer struct {
	Name      string
	Hand      []decks.Card
	HandScore int
}

func (d *Dealer) GetHandScore() int {
	var score int
	for _, card := range d.Hand {
		score += card.Value
	}
	return score
}
