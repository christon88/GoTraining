package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(1000)

	//Fan Out
	c1 := fact(in)
	c2 := fact(in)
	c3 := fact(in)
	c4 := fact(in)

	//Fan In
	for n := range merge(c1, c2, c3, c4) {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			for j := 1; j < 20; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fact(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
