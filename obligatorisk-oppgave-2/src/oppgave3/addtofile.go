package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fil, err := os.Create("C:/Users/simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave3/file.txt")
	errorCheck(err)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(fil)

	fmt.Print("Enter number 1: ")
	number1, err := reader.ReadString('\n')
	fmt.Print("Enter number 2: ")
	number2, err := reader.ReadString('\n')

	number1int := ("Number_1:" + number1)

	writer.WriteString(number1int)
	writer.WriteString(number2)

	fmt.Println()

	writer.Flush()
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}