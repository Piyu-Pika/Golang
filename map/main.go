//map in go

package main

import "fmt"


func main() {
	studentsGrades := make(map[string]int)
	studentsGrades["John"] = 90
	studentsGrades["Mary"] = 80
	studentsGrades["Peter"] = 85
	studentsGrades["Sarah"] = 95
	studentsGrades["Piyush"] = 100
	
	for name, grade := range studentsGrades {
		println(name, grade)
	}

	fmt.Println("Piyush's grade is", studentsGrades["Piyush"])
	fmt.Println("John's grade is", studentsGrades["John"])
	studentsGrades["John"] = 100
	fmt.Println("John's grade is", studentsGrades["John"])
	delete(studentsGrades, "John")
	fmt.Println("John's grade is", studentsGrades["John"])

	//check if key exists
	grade, ispresent := studentsGrades["John"]
	if ispresent {
		fmt.Println("John's grade is", grade)
	} else {
		fmt.Println("John is not present")
	}
	//range over map
	for name, grade := range studentsGrades {
		fmt.Println(name, grade)
	}



}