package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers: ", number)
	fmt.Printf("data type of number: %T\n", number)
	fmt.Println("Length of the slice: ", len(number))
}