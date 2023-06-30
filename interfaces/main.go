package main

import "fmt"

type shape interface {
	calculateArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {

	t := triangle{
		height: 7,
		base:   3,
	}

	s := square{
		sideLength: 10,
	}

	printArea(s)
	printArea(t)
}

func (t triangle) calculateArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) calculateArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) {
	fmt.Println(sh.calculateArea())
}
