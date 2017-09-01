package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 26 {
			break
		}
	}
	fmt.Println("That was dumb")
	for i := 0; i <= 26; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
