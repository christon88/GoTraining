package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 20; i++ {
		go factorial(c, i)
		fmt.Println(<-c)
	}
}

func factorial(c chan int, val int) {
	if val < 2 {
		c <- val
	} else {
		result := 1
		for i := val; i > 0; i-- {
			result *= i
		}
		c <- result
	}
}
