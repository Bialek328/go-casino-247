package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Bialek328/go-casino-247/internal/boards"
	"github.com/Bialek328/go-casino-247/internal/decks"
	"github.com/Bialek328/go-casino-247/internal/players"
)

var playersList []players.Player
var Deck decks.BlackJackDeck

func runBanner() {
	boards.DisplayBanner()
}

func getInput() string {
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

func AddPlayer() {
	var name string
	var cash int
	fmt.Println("Player name:")
	fmt.Scanln(&name)
	fmt.Println("Player cash:")
	fmt.Scanln(&cash)

	player := players.Player{Name: name, Cash: cash}
	pointerPlayer := &player
	playersList = append(playersList, *pointerPlayer)
	fmt.Println(playersList)
}

// func StartGame() {
// 	Deck = decks.InitBlackJackDeck()
// 	dealer := players.Dealer{Name: "Dealer"}
// 	// dealerPointer := &dealer
// 	// initialDealing(playersList, dealerPointer, Deck)
// }

func blackjack(input string) {
	const addPlayer = "a"
	const startGame = "s"
	const exit = "q"

	switch input {
	case exit:
		os.Exit(0)
	case addPlayer:
		AddPlayer()
	case startGame:
		fmt.Println("Starting game")
	default:
		err := errors.New("wrong option")
		fmt.Println(err)
	}
}

func main() {
	q := make(chan bool)

	go func() {
		for {
			select {
			case <-q:
				return
			default:
				runBanner()
			}
		}
	}()

	input := getInput()
	if input != "" {
		q <- true
		boards.ClearBanner()
		fmt.Println("Witamy w kasynie")
		fmt.Println("(a) add player\n(s) start game\n(q) exit game")
		fmt.Println("Choose option:")
		var in string

		if in == "q" {
			os.Exit(0)
		}
		for in != "q" {
			blackjack(in)
			fmt.Scanln(&in)
		}
	}
}
