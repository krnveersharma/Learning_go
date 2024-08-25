package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	uri := "https://www.google.com"
	parsedURL, err := url.Parse(uri)
	if err == nil {

		fmt.Printf("Type of parsedurl is %T\n", parsedURL)
		fmt.Println(parsedURL)
		fmt.Println("Scheme:-", parsedURL.Scheme)
	}
}
