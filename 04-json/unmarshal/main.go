package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	/*
		Because used the tag to json:"clan name", field will not change to "Mizuyabu"
		unless key in json byte slice is also "clan name" instead of "Last"
	*/
	Last string `json:"clan name"`
	Age  int
}

func main() {
	var p Person
	fmt.Println(p)
	data := []byte(`{"First":"Tyler","Last":"Mizuyabu","Age":20}`)
	_ = json.Unmarshal(data, &p)
	fmt.Println(p)
}
