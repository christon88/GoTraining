package main

import (
	"bufio"
	"fmt"
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

	//set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 12)

	//loop over words
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
	//fmt.Println("************")
	//for i := 28; i <= 126; i++ {
	//fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	//}
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
