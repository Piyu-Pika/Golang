package main

import "fmt"

func main() {
	//  without defer
	fmt.Println("Starting of the program")
	fmt.Println("Middle of the program")
	fmt.Println("End of the program")

	//  with defer
	defer fmt.Println("End of the program 1")
	fmt.Println("Starting of the program2")
	fmt.Println("Middle of the program3")

	// use LIFO (Last In First Out) order
	// use in os operations, file operations, etc.
}