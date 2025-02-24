package main

import (
	"fmt"
)

func main() {
    // Don't Change this!
    scores := [5]int{24,56,78,98,67}
    // Your Code Here
    sum := 0 
    for i := 0; i <len(scores); i++ {
    sum = sum + scores[i]

}
average := sum / len(scores)
fmt.Println(average)

    
}