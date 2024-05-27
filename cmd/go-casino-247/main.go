package main

import (
	"fmt"

	"github.com/Bialek328/go-casino-247/internal/decks"
)

func main() {
	for _, value := range decks.Colors {
		fmt.Println(value)
	}
}
