package main

import (
	"encoding/json"
	"fmt"
)

//Person is exported because changing from caps is tiresome
type Person struct {
	First       string
	Last        string `json:"-"`
	Age         int    `json:"experience"`
	notExported int
}

func main() {
	p1 := Person{"James", "Doe", 21, 42}
	bs, _ := json.Marshal(p1)
	var um Person
	json.Unmarshal(bs, &um)
	fmt.Println(um.First)
	fmt.Println(um.Last) //Will not show
	fmt.Println(um.Age)
	fmt.Printf("%T, \n", um)
}
