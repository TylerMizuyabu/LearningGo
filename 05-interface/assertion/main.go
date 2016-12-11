package main

import "fmt"

type namedObj struct {
	name string
}

type shape struct {
	color   int32
	centerX int
	centerY int
}

type rectangle struct {
	namedObj
	shape

	height int
	width  int
}

func main() {
	var name interface{} = "Tyler"
	//var name interface{} = 2
	var num interface{} = 3
	str, ok := name.(string)
	if ok {
		fmt.Println("Name is a string ", str)
	} else {
		fmt.Println("Name is not a string")
	}
	test(42, 36, "Mizuyabu", "Scott", 3.124159, 'a', rectangle{namedObj{"rectangle1"}, shape{12334, 4, 5}, 10, 20})

	//Must assert that num is an int
	fmt.Println(num.(int) + 2)
}

func test(v ...interface{}) {
	for _, val := range v {
		fmt.Printf("%T\n", val)
	}
}
