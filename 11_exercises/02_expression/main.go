package main

import "fmt"

func main() {
	half := func(n int) (float64, bool) {
		return float64(n) / 2, n%2 == 0
	}

	fmt.Println(half(4))
	fmt.Println(half(5))
}
