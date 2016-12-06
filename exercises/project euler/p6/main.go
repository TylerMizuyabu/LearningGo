package main

import "fmt"

func main() {
	var squareOfSum, sumOfSquares int
	num := 100

	for i := 1; i <= num; i++ {
		squareOfSum += i
		sumOfSquares += i * i
	}
	squareOfSum *= squareOfSum
	fmt.Println(squareOfSum - sumOfSquares)
}
