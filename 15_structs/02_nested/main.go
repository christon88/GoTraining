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

func (e employee) scan() string {
	var certified string
	if e.certification {
		certified = "certified"
	}
	return e.first + e.last + certified
}

func main() {
	p1 := employee{
		person: person{
			first: "John",
			last:  "Doe",
			age:   20,
		},
		employeeNo:    25,
		certification: true,
	}

	fmt.Println(p1.scan())
}
