package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported string
}

func main() {
	p := Person{First: "Tyler", Last: "Mizuyabu", Age: 20, notExported: "This string won't be marshalled"}
	fmt.Println(p)
	//Will only marshal fields that can be exported
	bs, _ := json.Marshal(p)
	fmt.Println(bs)
	fmt.Println()
	fmt.Println(string(bs))
}
