package main

import "fmt"

func main() {

	colors := map[string]string{ // map tanimlama alternatif 3 ve ayni zamanda degerler direk ataniyor
		"red":    "#fff42",
		"yellow": "#fff423",
		"white":  "#546546",
	}

	printMap(colors)

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for ", key, " is: ", value)
	}
}
