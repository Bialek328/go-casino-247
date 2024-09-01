package blackjack

import (
    "fmt"
    "testing"
    "reflect"
)

func TestNewPlayer(t *testing.T) {
    name := "Mat"
    player := NewPlayer(name)
    if reflect.TypeOf(player) != reflect.TypeOf(Player{}) {
        t.Errorf("NewPlayer returning wrong type")
    }
    if player.Name != name {
        fmt.Println(player.Name)
        t.Errorf("Wrong name assigned to player")
    }

}

func TestNewDealer(t *testing.T) {
    dealer := NewDealer()
    if reflect.TypeOf(dealer) != reflect.TypeOf(Dealer{}) {
        t.Errorf("Wrong type of the dealer")
    }
    if dealer.Name != "Dealer" {
        t.Errorf("Something went wrong with the name")
    }
}

func TestPlayerGetCard(t *testing.T) {
    player := NewPlayer("Mat")
    card := Card {Symbol: Queen, Color: "clubs", Value: 10}
    player.GetCard(card)
    if len(player.Hand) == 0 {
        t.Errorf("No card added to player hand")
    }
    playerCard := player.Hand[0]
    if playerCard.Color != "clubs" {
        t.Errorf("Different color than expected")
    }
    if playerCard.Value != 10 {
        t.Errorf("Different value than expected")
    }
    if playerCard.Symbol != Queen {
        t.Errorf("Different symbol than expected")
    }
}

func TestDealerGetCard(t *testing.T) {
    dealer := NewDealer()
    card := Card{Symbol: Jack, Color: "hearts", Value: 10}
    dealer.GetCard(card)
    dealerCard := dealer.Hand[0]
    if dealerCard.Color != "hearts" {
        t.Errorf("Different color than expected")
    }
    if dealerCard.Value != 10 {
        t.Errorf("Different value than expected")
    }
    if dealerCard.Symbol != Jack {
        t.Errorf("Different symbol than expected")
    }
}

func TestGetHandScore(t *testing.T) {
    player := NewPlayer("Mat")
    player.GetHandScore()
    score := player.HandScore
    if score != 0 {
        t.Errorf("Expected empty hand to have score 0")
    }
    cards :=[]Card{
        {Symbol: Queen, Color: "clubs", Value: 10},
        {Symbol: "2", Color: "clubs", Value: 2},
    }
    player.Hand = cards
    player.GetHandScore()
    if player.HandScore != 12 {
        t.Errorf("Error calculating hand value")
    }
    ace := Card{Symbol: Ace, Color: "spades", Value: 11}
    player.GetCard(ace)
    player.GetHandScore()
    if player.HandScore != 13 {
        t.Errorf("Error while adjusting Ace value")
    }
    player.Hand = []Card{
        {Symbol: Ace, Color: "spades",Value: 11},
        {Symbol: "10",Color: "clubs",Value: 10},
    }
    player.GetHandScore()
    if player.HandScore != 21 {
        t.Errorf("Error while getting base ace value")
    }
}
