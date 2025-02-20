package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning Urls in Go")
	myurl := "https://www.exaple.com/path/to/file.html?query=string#fragment"
	parsedUrl, _ := url.Parse(myurl)
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)
	fmt.Println(parsedUrl.Fragment)
}
