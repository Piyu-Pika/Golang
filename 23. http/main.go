package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning HTTP")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	fmt.Printf("type of res: %T\n", res)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
