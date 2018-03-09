package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"os/signal"
	"syscall"
)

func main() {
	filename := os.Args[1]
	checkFile(filename)
}

func checkFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	str := string(file) // Konverterer filen til en string.

	completeCheck := "Calculation complete!" // Referansen som sjekker om filen allerede er kalkulert.

	/**
	En if-statement som først sjekker om filen allerede er kalkulert.
	Hvis ikke blir tallene kalkulert, og summen blir skrevet inn i file.txt.
	 */
	if strings.Contains(str, completeCheck) {
		fmt.Println("This file is already calculated!")
	} else {
		sumFromFile(filename)
	}
}

func sumFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	counter := 0
	number1 := 0
	number2 := 0
	total := 0
	for reader.Scan() {
		for counter < 1 {
			lineOne, err := strconv.Atoi(reader.Text()) // Konverterer linje 1 til int.
			sigINT(err)
			number1 = lineOne
			counter++
		}
		lineTwo, err := strconv.Atoi(reader.Text()) // Konverterer linje 2 til int.
		sigINT(err)
		number2 = lineTwo
		total = number1 + number2
	}
	fmt.Println("Answer:", number1, "+", number2, "=", total)

	createFile, err := os.Create(filename)
	errorCheck(err)

	str := strconv.Itoa(total) // Konverterer variabel total med type int til string.

	createFile.WriteString(str)
	createFile.WriteString("\nCalculation complete!")
	fmt.Println("Calculation complete!")
}

// Feilhåndtering
func errorCheck(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func sigINT(err error) {
	c := make(chan os.Signal, 0x2)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Printf("\nYou canceled the program!\n")
		os.Exit(1)
	}()
}