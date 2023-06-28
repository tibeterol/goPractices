package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {

	bts, err := ioutil.ReadFile(fileName)

	if err != nil {
		println("Error : ", err)
		os.Exit(1) // programi derhal kapatmasi icin 0 dan farkli bir deger girilmesi gerekiyor
	}

	stringArr := strings.Split(string(bts), ",")

	return deck(stringArr)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}

}
