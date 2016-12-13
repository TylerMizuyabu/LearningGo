package main

import (
	"fmt"
	"sync"
)

var wb = sync.WaitGroup{}

func main() {
	wb.Add(2)
	go a()
	go b()
	wb.Wait()
}

func a() {
	for i := 1; i < 50; i++ {
		fmt.Println("func A")
	}
	wb.Done()
}

func b() {
	for i := 1; i < 50; i++ {
		fmt.Println("func B")
	}
	wb.Done()
}
