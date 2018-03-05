package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("File contents: %s", content)

	str := string(content)

	fmt.Println(str)
}