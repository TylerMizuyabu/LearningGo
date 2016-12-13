package main

import "fmt"

func main() {
	var inc1 = incrementor()
	fmt.Printf("%T\n", inc1)
	//var inc2 = incrementor()
	for v := range inc1 {
		fmt.Println(v)
	}

	//fmt.Println("Now Printing inc2")
	//for v := range inc2 {
	//	fmt.Println(v)
	//}
}

func incrementor() <-chan int {
	var c = make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
