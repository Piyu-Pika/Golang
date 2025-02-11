package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Println("Hello", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello", name)
}
