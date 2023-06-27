package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	hand.print()

	fmt.Println("")

	fmt.Println("Remaning Cards ---")

	remainingCards.print()

}
