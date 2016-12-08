package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p Person
	fmt.Println(p)
	/*
	   NewDecoder requires a Reader Type
	   To have implemented the Reader interface the struct must implement the
	   read method
	   strings.NewReader returns a new Reader reading from passed in string argument
	   Thus we pass this into NewReader and get a Decoder to then Decode
	   the string and store the json result in our interface p

	   more info in godocs online
	*/
	rdr := strings.NewReader(`{"First":"Tyler","Last":"Mizuyabu","Age":20}`)
	json.NewDecoder(rdr).Decode(&p)

	fmt.Println(p)
}
