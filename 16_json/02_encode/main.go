package main

import (
	"encoding/json"
	"os"
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
	json.NewEncoder(os.Stdout).Encode(p1)
}
