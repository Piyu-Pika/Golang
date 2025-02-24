package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello", name)
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", name,"again")
}

func sayhi(name string) {
	fmt.Println("Hi", name)
}

func main() {
	fmt.Println("Learning Goroutines in Go")
	go sayHello("Piyush")
	go sayhi("Piyush")

	time.Sleep(4* time.Second)
}