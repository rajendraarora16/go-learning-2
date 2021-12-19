package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	sideLength int
}

type rectangle struct {
	height float64
	width  float64
}

func main() {
	r := rectangle{
		height: 10.45,
		width:  14.20,
	}

	printArea(r)
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func printArea(s shape) {
	fmt.Println(s.area())
}
