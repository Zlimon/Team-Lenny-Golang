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

	fmt.Printf("Filinformasjon:", filename)
	fmt.Println("Bytes: ", Bytes)
	fmt.Println("Kilobytes: ", Kilo)
	fmt.Println("Megabytes: ", Mega)
	fmt.Println("Gigabytes: ", Giga)

	if fileInfo.Mode().IsDir() == true {
		fmt.Println("Er et directory")
	} else if fileInfo.Mode().IsDir() == false {
		fmt.Println("Er ikke et directory")


		if fileInfo.Mode().IsRegular() {
			fmt.Println("Er en regular file")
		} else {
			fmt.Println("Er ikke en regular file")
		}


		fmt.Println("Har Unix permission bits:", fileInfo.Mode().Perm())

		if fileInfo.Mode()&os.ModeAppend == os.ModeAppend {
			fmt.Println("Er append only")
		} else {
			fmt.Println("Er ikke append only")
		}


		if fileInfo.Mode()&os.ModeDevice == os.ModeDevice {
			fmt.Println("Er en device file")
		} else {
			fmt.Println("Er ikke en device file")
		}


		if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			fmt.Println("Er en symbolic link")
		} else {
			fmt.Println("Er ikke en Symbolic link")
		}
	}
}