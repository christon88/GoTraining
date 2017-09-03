package main

import "fmt"

func main() {
	greatest := func(n ...int) int {
		var max int
		for _, i := range n {
			if i > max {
				max = i
			}
		}
		return max
	}

	fmt.Println(greatest(4, 6, 8, 2, 22, 42, 41, 05, 26))
}
