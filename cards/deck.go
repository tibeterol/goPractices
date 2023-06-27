package main

import "strings"

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	//[]string(d) type conversion  deck i string dizisine cevirebiliyoruz bu sekilde. Zaten type deck aslinda string slice i oldugu icin string slicetan kalitim aliyo gibi dusunulebilir rahat bir sekilde tip donusumu yapildi. Sonrasinda kutuphane kullanilarak hepsi tek bir stringe donusturuldu

}
