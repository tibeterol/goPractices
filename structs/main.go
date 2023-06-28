package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//john := person{firstName: "John", lastName: "Wick"}
	var aziz person

	aziz.firstName = "Aziz"
	aziz.lastName = "Vefa"
	fmt.Printf("%+v", aziz) // %+v ile structin icindeki tum key value lari basar {firstName:Aziz lastName:Vefa}
	//eger deger atamasaydik bos string olacakti. Diger dillerde genelde null ama burada bos string eger intse 0 ataniyor
	//fmt.Println(john)

}
