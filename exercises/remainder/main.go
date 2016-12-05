package main

import "fmt"

func main() {
	var small, large int
	fmt.Print("Please enter a small number: ")
	fmt.Scanf("%d", &small)
	fmt.Print("Please enter a large number: ")
	fmt.Scanf("%d", &large)
	fmt.Print("The remainder of the large number divided by the small number is ...\n", large%small)
}
