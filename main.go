package main

import (
	"fmt"
	"log"

	"github.com/raja-dettex/go-cache-client/client"
)

// usecase
func main() {
	defer handlePanic()
	client, err := client.New(":4000")
	if err != nil {
		panic(err)
	}
	res, err := client.Get("testKey")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func handlePanic() {
	if err := recover(); err != nil {
		fmt.Println("can not create client")
	}
}
