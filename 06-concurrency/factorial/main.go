package main

import "fmt"

func main() {
	ints := make([]int, 0)
	for i := 3; i < 13; i++ {
		ints = append(ints, i)
	}
	input := gen(ints...)
	factorial(input)
}

func fac(i int) int {
	v := 1
	for n := 2; n <= i; n++ {
		v *= n
	}
	return v
}

func factorial(input chan int) {
	out := make(chan int)
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			for n := range input {
				out <- fac(n)
			}
			done <- true
		}()
	}
	go func() {
		i := 1
		for n := range out {
			fmt.Println(i, "\t", n)
			i++
		}
	}()
	for i := 0; i < 10; i++ {
		<-done
	}
	close(out)
	close(done)
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			for _, n := range nums {
				out <- n
			}
		}
		close(out)
	}()
	return out
}
