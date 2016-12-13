package main

import (
	"errors"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

//ErrNorGateMath is custom error for trying to square root a negative number
var ErrNorGateMath = errors.New("norgate math: square root of negative number")

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorGateMath
	}
	return math.Sqrt(f), nil
}

//examples in std library
//https://www.golang.org/src/bufio/bufio.go
