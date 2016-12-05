package main

import "fmt"

func main() {

	var upper int
	fmt.Print("Please enter an upper bound: ")
	fmt.Scanf("%d", &upper)

	//First find how the number of times 3 and 5 divide into 1000
	x := (upper - 1) / 3
	y := (upper - 1) / 5
	//Starts the sum off as the sum of all numbers soley divisible by 3 or 5
	//examples: 6, 25, 10
	sum := 3*(x*(x+1)/2) + 5*(y*(y+1)/2)

	//Eliminates numbers that were counted twice. Examples: 15, 30, 45
	z := (upper - 1) / 15
	sum = sum - 15*(z*(z+1)/2)

	fmt.Print(sum)

}
