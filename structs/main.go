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
	john.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
