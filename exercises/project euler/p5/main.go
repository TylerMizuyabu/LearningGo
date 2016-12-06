package main

import "fmt"

//Way too inneficient
func main() {
	num := 20
	res := (num/2 + 1) * num
	endLoop := false
	for ; !endLoop; res += 20 {
		endLoop = true
		for i := num/2 + 1; i <= num; i++ {
			if res%i > 0 {
				endLoop = false
				break
			}
		}
		fmt.Println(res)
	}
}
