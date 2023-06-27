package main

func main() {

	cards := newDeck()

	println(cards.toString())

	cards.saveToDisk("my_cards")

}
