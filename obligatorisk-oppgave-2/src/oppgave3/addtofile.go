package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	f, err := os.Create("C:/Users/simon/go/src/Go/Team-Lenny/Team-Lenny/obligatorisk-oppgave-2/src/oppgave3/fil.txt")

	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number 1: ")
	input, err := reader.ReadString('\n')
	fmt.Print("Enter number 1: ")
	input2, err := reader.ReadString('\n')


	write, err := writer.WriteString(input)
	write2, err := writer.WriteString(input2)

	fmt.Println(write, write2)

	writer.Flush()

	if err != nil {
		log.Fatal(err)
	}
}