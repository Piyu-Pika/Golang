// pointer in go

package main

import "fmt"


func main() {
	var x int = 10
	var stri string = "hello"
	var add *string = &stri
	var y *int = &x
	*y = 20
	fmt.Println(x, y)
	x = 30
	fmt.Println(x, y)
	y = &x
	*y = 40
	fmt.Println(x, y)
	fmt.Println(add)
}