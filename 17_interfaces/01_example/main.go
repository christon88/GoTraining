package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type triangle struct {
	base   float64
	height float64
}

func (z triangle) area() float64 {
	return z.base * z.height / 2
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	t := triangle{10, 5}
	info(s)
	info(t)
}
