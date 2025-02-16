//loop in go
package main

import "fmt"

func main() {
	for i:=0; i<10; i++{
		fmt.Println("number is",i)

	}
	counter:=0
	for {
		fmt.Println("counter is",counter)
		counter++
		if counter==10{
			break
		}
	}
	num := [] int{1,2,3,4,5,6,7,8,9,10}
	for index,value := range num{
		fmt.Println("index is",index,"value is",value)
	}
	data:="hello world"
	for index,char:= range data{
		fmt.Printf("index is %d char is %c\n",index,char)
	}
}

