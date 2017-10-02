package main

import "fmt"

func main() {
	c := make(chan int)
	n := 4
	go factorial(c, n)
	fmt.Println(<-c)
}

func factorial(c chan int, n int) {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	c <- result
}
