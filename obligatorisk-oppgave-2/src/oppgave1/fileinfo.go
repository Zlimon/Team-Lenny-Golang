package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Filename: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	fi, err := os.Lstat("")
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")

	}
}
