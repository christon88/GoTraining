package main

import "fmt"

func main() {
	for i := range factorial(numbers(20)) {
		fmt.Println(i)
	}
}

func numbers(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for a := range in {
			out <- fact(a)
		}

		close(out)
	}()
	return out
}

func fact(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}
