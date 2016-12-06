package main

import "fmt"

func main() {
	num := 600851475143
	for i := 2; i < num; i++ {
		if num%i == 0 {
			if isPrime(num / i) {
				fmt.Println(num / i)
				break
			}
		}
	}
}

func isPrime(x int) bool {
	for i := 2; i < x-1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
