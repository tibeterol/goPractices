package main

func main() {

	cards := deck{
		"kart 1", "kart 2",
	}

	cards = append(cards, "kart 3")

	cards.print()

}
