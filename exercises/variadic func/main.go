package main

import "fmt"

func main() {
	largest(15, 2, 6, 4, 8, 63, 6, 13, 5, 66, 1, 65)
}

func largest(nums ...float64) {
	var x float64
	if len(nums) > 0 {
		x = nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > x {
				x = nums[i]
			}
		}
	}
	fmt.Print(x)
}
