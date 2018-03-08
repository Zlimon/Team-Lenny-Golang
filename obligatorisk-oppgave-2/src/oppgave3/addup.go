package main

import (
	"fmt"
)

var number1 int
var number2 int
var total int

func main() {
	ch := make(chan int) // Oppretter kanaler med type int.

	// Leser input fra terminal, og putter dem inn i 2 kanaler.
	fmt.Print("Enter number 1: ")
	fmt.Scan(&number1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&number2)

	// Funksjon A.
	go func() {
		ch <- number1
		ch <- number2
	}()

	// Adderer tallene, og putter dem i 1 kanal.
	total = <-ch + <-ch
	// Funksjon B.
	go func() {
		ch <- total
	}()

	fmt.Println("Total =", <-ch) // Printer ut summen/siste kanalen.
}