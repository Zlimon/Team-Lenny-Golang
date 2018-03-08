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
	fil2, err := ioutil.ReadFile("C:/Users/simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave3/file.txt")
	errorCheck(err)

	str := string(fil2) //Konverterer filen til en string.

	completeCheck := "Calculation complete!" //Referansen som sjekker om filen allerede er kalkulert.

	/**
	En metode som først sjekker om filen allerede er kalkulert.
	Hvis ikke blir det spurt om å skrive inn 2 tall som skrives til file.txt.
	 */
	if strings.Contains(str, completeCheck) {
		fmt.Println("Summen av tallene er:", str)
	} else {
		fil, err := os.Create("C:/Users/simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave3/file.txt")
		errorCheck(err)

		totalReader := bufio.NewScanner(fil)

		totalReader.Split(bufio.ScanLines)

		reader := bufio.NewReader(os.Stdin)
		writer := bufio.NewWriter(fil)

		fmt.Print("Enter number 1: ")
		number1, err := reader.ReadString('\n')
		errorCheck(err)
		fmt.Print("Enter number 2: ")
		number2, err := reader.ReadString('\n')
		errorCheck(err)

		//number1int := ("Number_1:" + number1)

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
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}