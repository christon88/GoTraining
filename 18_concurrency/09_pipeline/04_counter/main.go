package main

import "fmt"

func main() {
	n := 5
	for i := range numbers(n) {
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
