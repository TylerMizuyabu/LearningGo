package main

import "fmt"

func main() {
	fibFunc := fibWrapper()
	var sum int = 2
	for {
		var nextFib int = fibFunc()
		if nextFib > 4000000 {
			break
		}
		if nextFib%2 == 0 {
			sum += nextFib
		}
	}
	fmt.Print(sum)
}

func fibWrapper() func() int {
	var n, m int
	n = 1
	m = 2
	nextFib := func() int {
		var nextInt = n + m
		n = m
		m = nextInt
		return nextInt
	}
	return nextFib
}
