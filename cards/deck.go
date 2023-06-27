package main

import (
	"io/ioutil"
	"strings"
)

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

}

func (d deck) saveToDisk(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) // 0666 kodu ile herkes bu dosyayi yazabilir ve okuyabilir bu yetkiyi vermis olduk
	// writefile kullanimdan kaldirilmis ileride bu fonk. komple kaldirilabilir ya da duzenlenebilir demek bu. Kullanilmasi go tarafindan onerilmiyor. Alternatif olarak os kutuphanesinin os.Create i kullanilabilir
}
