package main

import (
	"fmt"

	"github.com/Bialek328/go-casino-247/internal/boards"
	"github.com/Bialek328/go-casino-247/internal/decks"
	"github.com/Bialek328/go-casino-247/internal/players"
)

func runBanner() {
	for {
		boards.DisplayBanner()
	}
}

func getInput() string {
	fmt.Println("Press Q to exit")
	var q string
	fmt.Scanln(&q)
	return q
}

func initialDealing(players []*players.Player, dealer *players.Dealer, deck *decks.BlackJackDeck) {
	for i := 0; i < 2; i++ {
		for _, player := range players {
			hit(player, deck)
		}
		card := deck.DealCard()
		dealer.GetCard(card)
	}
}

func hit(player *players.Player, deck *decks.BlackJackDeck) {
	card := deck.DealCard()
	player.GetCard(card)
}

func main() {
	deck := decks.InitBlackJackDeck()
	deck_pointer := &deck
	player1 := players.Player{
		Name: "Mat",
		Cash: 1000,
	}
	dealer := players.Dealer{
		Name: "Dealer",
	}
	dealer_pointer := &dealer
	deck.Shuffle()
	players := []*players.Player{&player1}

	initialDealing(players, dealer_pointer, deck_pointer)
	dealer.GetHandScore()
	hit(players[0], deck_pointer)
	player1.GetHandScore()

	fmt.Println(player1)
	fmt.Println(dealer)
}
