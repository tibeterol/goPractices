package main

import (
	"math/rand"
)

func main() {
	/*	var numbers []int

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

		}*/

	println(getPattern(8))

}

func getPattern(patternLength int) string {
	pattern := ""
	closedOnes := 0
	openedOnes := 0

	if patternLength%2 != 0 {
		return "Girilen uzunluk mutlaka çift sayı olmalıdır"
	}

	for i := 0; i < patternLength; i++ {
		randomNum := rand.Intn(3)

		if openedOnes < patternLength/2 && (randomNum%2 == 0 || closedOnes == openedOnes) {
			pattern += "("
			openedOnes += 1
		} else {
			pattern += ")"
			closedOnes += 1
		}
	}

	return pattern
}
