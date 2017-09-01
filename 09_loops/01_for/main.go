package main

import"fmt"

func main() {
	i := 0
	for i < 5 {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
		i++
	}
}
