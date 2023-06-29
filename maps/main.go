package main

import "fmt"

type renk map[string]string

func main() {

	/*colors := map[string]string{ // map tanimlama alternatif 3 ve ayni zamanda degerler direk ataniyor
		"red":    "#fff42",
		"yellow": "#fff423",
		"white":  "#546546",
	}*/

	colors2 := renk{
		"red":    "#fff42",
		"yellow": "#fff423",
		"white":  "#546546",
	}

	//printMap(colors)
	colors2.printRenk()

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for ", key, " is: ", value)
	}
}

func (r renk) printRenk() {
	for key, value := range r {
		fmt.Println("Hex code for ", key, " is: ", value)
	}
}
