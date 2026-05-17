package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println(num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Println(data)
	fmt.Printf("Type of data is %T\n", data)

	str := strconv.Itoa(num)
	fmt.Println(str)
	fmt.Printf("Type of str is %T\n", str)

	number := "1234567898"

	num2, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num2)
	fmt.Printf("Type of num2 is %T\n", num2)

	flowt := "123.456"
	data2, _ := strconv.ParseFloat(flowt, 64)
	fmt.Println(data2)
	fmt.Printf("Type of data2 is %T\n", data2)
}