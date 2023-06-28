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

	john.updateName("Johny") // john. diyerek kisayol kullanmis olduk fonksiyon receiver olarak bizden pointer beklemesine ragmen person verdim ve dogru calisti. Go arka tarafta bizim icin cevirmeyi yapiyor

	john.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerOfPerson *person) updateName(newFirstName string) {
	(*pointerOfPerson).firstName = newFirstName
}
