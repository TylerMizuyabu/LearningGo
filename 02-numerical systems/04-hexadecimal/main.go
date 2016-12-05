package main

import "fmt"

func main() {
	fmt.Printf("%d - %#x\n", 42, 42)
	fmt.Printf("%d - %#X\n", 42, 42)
	fmt.Printf("%d - %x\n", 42, 42)
	fmt.Printf("%d - %X", 42, 42)
}
