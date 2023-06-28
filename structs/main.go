package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//john := person{"John","Wick"} // istenilmeyen kullanim
	john := person{firstName: "John", lastName: "Wick"} // dogru kullanim ornegi

	fmt.Println(john)

}
