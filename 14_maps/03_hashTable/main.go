package main

import (
	"bufio"
	"log"
	"net/http"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

}
