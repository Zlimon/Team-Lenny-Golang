package main

import (
	"fmt"
	"log"
	"os"
	//"bufio"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter Filename: ")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	fi, err := os.Lstat("fileinfo.go")
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
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")

	case mode&os.ModeAppend != 0:
		fmt.Println("Append")


	}
}
