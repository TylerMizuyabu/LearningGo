package main

import (
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("error happened", err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("noExist.txt")
	if err != nil {
		//log.Println("error Happened", err)
		log.Fatalln("error happened", err)
		//panic(err)
	}
}
