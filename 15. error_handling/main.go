package main

import "fmt"

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
 
func main() {
	fmt.Println("Started error handling!")
	ans, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
