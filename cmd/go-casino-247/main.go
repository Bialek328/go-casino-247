package main

import (
	"fmt"

	"github.com/Bialek328/go-casino-247/internal/decks"
)

func main() {
	deck := decks.InitDeck()
	for _, card := range deck.Cards {
		fmt.Println(card)
	}
}
