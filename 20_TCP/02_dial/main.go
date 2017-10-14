package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	ln, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	bs, _ := ioutil.ReadAll(ln)
	fmt.Println(string(bs))
}
