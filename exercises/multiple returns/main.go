package main

import "fmt"

func main() {
	div, isEven := divideEven(2)
	fmt.Printf("%v - %v", div, isEven)
}

func divideEven(x int) (int, bool) {
	return x / 2, x%2 == 0
}
