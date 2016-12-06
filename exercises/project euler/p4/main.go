package main

import (
	"fmt"
	"strconv"
)

func main() {
	//starts off as the largest number possible 999*999
	for num := 999 * 999; num > 100*100; num-- {
		fmt.Println(num)
		if isPalindrome(strconv.Itoa(num)) && threeDigitFactors(num) {
			//fmt.Println(num)
			break
		}
	}
}

//Makes sure that the factors are 3 digits
func threeDigitFactors(x int) bool {
	for i := 999; i > x/1000; i-- {
		if x%i == 0 {
			return true
		}
	}
	return false
}

//returns true if the string given is a palindrome, false otherwise
func isPalindrome(s string) bool {
	length := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[length-i] {
			return false
		}
	}
	return true
}
