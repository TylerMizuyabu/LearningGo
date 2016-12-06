package main

import "fmt"

func main() {
	nth := 10001
	primes := []int{2, 3, 5, 7, 11, 13}
	isPrime := func(x int) bool {
		for i := 0; i < len(primes); i++ {
			if x%primes[i] == 0 {
				return false
			}
		}
		return true
	}
	for num := 16; len(primes) < nth; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	fmt.Println(primes[nth-1])
}
