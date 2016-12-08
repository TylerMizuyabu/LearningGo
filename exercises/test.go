package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"clan name"`
	Age   int
}

func main() {
	var p Person
	data := []byte(`{"First":"Tyler","Last":"Mizuyabu","Age":20}`)
	fmt.Println(string(data))
	fmt.Printf("%T\n", data)
	_ = json.Unmarshal(data, &p)
	fmt.Println(p)
}
