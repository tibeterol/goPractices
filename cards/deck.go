package main

type deck []string

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}
