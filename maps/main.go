package main

import "fmt"

func main() {

	//var colors map[string]string // map tanimlama alternatif 1

	var colors = make(map[string]string) // map tanimlama alternatif 2
	var colors2 = make(map[int]string)
	/*colors := map[string]string{ // map tanimlama alternatif 3 ve ayni zamanda degerler direk ataniyor
		"red":    "#fff42",
		"yellow": "#fff423",
	} */

	colors["white"] = "#45456" // mape eleman ekleme ya da guncelleme

	colors["white"] = "#121212"

	colors["red"] = "#78787"

	colors2[10] = "red"

	delete(colors, "red") // mapten eleman silme
	delete(colors2, 10)

	fmt.Println(colors)
	fmt.Println(colors2)

}
