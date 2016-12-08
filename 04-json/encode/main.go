package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Tyler", "Mizuyabu", 20, 1}
	/*
	   NewEncoder requires a io.Writer Type
	   Stdout is created by making a new *File (i.e its a pointer to a File type)
	   File Types implement the Write method, which means they implicitly implement
	   the Writer interface
	   Therefore Stdout implements the Writer interface and can be used in NewEncoder

	   more info at:
	   https://golang.org/pkg/io
	   https://golang.org/pkg/os
	   https://golang.org/pkg/encoding/json
	*/
	json.NewEncoder(os.Stdout).Encode(p1)
}
