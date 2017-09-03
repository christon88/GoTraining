package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	n := 29
	fmt.Println(m, " - ", n)
	changeString(m)
	changeNo(&n)
	fmt.Println(m, " - ", n)
}

func changeString(z []string) {
	z[0] = "Christian"
}

func changeNo(n *int) {
	*n = 30
}
