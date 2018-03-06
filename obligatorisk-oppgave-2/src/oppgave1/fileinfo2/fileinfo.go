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

	bytes := float64(fileInfo.Size())
	kilo := bytes / 1024
	mega := kilo / 1024
	giga := mega / 1024

	//Informasjon om filen printes ut via Println
	fmt.Printf("Infomation about file:", filename)
	fmt.Println("Bytes: ", bytes)
	fmt.Println("Kilobytes: ", kilo)
	fmt.Println("Megabytes: ", mega)
	fmt.Println("Gigabytes: ", giga)

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