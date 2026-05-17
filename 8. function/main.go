package main

import "fmt"

func simpleFunc() {
	fmt.Println("I am a simple function!")
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("I am learning Go!")
	simpleFunc()
	sum := add(132326, 1232)
	fmt.Println(sum)
}
