package main

import "fmt"

var dizi2 = []string{
	"afgdg",
}

func main() {

	fmt.Println(dizi2)

	arraydizisi := [2]int{ // array
		1, 3,
	}
	dizi := []string{ // slice
		"kart1", newCard(),
	}

	dizi = append(dizi, "kart 3")

	for i, kart := range dizi {
		fmt.Println(i, kart)
	}

	for i := 1; i <= 10; i++ { // 1 den 10 a kadar iterate
		fmt.Println(i)
	}

	for i, sayi := range arraydizisi {
		fmt.Println(i, sayi)
	}

}

func newCard() string {
	return "yeniKart"
}
