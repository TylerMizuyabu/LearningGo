package main

import "fmt"

func main() {
	//Yes its redundant. Ignore the error
	fmt.Print((true && false) || (false && true) || !(false && false))
}
