package main

/*
Interface is a type!!!

Bill Kennedy Quote:
"Once a type implements an interface, an entire world of functionality can be
opened up to values of that type"
*/

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, length float64
}

func (z Rectangle) area() float64 {
	return z.width * z.length
}

type Square struct {
	side float64
}

/*
  Square implements area() float64 so it can be used in functions that require
  the Shape interface type because it implements the Shape interface
*/
func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	radius float64
}

func (z Circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

/*
  Anything implementing the area() float64 method, implements the Shape
  interface
*/
type Shape interface {
	area() float64
}

func main() {
	s := Square{10}
	fmt.Printf("%T\n", s)
	info(s)
	c := Circle{2}
	info(c)
}
