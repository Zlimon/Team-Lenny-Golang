package main

import (
	"fmt"
)

var number1 int
var number2 int
var total int

func main() {
	ch := make(chan int)

	fmt.Print("Enter number 1: ")
	fmt.Scan(&number1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&number2)
	go func() {
		ch <- number1
		ch <- number2
	}()

	//fmt.Println("Number 1:", <-ch,"\nNumber 2:", <-ch)

	total = number1 + number2
	go func() {
		ch <- total
	}()

	fmt.Println("Total =", <-ch)
}