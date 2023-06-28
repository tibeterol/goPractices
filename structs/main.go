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
	p := &john // john structının adresi p pointerina atandi

	p.updateName("Johny")

	john.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerOfPerson *person) updateName(newFirstName string) { // *person person pointer tipi demek
	(*pointerOfPerson).firstName = newFirstName // burada (*pointerOfPerson) ile pointerin isaret ettigi adresin degerine gidiyoruz. Yani structa ve structin nameini guncelliyoruz
}
