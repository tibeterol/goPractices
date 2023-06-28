package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Wick",
		contact: contactInfo{
			email:   "blabla@gmail.com",
			zipCode: 34516,
		},
	}

	//fmt.Printf("%+v", john)

	john.updateName("Johny")

	// slicelar mapler channellar falan hepsi reference type yani fnk.a parametre olarak verilse dahi pointersiz update yapilabiliyor cunku mesela slice icin bir slice olusturuldugunda arkada 1 array ve slice data type olusuyor ve slice arrayi point ediyor basindaki pointer ile pointer , capacity ve  lengthden olusuyo slice.
	// int,string,struct falan bunlar value type ve bunlari guncellerken pointer kullanimi gerekli

	john.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerOfPerson *person) updateName(newFirstName string) {
	(*pointerOfPerson).firstName = newFirstName
}
