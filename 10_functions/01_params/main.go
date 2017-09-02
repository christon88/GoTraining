package main

import "fmt"

func main() {
	greet("Christian")
	fmt.Println(formal_greet("Ella"))
	fmt.Println(average(15, 25, 35))
  numbers := []float64{25, 35, 45}
  fmt.Println(average(numbers...))
}

func greet(name string) {
	fmt.Println("Hello " + name + "!")
}
func formal_greet(name string) string {
	return fmt.Sprint("Greetings " + name + "!")
}

func average(an ...float64) float64 {
	var sum float64
	n := float64(len(an))
	for _, v := range an {
		sum += v
	}
	return sum / n
}
