package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#o \t %#X \t %q \n", i, i, i, i, i)
	}
}
