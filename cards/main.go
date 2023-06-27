package main

import "fmt"

// 2 tane dizi typei var go da . 1. si array 2.si slice bu kullanilan slice. Slice da istedigim boyutta olabiliyor ve istedigim gibi eleman ekleyip cikarabiliyorum. Arrayde boyut belli olmak zorunda mesela 30 elemanli dizi. Illaki 30 u belirteceksin arrayde

func main() {
	arraydizisi := [2]int{ // array
		1, 3,
	}
	dizi := []string{ // slice
		"kart1", newCard(),
	}

	dizi = append(dizi, "kart 3") // append   yeni  bir dizi doner. Eleman eklemeye yariyor

	for i, kart := range dizi { // i index ve kart elemanin kendisi range dizi ile slice in boyutu kadar donuyor ve surekli yeni index ve elemana degerler ataniyor := ile
		fmt.Println(i, kart)
	}

	for i, sayi := range arraydizisi {
		fmt.Println(i, sayi)
	}

}

func newCard() string {
	return "yeniKart"
}
