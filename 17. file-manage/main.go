package main

import (
	"fmt"
	"os"
)

func main() {
	//  create file
	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()
	// // why to use f.close()? : it free the resource of the file which is not needed anymore.the
	// fmt.Println("file created successfully")

	// //  write to file
	// _, err = f.WriteString("hello world")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // read from file
	// b, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(b))
	// // delete file
	// err = os.Remove("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// file,err :=os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// fmt.Println("file opened successfully")

	// buf := make([]byte, 1024)
	// for {
	// 	n, err := file.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		fmt.Println(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	fmt.Println(string(buf[:n]))
	// }

	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err != nil && err != io.EOF {
	// 		fmt.Println(err)
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))

}
