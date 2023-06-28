package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, num := range numbers {

		if num%2 == 0 {
			fmt.Printf("%v is even", num)
			fmt.Println("")
		} else {
			fmt.Printf("%v is odd", num)
			fmt.Println("")
		}

	}

}
