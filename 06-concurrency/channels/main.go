package main

import (
	"fmt"
	"sync"
	"time"
)

var c = make(chan int)
var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go a()
	go cFunc()
	wg.Wait()
}

func a() {
	i := 5
	for x := 0; x < 10; x++ {
		c <- i
		//fmt.Printf("sent i %v\n", i)
		//i = <-c
		i *= 4
		fmt.Println("Waiting 5 seconds...")
		time.Sleep(time.Millisecond * 5000)
	}
	wg.Done()
}

func b() {
	for x := 0; x < 10; x++ {
		i := <-c
		i /= 2
		fmt.Printf("divided i by 2, %v\n", i)
		c <- i
	}
	wg.Done()
}

func cFunc() {
	for {
		i := <-c
		fmt.Println(i)
	}

}
