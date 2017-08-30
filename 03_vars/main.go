package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%v \t %T \n", d, d)
}
