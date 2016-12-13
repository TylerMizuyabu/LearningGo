package main

import (
	"fmt"
	"strconv"
)

/*
var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Process: "+s+" printing:", i)
	}
	wg.Done()
}
*/

/*
Challenge
-- Take the code from above and change it to use channels instead of wait groups and anomicity
*/

var count int

func main() {
	for n := range incrementor(2) {
		fmt.Println(n)
		count++
	}
	fmt.Println("Final count is: ", count)

}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(n int) {
			for i := 0; i < 20; i++ {
				c <- "Process: " + strconv.Itoa(n) + " printing: " + strconv.Itoa(i)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}
