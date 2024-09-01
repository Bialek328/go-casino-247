package blackjack

const (

)
type Player struct {
	Name      string
	Cash      int
	Bet       int
	Hand      []Card
	HandScore int
}

type Dealer struct {
	Name      string
	Hand      []Card
	HandScore int
}

func NewPlayer(name string) *Player {
    player := &Player{}
    player.Name = name
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

// todo: consider renaming or changing logic
func (p *Player) GetHandScore() int {
    score := calculateHandScore(p.Hand)
    p.HandScore = score
    return score
}

func (p *Player) GetCard(card Card) {
	p.Hand = append(p.Hand, card)
}

func (d *Dealer) GetHandScore() int {
    score := calculateHandScore(d.Hand)
    d.HandScore = score
    return score
}

func (d *Dealer) GetCard(card Card) {
	d.Hand = append(d.Hand, card)
}
