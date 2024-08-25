package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
		return
	}
	defer res.Body.Close()
	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
		return
	}
	content := string(databytes)
	fmt.Println(content)
}
