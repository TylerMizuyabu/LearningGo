package main

import "fmt"

func main() {
	divideEven := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	div, isEven := divideEven(2)
	fmt.Printf("%v - %v", div, isEven)
}
