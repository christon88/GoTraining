package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{9, 5, 48, 16, 42, 2}
	sort.Ints(n)
	fmt.Println(n)
}
