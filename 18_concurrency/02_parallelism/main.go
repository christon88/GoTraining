package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//obsolete after go 1.5
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 30; i++ {
		fmt.Println("foo", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 30; i++ {
		fmt.Println("bar", i)
		time.Sleep(time.Duration(2 * time.Millisecond))
	}
	wg.Done()
}
