// strings in go

package main

import (
	"fmt"
	"strings"
)

func main() {
	// split string
	str := "apple,mango,banana"
	parts := strings.Split(str, ",")
	for _, part := range parts {
		fmt.Println(part)
	}
	// count the word in a string
	str2 := "one two three four five"
	wordCount := strings.Count(str2, "two")
	fmt.Println(wordCount)

    //trim string
	str3 := "  hello world  "
	trimmed := strings.TrimSpace(str3)
	fmt.Println(trimmed)

	//Join string
	str4:= "Piyush"
	str5 := "Bhardwaj"
	resut := strings.Join([]string{str4, str5}, " ")
	fmt.Println(resut)
}
