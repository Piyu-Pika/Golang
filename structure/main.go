// structure in go
package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks int
}

func main() {
	var student1 Student
	student1.name = "Piyush"
	student1.age = 20
	student1.marks = 100
	println(student1.name, student1.age, student1.marks)
	student1.age = 21
	println(student1.name, student1.age, student1.marks)
	fmt.Println(student1)

	var student2 Student
	student2.name = "John"
	student2.age = 20
	student2.marks = 100
	println(student2.name, student2.age, student2.marks)
	student2.age = 21
	println(student2.name, student2.age, student2.marks)
	fmt.Println(student2)

	var student3 Student
	student3.name = "modi"
	student3.age = 20
	student3.marks = 60
	println(student3.name, student3.age, student3.marks)
	student3.age = 21
	println(student3.name, student3.age, student3.marks)
	fmt.Println(student3)

}
