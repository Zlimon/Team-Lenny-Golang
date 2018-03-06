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
	text, _ := reader.ReadString( '\n')
	fmt.Println(text)
	fi, err := os.Lstat(text)
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
	case mode&os.ModeAppend != 0:
		fmt.Println("Append")
	case mode&os.ModeDevice != 0:
		fmt.Println("Device")
	case mode&os.ModePerm != 0:
		fmt.Println(mode&os.ModePerm)
	}
}
