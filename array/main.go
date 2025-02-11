package main

import "fmt"

func main() {
	fmt.Println("Learining Array!")
	var names [5]string
	names[0] = "Piyush"
	names[1] = "Rahul"
	names[2] = "Raj"
	fmt.Println("Names of the person: ", names)
	var nums = [8]int{1, 2, 3, 4, 5}
	fmt.Println("Length of the array: ", len(nums))

	fmt.Println("Numbers: ", nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
