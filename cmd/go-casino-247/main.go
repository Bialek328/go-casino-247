package main

import (
	// "fmt"
	"fmt"
	"os"

	"github.com/Bialek328/go-casino-247/internal/boards"
	// "github.com/Bialek328/go-casino-247/internal/decks"
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

func main() {
	go runBanner()
	q := getInput()
	if q == "q" || q == "Q" {
		os.Exit(1)
	}
}
