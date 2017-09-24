package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	//var c int

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
			//c = i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
			//fmt.Println(c)
		}
	}()

	time.Sleep(time.Second)
}
