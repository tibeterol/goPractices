package main

func main() {

	//cards := newDeck()

	//println(cards.toString())

	readedCards := newDeckFromFile("my_cards")
	readedCards.print()
	//cards.saveToDisk("my_cards")

}
