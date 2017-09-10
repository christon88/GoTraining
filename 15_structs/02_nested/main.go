package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type employee struct {
	person
	employeeNo    int
	certification bool
}

func main() {
	p1 := employee{
		person: person{
			first: "John",
			last:  "Doe",
			age:   20,
		},
		employeeNo: 25,
	}

	fmt.Println(p1.first, p1.last, p1.age)
}
