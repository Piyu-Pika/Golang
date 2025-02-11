// startig go program
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World")
	fmt.Println("=====================================")
	var version string = "1.0.0"
	fmt.Println(version)
	var num int32 = 42
	var pi float64 = 3.14159
	var isRunning bool = true
	var letter rune = 'A'
	var message string = "Test Message"
	var values []int = []int{1, 2, 3}
	fmt.Println("Integer:", num)
	fmt.Println("Float:", pi)
	fmt.Println("Boolean:", isRunning)
	fmt.Println("Rune:", letter)
	fmt.Println("String:", message)
	fmt.Println("Slice:", values)

	person := 23

	var Public int = 10
	var private int = 20
	fmt.Println("Public:", Public)
	fmt.Println("Private:", private)

	println(person)
}
