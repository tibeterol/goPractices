package main

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // 2 tane deck tipinden return bekliyor
	return d[:handSize], d[handSize:] // mesela 4 eleman var diyelim ve handSize 2 geldi bu 0. elemandan 2 dahil olmayacak sekilde al demek ilk return icin, 2.return icin 2 dahil hepsini al
}
