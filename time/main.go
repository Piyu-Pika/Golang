package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
    // fmt.Printf("Type of currentTime is %T\n", currentTime)
	formated := currentTime.Format("02-01-2006 15:04:05")
	fmt.Println(formated)
	formated2 := "2022-01-01 15:04:05"
	layout_str := "2006-01-02 15:04:05"
	layout, _ := time.Parse(layout_str, formated2)
	fmt.Println(layout)
}
