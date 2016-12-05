package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		printNum := true
		if i%3 == 0 {
			printNum = false
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			printNum = false
			fmt.Print("Buzz")
		}
		if printNum {
			fmt.Print(i)
		}
		fmt.Print("\n")
	}
}
