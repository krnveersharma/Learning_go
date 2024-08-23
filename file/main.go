package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file")
	// 	return
	// }
	// defer file.Close() //After main function execution is completed then file resource will be freed
	// content := "hello world by Karan"
	// io.WriteString(file, content+"\n")
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while creating file")
		return
	}
	defer file.Close()

	//create buffer to read file content
	buffer := make([]byte, 12)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading the file")
			return
		}
		//process the read content
		fmt.Println(string(buffer[:n]))
	}

	//shortcut to read file but load all the filein memory in single time but buffer takes chunk of data and put in bytes and read them
	content, error := os.ReadFile("example.txt")
	if error != nil {
		fmt.Println("Error while reading the file")
	}
	fmt.Println(string(content))

}
