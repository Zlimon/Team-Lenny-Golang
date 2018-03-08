package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	fil, err := os.Open("C:/Users/simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave3/file.txt")
	errorCheck(err)

	//str := string(fil) //Konverterer filen til en string.

	reader := bufio.NewScanner(fil)
	reader.Split(bufio.ScanLines)

	counter := 0
	var1 := 0
	var2 := 0
	total := 0
	for reader.Scan() {
		for counter < 1 {
			lineOne, err := strconv.Atoi(reader.Text()) //Konverterer linje 1 til int
			errorCheck(err)
			var1 = lineOne
			counter++
		}
		lineTwo, err := strconv.Atoi(reader.Text()) //Konverterer linje 2 til int
		errorCheck(err)
		var2 = lineTwo
		total = var1 + var2
	}
	fmt.Println("Summen:", var1, "+", var2, "=", total)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}