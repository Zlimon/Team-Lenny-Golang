package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	fileInfo(filename)
}

func fileInfo(filename string) {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}

	Bytes := float64(fileInfo.Size())
	Kilo := Bytes / 1024
	Mega := Kilo / 1024
	Giga := Mega / 1024

	fmt.Printf("Infomation about file:", filename)
	fmt.Println("Bytes: ", Bytes)
	fmt.Println("Kilobytes: ", Kilo)
	fmt.Println("Megabytes: ", Mega)
	fmt.Println("Gigabytes: ", Giga)

	if fileInfo.Mode().IsDir() == true {
		fmt.Println("Is a directory")
	} else if fileInfo.Mode().IsDir() == false {
		fmt.Println("Is not a directory")


		if fileInfo.Mode().IsRegular() {
			fmt.Println("Is a regular file")
		} else {
			fmt.Println("Is not a regular file")
		}


		fmt.Println("Has Unix permission bits:", fileInfo.Mode().Perm())

		if fileInfo.Mode()&os.ModeAppend == os.ModeAppend {
			fmt.Println("Is append only")
		} else {
			fmt.Println("Is not append only")
		}


		if fileInfo.Mode()&os.ModeDevice == os.ModeDevice {
			fmt.Println("Is a device file: true")
		} else {
			fmt.Println("Is a device file: false")
		}


		if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			fmt.Println("Is a symbolic link")
		} else {
			fmt.Println("Is not a Symbolic link")
		}
	}
}