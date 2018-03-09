package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"os/signal"
	"syscall"
)

func main() {
	filename := os.Args[1]
	checkFile(filename)
}

func checkFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	errorCheck(err)

	str := string(file) // Konverterer filen til en string.

	completeCheck := "Calculation complete!" // Referansen som sjekker om filen allerede er kalkulert.

	/**
	En if-statement som først sjekker om filen allerede er kalkulert.
	Hvis ikke blir det spurt om å skrive inn 2 tall som skrives til file.txt.
	 */
	if strings.Contains(str, completeCheck) {
		fmt.Println("The answer is:", str)
	} else {
		addToFile(filename)
	}
}

func addToFile(filename string) {
	file, err := os.Create(filename)
	errorCheck(err)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number 1: ")
	number1, err := reader.ReadString('\n')
	fmt.Print("Enter number 2: ")
	number2, err := reader.ReadString('\n')

	writer := bufio.NewWriter(file)

	// Sjekker om variablene inneholder uhyggelige tegn som ødelegger for sumfromfile.go.
	if number1 == "\n" {
		fmt.Println("You have to write something in number 1!")
	} else if strings.Contains(number1, " ") {
		fmt.Println("You can not use space in number 1")
	} else if number2 == "\n" {
		fmt.Println("You have to write something in number 2!")
	} else if strings.Contains(number2, " ") {
		fmt.Println("You can not use space in number 2")
	} else {
		writer.WriteString(number1)
		writer.WriteString(number2)
		fmt.Println("Writing numbers to file completed!")
	}

	writer.Flush()
}

func errorCheck(err error) {
	c := make(chan os.Signal, 0x2)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Printf("\nYou canceled the program!\n")
		os.Exit(1)
	}()
}